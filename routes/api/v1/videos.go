package v1

import (
	"context"
	"cska/db/rbdata"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Videos receives a body with:
//
//	videoName: string
//	count: int
//
// and returns latest videos, containing `videoName“, maximum `count“ videos with such JSON
//
//	{
//		"count": videos_returned
//		"videos": list_of_videos
//	}
func Videos(c *gin.Context) {
	type Request struct {
		VideoName string `json:"videoName"`
		Count     int    `json:"count"`
	}
	var r Request
	if err := c.ShouldBindJSON(&r); err != nil {
		fmt.Println("Received bad data")
		c.String(http.StatusBadRequest, err.Error())
	}

	fmt.Println("videoName:", r.VideoName)
	fmt.Println("count:", r.Count)

	//connecting to db and getting videos
	conn := rbdata.Connect()
	defer conn.Close(context.Background())
	videos := conn.GetVideos(r.VideoName, r.Count)

	c.JSON(http.StatusOK, gin.H{"videos": videos, "count": len(videos)})
}
