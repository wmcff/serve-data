package main

import (
	"github.com/wmcff/zoogeek/router"

	"github.com/wmcff/zoogeek/container"

	"github.com/wmcff/zoogeek/repository"

	"github.com/wmcff/zoogeek/logger"

	"github.com/wmcff/zoogeek/config"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	conf, env := config.Load()

	//初始化日志配置
	logger := logger.InitLogger(env)
	logger.GetZapLogger().Infof("Loaded this configuration : application." + env + ".yml")

	//服务池
	rep := repository.NewZooRepository(logger, conf)
	container := container.NewContainer(rep, conf, logger, env)

	//路由
	router.Init(e, container)

	//服务接收
	if err := e.Start(":8080"); err != nil {
		logger.GetZapLogger().Errorf(err.Error())
	}
}
