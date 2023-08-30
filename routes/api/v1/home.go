package v1

import (
	"cska/db/settings"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Home(c *gin.Context) {
	c.String(http.StatusOK, "Привет, это апи сервер для телеграм бота %v", settings.BotName)
}
