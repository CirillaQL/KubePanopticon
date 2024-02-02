package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

type Logger struct {
	*zap.Logger
}

var Log *Logger

func InitLogger() {
	//encoderConfig := zapcore.EncoderConfig{
	//	CallerKey:   "caller",
	//	MessageKey:  "msg",
	//	TimeKey:     "ts",
	//	EncodeTime:  zapcore.ISO8601TimeEncoder,
	//	FunctionKey: zapcore.OmitKey,
	//}
	encoder := zapcore.NewJSONEncoder(zapcore.EncoderConfig{
		MessageKey:    "msg",
		LevelKey:      "level",
		TimeKey:       "ts",
		CallerKey:     "caller",
		StacktraceKey: "trace",
		LineEnding:    zapcore.DefaultLineEnding,
		EncodeLevel:   zapcore.LowercaseLevelEncoder,
		EncodeTime:    zapcore.RFC3339TimeEncoder,
		EncodeCaller:  zapcore.ShortCallerEncoder,
	})
	core := zapcore.NewCore(
		encoder,
		zapcore.AddSync(os.Stdout),
		zap.DebugLevel)

	logEntry := zap.New(core, zap.AddCaller(), zap.AddStacktrace(zap.ErrorLevel))
	defer logEntry.Sync()
	Log = &Logger{logEntry}
}

func init() {
	InitLogger()
}
