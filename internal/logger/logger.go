package internallogger

import (
	"fmt"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Logger represents a customized logger instance.
type Logger struct {
	logger *zap.Logger
}

type LogField map[string]interface{}

// NewLogger creates a new Logger instance.
func NewLogger() (*Logger, error) {
	config := zap.NewDevelopmentConfig()
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder

	logger, err := config.Build()
	if err != nil {
		return nil, fmt.Errorf("failed to create logger: %w", err)
	}

	return &Logger{
		logger: logger,
	}, nil
}

// SetCore sets the core of the logger, useful for testing.
func (l *Logger) SetCore(core zapcore.Core) {
	l.logger = zap.New(core)
}

// Info logs a message at the info level.
func (l *Logger) Info(msg string, fields ...LogField) {
	zapFields := convertToZapFields(fields...)
	l.logger.Info(msg, zapFields...)
}

// Error logs a message at the error level.
func (l *Logger) Error(msg string, fields ...LogField) {
	zapFields := convertToZapFields(fields...)
	l.logger.Error(msg, zapFields...)
}

// convertToZapFields coverts our custom LogField type to an array of zap.Field
func convertToZapFields(fields ...LogField) []zap.Field {
	var zapFields []zap.Field

	for _, field := range fields {
		for k, v := range field {
			switch value := v.(type) {
			case string:
				zapFields = append(zapFields, zap.String(k, value))
			case int:
				zapFields = append(zapFields, zap.Int(k, value))
				// ... Add other types as needed
			}
		}
	}

	return zapFields
}
