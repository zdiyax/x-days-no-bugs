package main

type getCounterRequest struct {
}

type getCounterResponse struct {
	Counter int `json:"counter"`
	ServerError *ServerError `json:"-"`
}

type nilCounterRequest struct {
}

type nilCounterResponse struct {
	Success bool `json:"success"`
	ServerError *ServerError `json:"-"`
}
