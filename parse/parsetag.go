package parse

import (
	"regexp"
		"github.com/treeforest/dcs/engin"
)

const regexpStr = `<a href="([^"]+)" class="tag">([^<]+)</a>`
func ParseTag(content []byte) engine.ParseResult{
	re := regexp.MustCompile(regexpStr)

	matches := re.FindAllSubmatch(content, -1)

	result := engine.ParseResult{}

	for _, m := range matches {
		result.Items = append(result.Items, m[2])
		result.Reqs = append(result.Reqs, engine.Request{
			Url:"https://book.douban.com" + string(m[1]),
			ParseFunc: ParseBookList,
		})
	}

	return result
}