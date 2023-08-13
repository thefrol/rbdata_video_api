package main

type RequestBody struct {
	VideoName string `json:"videoName"`
}

type Request struct {
	Body RequestBody `json:"body"`
}
