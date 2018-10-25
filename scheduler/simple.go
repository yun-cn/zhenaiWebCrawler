package scheduler

import "crawler/engine"

// SimpleScheduler type
type SimpleScheduler struct {
	workerChan chan engine.Request
}

// ConfigureMasterWorkerChan func
func (s *SimpleScheduler) ConfigureMasterWorkerChan(c chan engine.Request) {
	s.workerChan = c
}

// Submit func
func (s *SimpleScheduler) Submit(r engine.Request) {
	// Send request down to woker chan
	go func() { s.workerChan <- r }()
}
