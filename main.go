package main

import (
	"github.com/treeforest/dcs/engin"
	"github.com/treeforest/dcs/parse"
	"github.com/treeforest/dcs/scheduler"
)

func main() {
	e := engine.NewConcurrentEngine(scheduler.CreateQueueScheduler(), 100)
	e.Run(engine.Request{
		Url:"https://book.douban.com",
		ParseFunc:parse.ParseTag,
	})

	//engine.Run(engine.Request{
	//	Url:"https://book.douban.com",
	//	ParseFunc:parse.ParseTag,
	//})
}