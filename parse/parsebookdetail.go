package parse

import (
	"regexp"
	"github.com/treeforest/dcs/model"
	"github.com/treeforest/dcs/engin"
	)

var authorRe = regexp.MustCompile(`<span class="pl">\s*作者:?</span>[\d\D]*?<a.*?>\s*([^<]+)</a>`)
var publicRe = regexp.MustCompile(`<span class="pl">出版社:</span>\s*(.*)<br/>`)
var pageRe = regexp.MustCompile(`<span class="pl">页数:</span>\s(\d*)<br/>`)
var priceRe = regexp.MustCompile(`<span class="pl">定价:</span>([^<]+)<br/>`)
var scoreRe = regexp.MustCompile(`<strong class="ll rating_num " property="v:average">([^<]+)</strong>`)
var introRe = regexp.MustCompile(`<div class="intro">[\d\D]*?<p>([^<]+)</p></div>`)
func ParseBookDetail(contents []byte, bookName string) engine.ParseResult{

	bookdetail := model.BookDetail{}

	bookdetail.BookName = bookName
	bookdetail.Author = extraString(contents, authorRe)
	bookdetail.Publicer = extraString(contents, publicRe)
	bookdetail.Pages = extraString(contents, pageRe)
	bookdetail.Price = extraString(contents, priceRe)
	bookdetail.Score = extraString(contents,scoreRe)
	bookdetail.Intro = extraString(contents, introRe)

	return engine.ParseResult{
		Items: []interface{}{bookdetail},
	}
}

func extraString(contents []byte, re *regexp.Regexp) string{
	match := re.FindSubmatch(contents)

	if len(match) >= 2 {
		return string(match[1])
	}

	return ""
}