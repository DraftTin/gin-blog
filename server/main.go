package main

import (
	"fmt"

	"github.com/DraftTin/gin-blog/core"
	"github.com/DraftTin/gin-blog/global"
)

func main() {
	// Read conf file
	core.InitConf()
	// 连接数据库
	fmt.Println(global.Config)
}
