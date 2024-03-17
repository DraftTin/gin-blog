package routers

import (
	"github.com/DraftTin/gin-blog/global"
	"github.com/gin-gonic/gin"
)

type RouterGroup struct {
	*gin.Engine
}

func InitRouter() *gin.Engine {
	gin.SetMode(global.Config.System.Env)
	basicRouter := gin.Default()
	routerGroup := RouterGroup{Engine: basicRouter}
	routerGroup.SettingsRouter()
	return basicRouter
}
