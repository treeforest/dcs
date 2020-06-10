package scheduler

import "github.com/treeforest/dcs/engin"

type queueScheduler struct {
	requestChan chan engine.Request
	workerChan chan chan engine.Request
}

func CreateQueueScheduler() engine.Scheduler{
	return &queueScheduler{
		requestChan:make(chan engine.Request),
		workerChan:make(chan chan engine.Request),
	}
}

func (s *queueScheduler) Submit(req engine.Request) {
	s.requestChan <- req
}

func (s *queueScheduler) ConfigWorkChan(c chan engine.Request) {

}

func (s *queueScheduler) WorkReady(w chan engine.Request) {
	s.workerChan <- w
}

func (s *queueScheduler) Run() {
	go func() {
		var requestQueue []engine.Request
		var workQueue []chan engine.Request

		for {
			var activeRequest engine.Request
			var activeWork chan engine.Request
			if len(requestQueue) > 0 && len(workQueue) > 0 {
				activeRequest = requestQueue[0]
				activeWork = workQueue[0]
			}

			select {
				case r := <-s.requestChan:
					requestQueue = append(requestQueue, r)
				case w := <-s.workerChan:
					workQueue = append(workQueue, w)
				case activeWork <- activeRequest:
					workQueue = workQueue[1:]
			}
		}
	}()
}