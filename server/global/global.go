package global

import (
	"github.com/DraftTin/gin-blog/config"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

var (
	Config *config.Config
	DB     *gorm.DB
	Log    *logrus.Logger
	Router *gin.Engine
)
