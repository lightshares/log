package log

import (
	"fmt"
	"github.com/spf13/viper"
)

/**
配置信息
*/
type logConfig struct {
	FileName   string     // 日志文件名称
	FilePath   string     // 日志路径
	Level      string     // 日志级别
	MaxSize    int        // 单个文件大小
	Compress   bool       // true表示压缩成zip文件,默认false
	MaxBackups int        // 保存几个历史文件
	Type       outputType // console表示是控制台输出,file表示是文件输出
}

/**
初始化配置分为两个部分
1. 加载配置问津。
2. 设置默认值
*/
func initConfig() (*logConfig, error) {
	logConfig, err := loadConfig()
	if err != nil {
		return nil, err
	}
	setDefaultConfig(logConfig)
	return logConfig, nil
}

/**
加载配置文件，文件路径在当前项目下的etc目录下的log.yaml中
*/
func loadConfig() (*logConfig, error) {
	v := viper.New()
	v.AddConfigPath("./../etc")
	v.AddConfigPath("etc")
	v.SetConfigName("log")
	v.SetConfigType("yaml")
	err := v.ReadInConfig()
	if err != nil {
		fmt.Println("在etc目录下没有log.yaml文件使用默认配置")
		return &logConfig{}, nil
	}
	logV := v.Get("log")
	if logV == nil {
		fmt.Println("log.yaml文件中没有log字段使用默认配置")
		return &logConfig{}, nil
	}
	config := logConfig{}
	err = v.Sub("log").Unmarshal(&config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}
func setDefaultConfig(config *logConfig) {
	if config.FileName == "" {
		config.FileName = "app.log"
	}
	if config.FilePath == "" {
		config.FilePath = "logs"
	}
	if config.Level == "" {
		config.Level = "debug"
	}
	if config.MaxBackups == 0 {
		config.MaxBackups = 10
	}
	if config.MaxSize == 0 {
		config.MaxSize = 500 * 1024 * 1024
	}
	if config.Type == "" {
		config.Type = ConsoleOutput
	}
}

type outputType string

var ConsoleOutput outputType = "console"
