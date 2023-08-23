package rbdata

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
)

type RbData struct {
	*pgx.Conn
}

func (rb RbData) GetVideos(name string, count int) []Video {
	if count < 0 {
		count = 10 // default value
	}

	rows, err := rb.Query(context.Background(),
		`SELECT id, "name",created_at
	FROM public.videos
	WHERE "name" ILIKE $1
	ORDER BY created_at DESC
	LIMIT $2;`, "%"+name+"%", count)

	if err != nil {
		fmt.Println(`query error:`, err)
		os.Exit(1)
	}
	defer rows.Close()

	videos := []Video{}
	for rows.Next() {
		var video Video
		rows.Scan(&video.Id, &video.Name, &video.CreatedAt)
		//fmt.Println(video.get_link(), video.name)
		videos = append(videos, video)
	}

	return videos
}

func Connect() *RbData {
	conn_string := os.Getenv("DB_CONN")
	conn, err := pgx.Connect(context.Background(), conn_string)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	//defer conn.Close(context.Background())
	return &RbData{Conn: conn}
}
