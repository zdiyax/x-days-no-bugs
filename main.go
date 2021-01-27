package main

import (
	"fmt"
	"net/http"
)

func main() {
	metricsService := &counterService{}
	metricsService.Init()

	m := http.NewServeMux()
	m.Handle("/days/", MakeHandler(*metricsService))
	http.Handle("/", handleAll(m))

	fmt.Println("listening on port :8083")
	http.ListenAndServe(":8083", m)
}



func handleAll(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		h.ServeHTTP(w, r)
	})
}
