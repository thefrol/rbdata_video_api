//This file used for cloud function in yandex.cloud

package main

import (
	"context"
	"encoding/json"
	"fmt"
)

func Handler(ctx context.Context, request *Request) (*Response, error) {
	conn := Connect()
	defer conn.Close(context.Background())
	videos := conn.GetVideos(request.Body.VideoName)

	r := ResponseBody{List: videos}
	body, err := json.Marshal(&r)
	if err != nil {
		println("Error marshalling json")
		panic(`json error`)
	}
	fmt.Println(string(body))
	resp := Response{StatusCode: 200, Body: string(body)}

	return &resp, nil
}
