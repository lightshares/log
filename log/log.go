package log

import (
	"github.com/lightshares/log/trace"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

// error logger
var log *zap.SugaredLogger
var enableTrace bool //默认是false表示不开启trace功能

func SetEnableTrace(b bool) {
	enableTrace = b
}

var levelMap = map[string]zapcore.Level{
	"debug":  zapcore.DebugLevel,
	"info":   zapcore.InfoLevel,
	"warn":   zapcore.WarnLevel,
	"error":  zapcore.ErrorLevel,
	"dPanic": zapcore.DPanicLevel,
	"panic":  zapcore.PanicLevel,
	"fatal":  zapcore.FatalLevel,
}

func getLoggerLevel(lvl string) zapcore.Level {
	if level, ok := levelMap[lvl]; ok {
		return level
	}
	return zapcore.InfoLevel
}

/**
初始化log实例
*/
func init() {
	config, err := initConfig()
	if err != nil {
		panic(err)
	}
	if config.Type == ConsoleOutput {
		prodLog, err := zap.NewProduction()
		if err != nil {
			panic(err)
		}
		log = prodLog.Sugar()
		return
	}
	fileName := config.FilePath + config.FileName
	level := getLoggerLevel(config.Level)
	syncWriter := zapcore.AddSync(&lumberjack.Logger{
		Filename:   fileName,
		MaxSize:    config.MaxSize, //1G
		LocalTime:  true,
		Compress:   config.Compress,
		MaxBackups: config.MaxBackups,
	})
	encoder := zap.NewProductionEncoderConfig()
	encoder.EncodeTime = zapcore.ISO8601TimeEncoder
	core := zapcore.NewCore(zapcore.NewJSONEncoder(encoder), syncWriter, zap.NewAtomicLevelAt(level))
	logger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
	log = logger.Sugar()
}

func getLog() *zap.SugaredLogger {
	//是否开启trace
	if enableTrace {
		_, traceId := trace.GetTraceId()
		return log.With("traceId", traceId)
	}
	return log
}

func Debug(args ...interface{}) {
	getLog().Debug(args...)
}

func Debugf(template string, args ...interface{}) {
	getLog().Debugf(template, args...)
}

func Info(args ...interface{}) {
	getLog().Info(args...)
}

func Infof(template string, args ...interface{}) {
	getLog().Infof(template, args...)
}

func Warn(args ...interface{}) {
	getLog().Warn(args...)
}

func Warnf(template string, args ...interface{}) {
	getLog().Warnf(template, args...)
}

func Error(args ...interface{}) {
	getLog().Error(args...)
}

func Errorf(template string, args ...interface{}) {
	getLog().Errorf(template, args...)
}

func DPanic(args ...interface{}) {
	getLog().DPanic(args...)
}

func DPanicf(template string, args ...interface{}) {
	getLog().DPanicf(template, args...)
}

func Panic(args ...interface{}) {
	getLog().Panic(args...)
}

func Panicf(template string, args ...interface{}) {
	getLog().Panicf(template, args...)
}

func Fatal(args ...interface{}) {
	getLog().Fatal(args...)
}

func Fatalf(template string, args ...interface{}) {
	getLog().Fatalf(template, args...)
}
