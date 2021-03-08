package logger

import (
    "fmt"
    "gin-example/config"
    "go.uber.org/zap"
    "go.uber.org/zap/zapcore"
    "gopkg.in/natefinch/lumberjack.v2"
    "os"
    "strings"
)

var log *zap.SugaredLogger

func Setup() {
    log = NewLogger();
    fmt.Println("加载logger")
}
func NewLogger() *zap.SugaredLogger{
    encoder := zap.NewProductionEncoderConfig()
    encoder.EncodeTime = zapcore.ISO8601TimeEncoder

    syncWriter := zapcore.AddSync(&lumberjack.Logger{
        Filename:  getFilePath(),
        MaxSize:   config.Logger.MaxSize, // 1G
        MaxAge:    config.Logger.MaxAge,   // 1G
        LocalTime: true,
        Compress:  false,
    })
    // 设置日志级别
    atomicLevel := zap.NewAtomicLevel()
    atomicLevel.SetLevel(getLoggerLevel(config.Logger.AtomicLevel))
    // 创建核心配置
    core := zapcore.NewCore(
        zapcore.NewJSONEncoder(encoder),
        syncWriter,
        atomicLevel,
    )
    // 获取一个日志对象
    zapLog := zap.New(
        core,
        zap.AddCaller(),
        zap.AddCallerSkip(1),
    ).Sugar()
    return zapLog;
}
func getLoggerLevel(lvl string) zapcore.Level {
    var levelMap = map[string]zapcore.Level{
        "debug":  zapcore.DebugLevel,
        "info":   zapcore.InfoLevel,
        "warn":   zapcore.WarnLevel,
        "error":  zapcore.ErrorLevel,
        "dpanic": zapcore.DPanicLevel,
        "panic":  zapcore.PanicLevel,
        "fatal":  zapcore.FatalLevel,
    }
    if level, ok := levelMap[lvl]; ok {
        return level
    }
    return zapcore.DebugLevel
}

func getFilePath() string {

    logfile := config.Logger.FilePath+"/" + getAppname() + ".log"
    return logfile
}

func getAppname() string {
    full := os.Args[0]
    full = strings.Replace(full, "\\", "/", -1)
    splits := strings.Split(full, "/")
    if len(splits) >= 1 {
        name := splits[len(splits)-1]
        name = strings.TrimSuffix(name, ".exe")
        return name
    }

    return ""
}
func Debug(args ...interface{}) {

    log.Debug(args...)

}

func Debugf(template string, args ...interface{}) {

    log.Debugf(template, args...)

}

func Info(args ...interface{}) {

    log.Info(args...)

}

func Infof(template string, args ...interface{}) {

    log.Infof(template, args...)

}

func Warn(args ...interface{}) {

    log.Warn(args...)

}

func Warnf(template string, args ...interface{}) {

    log.Warnf(template, args...)

}

func Error(args ...interface{}) {

    log.Error(args...)

}

func Errorf(template string, args ...interface{}) {

    log.Errorf(template, args...)

}

func DPanic(args ...interface{}) {

    log.DPanic(args...)

}

func DPanicf(template string, args ...interface{}) {

    log.DPanicf(template, args...)

}

func Panic(args ...interface{}) {

    log.Panic(args...)

}

func Panicf(template string, args ...interface{}) {

    log.Panicf(template, args...)
}

func Fatal(args ...interface{}) {

    log.Fatal(args...)

}

func Fatalf(template string, args ...interface{}) {

    log.Fatalf(template, args...)

}
