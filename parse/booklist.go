package parse

import (
	"github.com/treeforest/dcs/engine"
	"regexp"
	"fmt"
)

const bookListRe = `<a href="([^"]+)" title="([^"]+)"`
func ParseBookList(contents []byte) engine.ParseResult {

	fmt.Printf("%s")

	re := regexp.MustCompile(bookListRe)

	matches := re.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}

	for _, m := range matches {
		result.Items = append(result.Items, string(m[2]))
		result.Reqs = append(result.Reqs, engine.Request{
			Url:string(m[1]),
			ParseFunc:engine.NilParse,
		})
	}

	return result
}
