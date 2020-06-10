package scheduler

import "github.com/treeforest/dcs/engin"

type simpleScheduler struct {
	workerChan chan engine.Request
}

//func NewSimpleScheduler() Scheduler {
//	ss := new(simpleScheduler)
//	ss.workerChan = make(chan Request)
//	return ss
//}

func (ss *simpleScheduler)Submit(req engine.Request) {
	go func() {
		ss.workerChan <- req
	}()
}

func (ss *simpleScheduler)ConfigWorkChan(c chan engine.Request) {
	ss.workerChan = c
}