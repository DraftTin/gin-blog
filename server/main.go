package main

import (
	"github.com/DraftTin/gin-blog/core"
	"github.com/DraftTin/gin-blog/global"
	"github.com/DraftTin/gin-blog/routers"
)

func main() {
	// Read conf file
	core.InitConf()
	// 连接数据库
	global.DB = core.InitGorm()
	global.Log = core.InitLogger()
	global.Router = routers.InitRouter()
	global.Log.Infof("gin-blog is runnning at %s:%s", global.Config.System.Host, global.Config.System.Port)
	err := global.Router.Run(global.Config.System.GetAddr())
	if err != nil {
		global.Log.Errorf("router error: %s", err)
	}
}
