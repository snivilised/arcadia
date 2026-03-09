package log

import (
	"github.com/snivilised/arcadia/src/internal/log"
)

// Logger defines the logging interface used by the arcadia application,
// delegating to the internal logging implementation.
type Logger interface {
	Debug(msg string, fields ...log.Field)
	Info(msg string, fields ...log.Field)
	Warn(msg string, fields ...log.Field)
	Error(msg string, fields ...log.Field)
}
