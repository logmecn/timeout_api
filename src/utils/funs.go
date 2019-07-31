package utils

import (
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
)

/*
获取当前运行程序的目录，用来拼接配置文件的地址
*/
func GetCurrentDirectory() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Panic(err)
	}
	return strings.Replace(dir, "\\", "/", -1)
}

// 日志格式初始化
func InitLogger(logPath string, MaxSize int, logLever string) *zap.Logger {
	hook := lumberjack.Logger{
		Filename:   logPath, // 日志文件路径
		MaxSize:    MaxSize, // 每个日志文件保存的最大尺寸 单位：M
		MaxBackups: 1000,    // 日志文件最多保存多少个备份
		MaxAge:     1000,    // 文件最多保存多少天
		Compress:   true,    // 是否压缩
	}
	w := zapcore.AddSync(&hook)
	var (
		level zapcore.Level
	)
	switch logLever {
	case "debug":
		level = zap.DebugLevel
	case "info":
		level = zap.InfoLevel
	case "error":
		level = zap.ErrorLevel
	default:
		level = zap.InfoLevel
	}
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = MyDefineTimeEncoder
	core := zapcore.NewCore(
		zapcore.NewConsoleEncoder(encoderConfig),
		w,
		level,
	)
	logger := zap.New(core)
	logger.Info("Loger Init Success")
	return logger
}

// 自定义格式的日期时间。标准的 iso 格式看不习惯
func MyDefineTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02.15:04:05.000"))
}
