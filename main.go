package main

import (
	"github.com/treeforest/dcs/engine"
	"github.com/treeforest/dcs/parse"
)

func main() {
	engine.Run(engine.Request{
		Url:"https://book.douban.com/tag/%E7%A5%9E%E7%BB%8F%E7%BD%91%E7%BB%9C",
		ParseFunc:parse.ParseBookList,
	})
}