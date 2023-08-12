package main

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
)

type Video struct {
	id   int
	name string
}

func (v Video) get_link() string {
	return fmt.Sprintf("https://video.rbdata.ru/video/%v", v.id)
}

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

	for rows.Next() {
		var video Video
		rows.Scan(&video.id, &video.name)
		fmt.Println(video.get_link(), video.name)
	}

}
