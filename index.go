//This file used for cloud function in yandex.cloud

package main

import (
	"context"
	"cska/db/rbdata"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func init() {
	//router inits here, so newer call
	//to cloud function would use this global value
	//and wont create router again
	router = gin.Default()

	apiv1 := router.Group("api/v1")
	apiv1.GET("/", MyHandler)
	//router.GET("/", MyHandler)
	//router.POST("/stuff", PostHandler)
}

func MyHandler(c *gin.Context) {
	c.String(http.StatusOK, "Привет, это апи сервер для телеграм бота @rbvideobot")
}

func Handler(rw http.ResponseWriter, req *http.Request) {
	//parsing request
	name := req.URL.Query().Get("videoName")
	count_string := req.URL.Query().Get("count")

	count, err := strconv.Atoi(count_string)
	if err != nil {
		fmt.Printf("cant parse query parameter %v %v\n", count_string, err)
		count = -1
	}

	fmt.Println("videoName:", name)

	//connecting to db and getting videos
	conn := rbdata.Connect()
	defer conn.Close(context.Background())
	videos := conn.GetVideos(name, count)

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
