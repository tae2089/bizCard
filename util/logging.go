package util

import (
	"bytes"
	"encoding/json"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"io"
	"os"
	"sync"
)

var Log *zap.Logger
var logOnce sync.Once

func SetupLogging() *zap.Logger {

	if Log == nil {
		logOnce.Do(func() {
			config := zap.NewProductionConfig()
			encoderConfig := zap.NewProductionEncoderConfig()
			encoderConfig.TimeKey = "timestamp"
			encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
			config.EncoderConfig = encoderConfig

			w := zapcore.AddSync(&lumberjack.Logger{
				Filename:   "../logging.log",
				MaxSize:    500, // megabytes
				MaxBackups: 3,
				MaxAge:     28, // days
			})
			writeStdout := zapcore.AddSync(CustomWriter{})

			core := zapcore.NewCore(
				zapcore.NewJSONEncoder(encoderConfig),
				zapcore.NewMultiWriteSyncer(w, writeStdout),
				zap.InfoLevel)

			Log = zap.New(core,
				zap.AddCaller(),
				zap.AddStacktrace(zap.ErrorLevel),
			)
		})
	}
	return Log
}

type CustomWriter struct{}

func (e CustomWriter) Write(p []byte) (int, error) {
	var prettyJSON bytes.Buffer
	if err := json.Indent(&prettyJSON, p, "", "    "); err != nil {
		return 0, err
	}
	n, err := os.Stdout.Write(prettyJSON.Bytes())
	if err != nil {
		return n, err
	}
	if n != len(prettyJSON.Bytes()) {
		return n, io.ErrShortWrite
	}
	return len(p), nil
}
