package main

import (
	"github.com/wmcff/zoogeek/config"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	conf := config.Load()

	//初始化配置
	logger.Init(&conf.Logger)
	//服务池

	//服务接收
	if err := e.Start(":8080"); err != nil {
		logger.GetZapLogger().Errorf(err.Error())
	}
}
