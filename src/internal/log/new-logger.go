package log

import (
	"fmt"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func NewLogger(info *LoggerInfo) (Logger, error) {
	if !info.Enabled {
		return zap.NewNop(), nil
	}

	if info.Path == "" {
		return nil, fmt.Errorf("invalid config, value: %q", info.Path)
	}

	ws := zapcore.AddSync(&lumberjack.Logger{
		Filename:   info.Path,
		MaxSize:    info.Rotation.MaxSizeInMb,
		MaxBackups: info.Rotation.MaxNoOfBackups,
		MaxAge:     info.Rotation.MaxAgeInDays,
	})
	config := zap.NewProductionEncoderConfig()
	config.EncodeTime = zapcore.TimeEncoderOfLayout(info.TimeStampFormat)
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(config),
		ws,
		info.Level,
	)

	return zap.New(core), nil
}
