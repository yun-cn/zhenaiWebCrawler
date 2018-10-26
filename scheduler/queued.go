package scheduler

import "crawler/engine"

// QueuedScheduler type
type QueuedScheduler struct {
	requestChan chan engine.Request
	workerChan  chan chan engine.Request
}

// Submit Func
func (s QueuedScheduler) Submit(r engine.Request) {
	s.requestChan <- r
}

// WorkerReady Func
func (s QueuedScheduler) WorkerReady(w chan engine.Request) {
	s.workerChan <- w
}

func (s QueuedScheduler) Run() {
	go func() {
		var requestQ []engine.Request
		var workerQ []chan engine.Request
		for {
			select {
			case r := <-s.requestChan:
				requestQ = append(requestQ, r)
			case w := <-s.workerChan:
				workerQ = append(workerQ, w)
			}
		}
	}()
}
