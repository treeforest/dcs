package engine

import (
	"log"
	"github.com/treeforest/dcs/fetcher"
	)

type concurrentEngine struct {
	scheduler   Scheduler
	workerCount int
}

func NewConcurrentEngine(scheduler Scheduler, workerCount int) Enginer {
	return &concurrentEngine{
		scheduler:scheduler,
		workerCount:workerCount,
	}
}

func (ce *concurrentEngine) Run(seeds ...Request) {
	out := make(chan ParseResult)

	ce.scheduler.Run()

	for i := 0; i < ce.workerCount; i++ {
		ce.createWorker(out, ce.scheduler)
	}

	for _, s := range seeds {
		ce.scheduler.Submit(s)
	}

	itemCount := 1
	for {
		result := <-out

		for _, item := range result.Items {
			log.Printf("Got item: %d, %v\n", itemCount, item)
			itemCount++
		}

		for _, req := range result.Reqs {
			ce.scheduler.Submit(req)
		}
	}
}

func (ce *concurrentEngine) createWorker(out chan ParseResult, s Scheduler) {
	in := make(chan Request)
	go func() {
		for {
			s.WorkReady(in)

			request := <-in

			result, err := ce.worker(request)
			if err != nil {
				continue
			}

			out <- result
		}
	}()
}

func (ce *concurrentEngine) worker(req Request) (ParseResult, error){
	log.Printf("Fetch url: %s", req.Url)

	body, err := fetcher.WebFetch(req.Url)
	if err != nil {
		log.Printf("Fetch error: %s", req.Url)
		return ParseResult{}, err
	}

	return req.ParseFunc(body), nil
}