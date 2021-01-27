package main

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

func makeGetCounterEndpoint(s *counterService) endpoint.Endpoint {
	return func(ctx context.Context, r interface{}) (resp interface{}, err error) {
		req := r.(getCounterRequest)

		resp = s.GetCounter(req)

		return resp, nil
	}
}

func makeIncrementCounterEndpoint(s *counterService) endpoint.Endpoint {
	return func(ctx context.Context, r interface{}) (resp interface{}, err error) {
		req := r.(incrementCounterRequest)

		resp = s.IncrementCounter(req)

		return resp, nil
	}
}

func makeNilCounterEndpoint(s *counterService) endpoint.Endpoint {
	return func(ctx context.Context, r interface{}) (resp interface{}, err error) {
		req := r.(nilCounterRequest)

		resp = s.NilCounter(req)

		return resp, nil
	}
}
