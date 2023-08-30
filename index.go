//This file used for cloud function in yandex.cloud

package main

import (
	v1 "cska/db/routes/api/v1"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func init() {
	//router inits here, so newer call
	//to cloud function would use this global value
	//and wont create router again
	router = gin.Default()

	apiv1 := router.Group("api/v1")
	apiv1.GET("/", v1.Home)
	apiv1.POST("videos", v1.Videos)
}

// Handler is a basic handle for Yandex.Cloud.Functions
// Though we can use "index.router" for this role when configuring our function
func Handler(w http.ResponseWriter, r *http.Request) {

	fmt.Println(*r)
	router.ServeHTTP(w, r) // ServeHTTP conforms to the http.Handler interface (https://godoc.org/github.com/gin-gonic/gin#Engine.ServeHTTP)
}
