package main

import (
	"github.com/wmcff/serve-data/router"

	"github.com/wmcff/serve-data/container"

	"github.com/wmcff/serve-data/repository"

	"github.com/wmcff/serve-data/logger"

	"github.com/wmcff/serve-data/config"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	conf, env := config.Load()

	//初始化日志配置
	logger := logger.InitLogger(env)
	logger.GetZapLogger().Infof("Loaded this configuration : conf." + env + ".yml")

	//服务池
	rep := repository.NewObjRepository(logger, conf)
	container := container.NewContainer(rep, conf, logger, env)

	//路由
	router.Init(e, container)

	//服务接收
	if err := e.Start(":8081"); err != nil {
		logger.GetZapLogger().Errorf(err.Error())
	}
}
