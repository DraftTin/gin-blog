package settings_api

import (
	"github.com/DraftTin/gin-blog/models/response"
	"github.com/gin-gonic/gin"
)

func (SettingsApi) SettingInfoView(c *gin.Context) {
	response.FailWithCode(response.SettingError, c)
}
