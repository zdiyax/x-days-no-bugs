package main

import (
	"context"
	"encoding/json"
	"net/http"

	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

type ServerError struct {
	 StatusCode uint
	 ErrorMessage string
}

func(rs ServerError) Error () string {
	return rs.ErrorMessage
}

func MakeHandler(cs counterService) http.Handler {
	options := []kithttp.ServerOption{
		kithttp.ServerErrorEncoder(encodeError),
	}

	getCounter := kithttp.NewServer(
		makeGetCounterEndpoint(cs),
		decodeGetCounterRequest,
		encodeResponse,
		options...)

	nilCounter := kithttp.NewServer(
		makeNilCounterEndpoint(cs),
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
	if e, ok := r.(ServerError); ok && &e != nil {
		encodeError(context.Background(), e, w)
		return nil
	}

	return json.NewEncoder(w).Encode(r)
}

func encodeError(_ context.Context, err error, w http.ResponseWriter) {
	servErr := err.(ServerError)

	switch servErr.StatusCode {
	case 500:
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Sorry :( \n Server Internal error"))
	case 400:
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Bad request"))
	case 404:
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Not found"))
	case 413:
		w.WriteHeader(http.StatusRequestEntityTooLarge)
		w.Write([]byte("Request entity too large"))
	case 503:
		w.WriteHeader(http.StatusServiceUnavailable)
		w.Write([]byte("Service unavailable"))
	}

}
