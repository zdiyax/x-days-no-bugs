package main

type getCounterRequest struct {
}

type getCounterResponse struct {
	Counter int `json:"counter"`
}

type nilCounterRequest struct {
}

type nilCounterResponse struct {
	Success bool `json:"success"`
}
