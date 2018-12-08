package search

import (
	_ "bytes"
	_ "fmt"
	"github.com/jtcasper/imsights/types"
	"regexp"
)

type Searcher interface {
	New(ext string) *Searcher
	Search([]byte, *regexp.Regexp) []byte
	SearchAll([]byte) types.Class
}
