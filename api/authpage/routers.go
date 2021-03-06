package authpage

import (
	"github.com/gin-gonic/gin"
)

// Routers 路由
func Routers(e *gin.Engine) {
	g := e.Group("/authpage")
	g.GET("/componentinfo", getComponentInfoHandler)
	g.GET("/preauthcode", getPreAuthCodeHandler)
}
