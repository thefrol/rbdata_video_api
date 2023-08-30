package tests

import (
	"context"
	"cska/db/rbdata"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Проверяет работает ли соединение с базой данных
func Database(c *gin.Context) {
	//connecting to db and getting videos

	conn := rbdata.Connect()
	defer conn.Close(context.Background())
	defer func() {
		err := recover()
		if err != nil {
			c.String(http.StatusOK, "Cant connect to DB")
		}
	}()

	c.String(http.StatusOK, "Connected to DB")
}
