package zap

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"time"
)

var logger *zap.Logger

func Setup(opts ...Option) {
	opt := newOption(opts...)

	encoderConfig := zapcore.EncoderConfig{
		TimeKey:       "datetime",
		LevelKey:      "level",
		NameKey:       "logger",
		CallerKey:     "caller",
		MessageKey:    "message",
		StacktraceKey: "stacktrace",
		LineEnding:    zapcore.DefaultLineEnding,
		EncodeLevel:   zapcore.LowercaseLevelEncoder, // 小写编码器
		EncodeTime: func(t time.Time, encoder zapcore.PrimitiveArrayEncoder) {
			encoder.AppendString(t.Format("2006-01-02 15:04:05"))
		},
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.FullCallerEncoder, // 全路径编码器
		EncodeName:     zapcore.FullNameEncoder,
	}

	// 设置日志级别区间
	atomicLevel := zap.NewAtomicLevel()
	switch opt.LogLevel {
	case "debug":
		atomicLevel.SetLevel(zap.DebugLevel)
	case "info":
		atomicLevel.SetLevel(zap.InfoLevel)
	case "warn":
		atomicLevel.SetLevel(zap.WarnLevel)
	case "error":
		atomicLevel.SetLevel(zap.ErrorLevel)
	default:
		atomicLevel.SetLevel(zap.ErrorLevel)
	}
	// 指定日志等级
	//atomicLevel = zap.LevelEnablerFunc(func(level zapcore.Level) bool {
	//	return level == zap.ErrorLevel
	//})

	var ws zapcore.WriteSyncer
	if opt.LogPath != "" {
		hook := lumberjack.Logger{
			Filename:   opt.LogPath,        // 日志文件路径
			MaxSize:    opt.FileMaxSize,    // 每个日志文件保存的最大尺寸 单位：M
			MaxBackups: opt.FileMaxBackups, // 日志文件最多保存多少个备份
			MaxAge:     opt.FileMaxAge,     // 文件最多保存多少天
			Compress:   opt.FileCompress,   // 是否压缩
		}
		ws = zapcore.NewMultiWriteSyncer(zapcore.AddSync(&hook))
	} else {
		ws = os.Stdout
	}

	// 一个 NewCore 代表一个日志文件；若需要多个日志文件，则需要多个 NewCore
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig), // 编码器配置
		ws,
		atomicLevel, // 日志级别
	)

	logger = zap.New(core, zap.AddCaller())
}

func Logger() *zap.Logger {
	return logger
}
