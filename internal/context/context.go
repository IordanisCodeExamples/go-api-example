// Package context provides the context of the service
// where objects are added to the context and can be retrieved from it
package context

import (
	"context"

	internallogger "github.com/junkd0g/go-api-example/internal/logger"
)

type contextKey string

func (c contextKey) String() string {
	return string(c)
}

var (
	contextKeyLogger = contextKey("logger")
)

// AddLoggerToContex adding logger to ctx with the key logger
func AddLoggerToContex(ctx context.Context, logger *internallogger.Logger) context.Context {
	return context.WithValue(ctx, contextKeyLogger, logger)
}

// GetLoggerFromContext return logger from context and if it exist in the context
func GetLoggerFromContext(ctx context.Context) (*internallogger.Logger, bool) {
	logger, ok := ctx.Value(contextKeyLogger).(*internallogger.Logger)
	return logger, ok
}
