package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Field is an alias for zap.Field used by the internal logging package.
type Field = zap.Field

// Level is an alias for zapcore.Level representing the logging level.
type Level = zapcore.Level

const (
	DebugLevel = zapcore.DebugLevel
	InfoLevel  = zapcore.InfoLevel
	WarnLevel  = zapcore.WarnLevel
	ErrorLevel = zapcore.ErrorLevel
)

// Logger describes the logging contract used internally by arcadia.
type Logger interface {
	Debug(msg string, fields ...Field)
	Info(msg string, fields ...Field)
	Warn(msg string, fields ...Field)
	Error(msg string, fields ...Field)
	Sync() error
}

// Rotation captures file-based log rotation settings such as size,
// number of backups and age.
type Rotation struct {
	// Filename of the log file.
	Filename string
	// MaxSizeInMb specifies the maximum size in megabytes before rotation.
	MaxSizeInMb int
	// MaxNoOfBackups specifies the maximum number of old log files to retain.
	MaxNoOfBackups int
	// MaxAgeInDays specifies the maximum number of days to retain old log files.
	MaxAgeInDays int
}

// LoggerInfo groups the configuration needed to construct a Logger,
// including rotation, output path, time format and level.
type LoggerInfo struct {
	Rotation

	// Enabled indicates whether the logger is initially enabled.
	Enabled bool
	// Path to the log file.
	Path string
	// TimeStampFormat specifies the format for timestamps.
	TimeStampFormat string
	// Level specifies the logging level.
	Level Level
}
