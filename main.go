package main

import (
	"context"
	"encoding/json"
	"fmt"
)

func main() {
	textToSearch := "ЦСКА"

	conn := Connect()
	defer conn.Close(context.Background())
	videos := conn.GetVideos(textToSearch)

	r := ResponseBody{List: videos}
	body, err := json.Marshal(&r)
	if err != nil {
		println("Error marshalling json")
		panic(`json error`)
	}
	fmt.Println(string(body))
	resp := Response{StatusCode: 200, Body: string(body)}

	fmt.Printf("%v", resp)
}
