package main

import "fmt"

type Video struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func (v Video) get_link() string {
	return fmt.Sprintf("https://video.rbdata.ru/video/%v", v.Id)
}
