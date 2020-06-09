package parse

import (
	"regexp"
		"github.com/treeforest/dcs/engine"
)

const regexpStr = `<a href="([^"]+)" class="tag">[^<]+</a>]`
func ParseContent(content []byte) engine.ParseResult{
	re := regexp.MustCompile(regexpStr)

	matches := re.FindAllSubmatch(content, -1)

	result := engine.ParseResult{}

	for _, m := range matches {
		result.Items = append(result.Items, m[2])
		result.Reqs = append(result.Reqs, engine.Request{
			Url:"https://jgxy.xmu.edu.cn/" + string(m[1]),
			ParseFunc: engine.NilParse,
		})
	}

	return result
}