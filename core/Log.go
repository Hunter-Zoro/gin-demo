package core

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"time"
)

var (
	level   zapcore.Level
	options []zap.Option
)

func InitializeLog() *zap.Logger {
	if ok, _ := utils.PathExists(global.DEMO_CONFIG.Log.RootDir); !ok {
		_ = os.MkdirAll(global.DEMO_CONFIG.Log.RootDir, os.ModePerm)
	}
	setLogLevel()
	if global.DEMO_CONFIG.Log.ShowLine {
		options = append(options, zap.AddCaller(), zap.AddCallerSkip(1))
	}
	return zap.New(getZapCore(), options...)

}

func setLogLevel() {
	switch global.DEMO_CONFIG.Log.Level {
	case "debug":
		level = zap.DebugLevel
		options = append(options, zap.AddStacktrace(level))
	case "info":
		level = zap.InfoLevel
	case "warn":
		level = zap.WarnLevel
	case "error":
		level = zap.ErrorLevel
		options = append(options, zap.AddStacktrace(level))
	case "dpanic":
		level = zap.DPanicLevel
	case "panic":
		level = zap.PanicLevel
	case "fatal":
		level = zap.FatalLevel
	default:
		level = zap.InfoLevel
	}
}

// 扩展 Zap
func getZapCore() zapcore.Core {
	var encoder zapcore.Encoder

	// 调整编码器默认配置
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = func(time time.Time, encoder zapcore.PrimitiveArrayEncoder) {
		encoder.AppendString(time.Format("[" + "2006-01-02 15:04:05.000" + "]"))
	}
	encoderConfig.EncodeLevel = func(l zapcore.Level, encoder zapcore.PrimitiveArrayEncoder) {
		encoder.AppendString(global.DEMO_CONFIG.App.Env + "." + l.String())
	}

	// 设置编码器
	if global.DEMO_CONFIG.Log.Format == "json" {
		encoder = zapcore.NewJSONEncoder(encoderConfig)
	} else {
		encoder = zapcore.NewConsoleEncoder(encoderConfig)
	}

	return zapcore.NewCore(encoder, getLogWriter(), level)
}

// 使用 lumberjack 作为日志写入器
func getLogWriter() zapcore.WriteSyncer {
	file := &lumberjack.Logger{
		Filename:   global.DEMO_CONFIG.Log.RootDir + "/" + global.DEMO_CONFIG.Log.Filename,
		MaxSize:    global.DEMO_CONFIG.Log.MaxSize,
		MaxBackups: global.DEMO_CONFIG.Log.MaxBackups,
		MaxAge:     global.DEMO_CONFIG.Log.MaxAge,
		Compress:   global.DEMO_CONFIG.Log.Compress,
	}

	return zapcore.AddSync(file)
}
