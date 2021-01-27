package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	dbConfig := MongoConfig{
		Host:     "localhost",
		Database: "example",
		Port:     "27017",
	}
	
	metricsService := &counterService{}
	counterCol, err := InitCounterCollection(dbConfig)
	if err != nil {
		panic(err)
	}
	metricsService.Init(counterCol)

	m := mux.NewRouter()

	go CounterTicker()

	m.PathPrefix("/days").Handler( MakeHandler(*metricsService))

	fmt.Println("listening on port :8080")
	http.ListenAndServe(":8080", m)
}

