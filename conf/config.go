package config

import (
	"flag"
	"fmt"
	"os"

	"github.com/jinzhu/configor"
)

type Config struct {
	Database       Database
	Redis          Redis
	Extension      Extension
	Log            Log
	StaticContents StaticContents
	Security       Security
}

type Database struct {
	Host     string // 主机名
	Port     int    // 端口
	Database string // 数据库名
	Username string // 用户名
	Password string // 密码
	Charset  string // 字符集
	Debug    bool   //调试模式
	MaxOpen  int32  //最大连接数
	MaxIdle  int32  // 最大空闲数
}
type Redis struct {
	Enabled            bool `default:"false"`
	ConnectionPoolSize int  `yaml:"connection_pool_size" default:"10"`
	Host               string
	Port               string
}

type Extension struct {
	MasterGenerator bool `yaml:"master_generator" default:"false"`
	CorsEnabled     bool `yaml:"cors_enabled" default:"false"`
	SecurityEnabled bool `yaml:"security_enabled" default:"false"`
}
type Log struct {
	RequestLogFormat string `yaml:"request_log_format" default:"${remote_ip} ${account_name} ${uri} ${method} ${status}"`
}
type StaticContents struct {
	Path string `yaml:"path"`
}
type Security struct {
	AuthPath    []string `yaml:"auth_path"`
	ExculdePath []string `yaml:"exclude_path"`
	UserPath    []string `yaml:"user_path"`
	AdminPath   []string `yaml:"admin_path"`
}

const (
	//environment
	DEV     = "dev"
	TEST    = "test"
	SANDBOX = "sandbox"
	PRD     = "prod"
)

func Load(*Config, string) (config *Config) {
	var env *string
	if value := os.Getenv("APP_ENV"); value != "" {
		env = &value
	} else {
		env = flag.String("env", "dev", "To switch configurations.")
		flag.Parse()
	}

	config = &Config{}
	if err := configor.Load(config, *env+".yml"); err != nil {
		fmt.Printf("Failed to read %s.yml: %s", *env, err)
		os.Exit(2)
	}
	return config
}
