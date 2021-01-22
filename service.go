package main

import (
	"time"
)

type CounterInterface interface {
	GetCounter(req getCounterRequest) *getCounterResponse
	NilCounter(req nilCounterRequest) *nilCounterResponse
}

type counterService struct {
	CounterColl CounterCollectionInterface
}

func (s *counterService) Init(mg CounterCollectionInterface) {
	s.CounterColl = mg
	//= Counter{
	//	Days:        0,
	//	CurrentDate: time.Now(),
	//}
}

func (s *counterService) GetCounter(req getCounterRequest) *getCounterResponse {

	counter, err := s.CounterColl.GetCounterDB()
	if err != nil {
		resp := getCounterResponse{
			ServerError: &ServerError{},
		}
		resp.ServerError.StatusCode = 500
		resp.ServerError.ErrorMessage = err.Error()
		return &resp
	}

	t := time.Now()
	daysPassed := t.Sub(counter.CurrentDate).Hours() / 24
	resp := getCounterResponse{
		Counter: int(daysPassed),
	}
	return &resp
}

func (s *counterService) NilCounter(req nilCounterRequest) *nilCounterResponse {

	err := s.CounterColl.NilCounterDB()
	if err != nil {
		resp := nilCounterResponse{
			Success: false,
			ServerError: &ServerError{
				StatusCode: 500,
				ErrorMessage: err.Error(),
			},

		}
		return &resp
	}
	resp := nilCounterResponse{
		Success: true,
	}

	return &resp
}
