// Default package consits of default routes to api on root path on the server "/"
package root

import (
	"cska/db/settings"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Home(c *gin.Context) {
	//#TODO использовать Builder для строк
	c.String(http.StatusOK, "Тут находятся апи для бота %v, например api/v1", settings.BotName)
}
