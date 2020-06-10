package engine

type Enginer interface {
	Run(seeds ...Request)
}

type Scheduler interface {
	Submit(req Request)
	ConfigWorkChan(c chan Request)
	WorkReady(w chan Request)
	Run()
}