package tests

import (
	"context"
	"cska/db/rbdata"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Возвращает 10 видео содержищих в названии строку "ЦСКА"
func Find(c *gin.Context) {
	//connecting to db and getting videos
	const (
		videoName = "ЦСКА"
		count     = 10
	)
	conn := rbdata.Connect()
	defer conn.Close(context.Background())
	videos := conn.GetVideos(videoName, count)

	c.IndentedJSON(http.StatusOK, gin.H{"videos": videos, "count": len(videos)})
}

// Мучался сегодня очень долго с ошибкой, что
// This file is within module ".", which is not included in your workspace.
// И мой пакет не находил компилятор.
// Оказалось я назвал файл videos, вместо videos.go
