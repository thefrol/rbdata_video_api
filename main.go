package main

import (
	"context"
	"cska/db/rbdata"
	"encoding/json"
	"fmt"
)

func main() {
	textToSearch := "ЦСКА"

	conn := rbdata.Connect()
	defer conn.Close(context.Background())
	videos := conn.GetVideos(textToSearch, 10)

	r := ResponseBody{List: videos}
	body, err := json.Marshal(&r)
	if err != nil {
		println("Error marshalling json")
		panic(`json error`)
	}
	fmt.Println(string(body))
}
