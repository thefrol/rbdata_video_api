package main

type ResponseBody struct {
	List []Video `json:"videos"`
}

type Response struct {
	StatusCode int         `json:"statusCode"`
	Body       interface{} `json:"body"`
}
