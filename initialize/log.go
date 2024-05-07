package initialize

import (
	"github.com/ryze2048/kafka-example/global"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

func InitLog() {
	// 定义颜色配置
	colorCfg := zapcore.EncoderConfig{
		MessageKey:     "message",
		LevelKey:       "level",
		TimeKey:        "time",
		NameKey:        "logger",
		CallerKey:      "caller",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalColorLevelEncoder, // 设置日志级别的颜色
		EncodeTime:     zapcore.ISO8601TimeEncoder,       // 使用 ISO8601 格式输出时间
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	// 创建 ConsoleEncoder 实例
	consoleEncoder := zapcore.NewConsoleEncoder(colorCfg)

	// 创建一个日志级别
	level := zap.NewAtomicLevel()

	// 设置日志级别
	level.SetLevel(zapcore.DebugLevel)

	// 创建 Core 实例
	core := zapcore.NewCore(
		consoleEncoder,
		zapcore.Lock(zapcore.AddSync(os.Stderr)),
		level,
	)

	// 创建 Logger 实例
	global.ZAPLOG = zap.New(core)

}
