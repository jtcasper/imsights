package php

import (
	"bytes"
	_ "fmt"
	"github.com/jtcasper/imsights/types"
	"regexp"
)

type PhpSearcher struct {
	classRegex, funcRegex, callRegex *regexp.Regexp
}

func New() (s *PhpSearcher) {
	s = &PhpSearcher{
		classRegex: regexp.MustCompile("(?:class)\\s+(\\w+)"),
		funcRegex:  regexp.MustCompile("(?:public|protected|private)?\\s*(?:static\\s+)?(?:function)\\s+(\\w+)\\("),
		callRegex:  regexp.MustCompile("->(\\w+)\\("),
	}
	return
}

func (p *PhpSearcher) Search(b []byte, r *regexp.Regexp) [][]byte {
	return r.FindSubmatch(b)
}

func (p *PhpSearcher) SearchAll(b []byte) (c types.Class) {
	bcStart := regexp.MustCompile("/[*]")
	bcEnd := regexp.MustCompile("[*]/")
	comment := regexp.MustCompile("//")
	bcCount := 0
	for _, line := range bytes.Split(b, []byte{'\n'}) {
		if bcMatch := p.Search(line, bcStart); bcMatch != nil {
			bcCount++
		}
		if bcMatch := p.Search(line, bcEnd); bcMatch != nil {
			bcCount--
		}
		if bcCount > 0 {
			continue
		}
		if commentMatch := p.Search(line, comment); commentMatch != nil {
			continue
		}
		if c.Name == "" {
			classMatch := p.Search(line, p.classRegex)
			if classMatch != nil {
				c.Name = string(classMatch[0])
			}
		}
		if funcMatch := p.Search(line, p.funcRegex); funcMatch != nil {
			c.Functions = append(c.Functions, types.Function{
				Name: string(funcMatch[1]),
			})
		}
		if callMatch := p.Search(line, p.callRegex); callMatch != nil {
			calls := c.Functions[len(c.Functions)-1].Calls
			c.Functions[len(c.Functions)-1].Calls = append(calls, types.Call{
				Name: string(callMatch[1]),
			})
		}

	}
	return
}
