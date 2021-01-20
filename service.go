package main

import (
	"time"
)

type CounterInterface interface {
	GetCounter(req getCounterRequest) *getCounterResponse
	NilCounter(req nilCounterRequest) *nilCounterResponse
}

type counterService struct {
	C Counter `json:"c"`
}

func (s *counterService) Init() {
	s.C = Counter{
		Days:        0,
		CurrentDate: time.Now(),
	}
}

func (s *counterService) GetCounter(req getCounterRequest) *getCounterResponse {
	resp := getCounterResponse{}

	t := time.Now()
	daysPassed := t.Sub(s.C.CurrentDate).Hours() / 24
	resp.Counter = int(daysPassed)

	return &resp
}

func (s *counterService) NilCounter(req nilCounterRequest) *nilCounterResponse {
	resp := nilCounterResponse{}

	s.C.Days = 0
	s.C.CurrentDate = time.Now()

	return &resp
}
