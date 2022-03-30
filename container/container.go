package container

import (
	"github.com/wmcff/zoogeek/config"
	"github.com/wmcff/zoogeek/logger"
	"github.com/wmcff/zoogeek/repository"
)

type Container interface {
	GetRepository() repository.Repository
	GetConfig() *config.Config
	GetLogger() logger.Logger
	GetEnv() string
}

type container struct {
	rep    repository.Repository
	config *config.Config
	logger logger.Logger
	env    string
}

func NewContainer(rep repository.Repository, config *config.Config, logger logger.Logger, env string) Container {
	return &container{rep: rep, config: config, logger: logger, env: env}
}

func (c *container) GetRepository() repository.Repository {
	return c.rep
}

func (c *container) GetConfig() *config.Config {
	return c.config
}

func (c *container) GetLogger() logger.Logger {
	return c.logger
}

func (c *container) GetEnv() string {
	return c.env
}
