package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
)

func main() {
	conn_string := os.Getenv("DB_CONN")
	conn, err := pgx.Connect(context.Background(), conn_string)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	rows, err := conn.Query(context.Background(),
		`SELECT id, "name"
		FROM public.videos
		WHERE "name" ILIKE '%ЦСКА%'
		ORDER BY created_at DESC
		LIMIT 10;`)

	if err != nil {
		fmt.Println("wrong query")
	}

	var videos []Video
	for rows.Next() {
		var video Video
		rows.Scan(&video.Id, &video.Name)
		//fmt.Println(video.get_link(), video.name)
		videos = append(videos, video)
	}
	r := Response{List: videos}
	body, err := json.Marshal(&r)
	if err != nil {
		println("Error marshalling json")
		panic(`json error`)
	}
	fmt.Println(string(body))

}
