package main

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

func makeGetCounterEndpoint(cs counterService) endpoint.Endpoint {
	return func(ctx context.Context, r interface{}) (resp interface{}, err error) {
		req := r.(getCounterRequest)
		respCounter := cs.GetCounter(req)

		if respCounter.ServerError != nil {
			return nil, *respCounter.ServerError
		}


		return respCounter, nil
	}
}

func makeNilCounterEndpoint(cs counterService) endpoint.Endpoint {
	return func(ctx context.Context, r interface{}) (resp interface{}, err error) {
		req := r.(nilCounterRequest)

		respCounter := cs.NilCounter(req)

		if respCounter.ServerError != nil {
			return nil, *respCounter.ServerError
		}

		return respCounter, nil
	}
}
