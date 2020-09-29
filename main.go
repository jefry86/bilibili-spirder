//@Description todo
//@Author 凌云  jefry52@gmail.com
//@DateTime 2020/9/15 12:25 下午

package main

import (
	"bilibili-spirder/global"
	"bilibili-spirder/services"
	"bilibili-spirder/utils"
	"os"
	"os/signal"
)

func main() {
	utils.LoadCfg()
	global.NewGlobal()
	global.Logger.Info("服务启动中...")
	go services.SpiderVideo()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	global.Logger.Info("服务已停止")

}
