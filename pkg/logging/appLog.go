package logging

import (
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var AppLog *zap.SugaredLogger

func init() {
	var level zapcore.Level = zapcore.DebugLevel

	encoderConfig := zapcore.EncoderConfig{
		TimeKey:       "time",
		LevelKey:      "level",
		NameKey:       "logger",
		CallerKey:     "caller",
		MessageKey:    "msg",
		StacktraceKey: "stacktrace",
		LineEnding:    zapcore.DefaultLineEnding,
		EncodeLevel:   zapcore.CapitalColorLevelEncoder,
		EncodeTime: func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendString(t.Format("2006-01-02 15:04:05"))
		},
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	// 设置日志级别
	atom := zap.NewAtomicLevelAt(level)

	logConfig := zap.Config{
		Level:             atom,               // 日志级别
		Development:       false,              // 开发模式,堆栈跟踪
		DisableStacktrace: true,               // 关闭自动堆栈捕获
		Encoding:          "console",          // 输出格式 console 或 json
		EncoderConfig:     encoderConfig,      // 编码器配置
		InitialFields:     nil,                // 初始化字段,如：添加一个服务器名称
		OutputPaths:       []string{"stdout"}, // 输出到指定文件 stdout（标准输出,正常颜色） stderr（错误输出,红色）
		ErrorOutputPaths:  []string{"stderr"},
	}

	// 构建日志
	logger, _ := logConfig.Build()
	AppLog = logger.Sugar()
}
