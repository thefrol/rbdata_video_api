// Tests package consists of routes for testing
package tests

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Home(c *gin.Context) {
	//#TODO использовать Builder для строк
	c.String(http.StatusOK, "Это отделение тестирования \nвоспользуйтесь %vfind чтобы попробовать достать несколько видео\n%vdabase чтобы проверить как работает база данных", c.FullPath(), c.FullPath())
}
