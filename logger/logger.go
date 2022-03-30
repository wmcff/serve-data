package logger

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"go.uber.org/zap"
	"gopkg.in/natefinch/lumberjack.v2"
	"gopkg.in/yaml.v2"
	gormLogger "gorm.io/gorm/logger"
)

type Config struct {
	ZapConfig zap.Config        `json:"zap_config" yaml:"zap_config"`
	LogRotate lumberjack.Logger `json:"log_rotate" yaml:"log_rotate"`
}

type Logger interface {
	GetZapLogger() *zap.SugaredLogger
	LogMode(level gormLogger.LogLevel) gormLogger.Interface
	Info(ctx context.Context, msg string, data ...interface{})
	Warn(ctx context.Context, msg string, data ...interface{})
	Error(ctx context.Context, msg string, data ...interface{})
	Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error)
}

type logger struct {
	Zap *zap.SugaredLogger
}

func NewLogger(sugar *zap.SugaredLogger) Logger {
	return &logger{Zap: sugar}
}

func InitLogger(env string) Logger {
	configYaml, err := ioutil.ReadFile("./zaplogger." + env + ".yml")
	if err != nil {
		fmt.Printf("Failed to read logger configuration: %s", err)
		os.Exit(2)
	}
	var myConfig *Config
	if err = yaml.Unmarshal(configYaml, &myConfig); err != nil {
		fmt.Printf("Failed to read zap logger configuration: %s", err)
		os.Exit(2)
	}
	var zap *zap.Logger
	zap, err = build(myConfig)
	if err != nil {
		fmt.Printf("Failed to compose zap logger : %s", err)
		os.Exit(2)
	}
	sugar := zap.Sugar()
	// set package varriable logger.
	log := NewLogger(sugar)
	log.GetZapLogger().Infof("Success to read zap logger configuration: zaplogger." + env + ".yml")
	_ = zap.Sync()
	return log
}

// GetZapLogger returns zapSugaredLogger
func (log *logger) GetZapLogger() *zap.SugaredLogger {
	return log.Zap
}
