package main

import "cska/db/rbdata"

type ResponseBody struct {
	List []rbdata.Video `json:"videos"`
}

type Response struct {
	StatusCode int         `json:"statusCode"`
	Body       interface{} `json:"body"`
}
