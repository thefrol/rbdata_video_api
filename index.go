//This file used for cloud function in yandex.cloud

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func Handler(rw http.ResponseWriter, req *http.Request) {
	//parsing request
	name := req.URL.Query().Get("name")
	fmt.Println("videoName:", name)

	//connecting to db and getting videos
	conn := Connect()
	defer conn.Close(context.Background())
	videos := conn.GetVideos(name)

	r := ResponseBody{List: videos}
	body, err := json.Marshal(&r)
	if err != nil {
		println("Error marshalling json")
		panic(`json error`)
	}

	//writing response

	rw.Header().Set("X-Custom-Header", "Test")
	rw.WriteHeader(200)

	io.WriteString(rw, string(body))
}
