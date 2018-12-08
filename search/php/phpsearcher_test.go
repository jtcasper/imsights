package php_test

import (
	"fmt"
	"github.com/jtcasper/imsights/search/php"
	_ "testing"
)

func ExampleSearchAll() {
	b := []byte("class Test\nprotected function main()\n/* ->hey\n*/ test->Call(i)")
	r := php.New()
	c := r.SearchAll(b)
	fmt.Printf("%+v\n", c)
	// Output: {Name:class Test Functions:[{Name:main Calls:[{Name:Call}]}]}
}
