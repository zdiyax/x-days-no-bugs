package main

import (
	"context"
	"encoding/json"
	"net/http"

	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

func MakeHandler(s counterService) http.Handler {
	options := []kithttp.ServerOption{
		kithttp.ServerErrorEncoder(encodeError),
	}

	getCounter := kithttp.NewServer(
		makeGetCounterEndpoint(&s),
		decodeGetCounterRequest,
		encodeResponse,
		options...)

	nilCounter := kithttp.NewServer(
		makeNilCounterEndpoint(&s),
		decodeNilCounterRequest,
		encodeResponse,
		options...)


	//todo: incrementCounter должен прибавлять +1 к дням и обновлять дату
	//incrementCounter := kithttp.NewServer(
	//	makeNilCounterEndpoint(&s),
	//	decodeNilCounterRequest,
	//	encodeResponse,
	//	options...)

	r := mux.NewRouter()

	r.Handle("/days/counter/", getCounter).Methods("GET")
	r.Handle("/days/nil/", nilCounter).Methods("GET")

	return r
}

func decodeGetCounterRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	apiRequest := getCounterRequest{}

	return apiRequest, nil
}

func decodeNilCounterRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	apiRequest := nilCounterRequest{}

	return apiRequest, nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, r interface{}) error {
	if e, ok := r.(errorInterface); ok && e.error() != nil {
		encodeError(context.Background(), e.error(), w)
		return nil
	}

	return json.NewEncoder(w).Encode(r)
}

func encodeError(_ context.Context, err error, w http.ResponseWriter) {
	_ = json.NewEncoder(w).Encode(err)
}

type errorInterface interface {
	error() error
}
