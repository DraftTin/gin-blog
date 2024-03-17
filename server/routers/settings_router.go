package routers

import (
	"github.com/DraftTin/gin-blog/api/settings_api"
)

func (router RouterGroup) SettingsRouter() {
	api := settings_api.SettingsApi{}
	router.GET("/", api.SettingInfoView)
}
