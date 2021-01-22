package main

import (
	"fmt"
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

	m := http.NewServeMux()
	go CounterTicker()
	m.Handle("/days/", MakeHandler(*metricsService))
	http.Handle("/", m)

	fmt.Println("listening on port :8080")
	http.ListenAndServe(":8080", m)
}

