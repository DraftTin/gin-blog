package routers

import (
	"github.com/DraftTin/gin-blog/api/settings_api"
	"github.com/DraftTin/gin-blog/global"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	gin.SetMode(global.Config.System.Env)
	basicRouter := gin.Default()
	settingsApi := settings_api.SettingsApi{}
	apiGroup := basicRouter.Group("api")
	{
		apiGroup.GET("setting_info", settingsApi.SettingInfoView)
	}

	return basicRouter
}
