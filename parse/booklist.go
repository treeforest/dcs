package parse

import (
	"github.com/treeforest/dcs/engin"
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
		bookName := string(m[2])
		result.Items = append(result.Items, bookName)
		result.Reqs = append(result.Reqs, engine.Request{
			Url:string(m[1]),
			ParseFunc:func(c []byte) engine.ParseResult {
				return ParseBookDetail(c, bookName)
			},
		})
	}

	return result
}
