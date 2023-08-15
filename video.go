package main

import (
	"fmt"
	"time"
)

type Video struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"createdAt"`
}

func (v Video) get_link() string {
	return fmt.Sprintf("https://video.rbdata.ru/video/%v", v.Id)
}
