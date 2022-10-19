package loggers

import (
	"errors"
	"github.com/gunjdesai/kafka-connect-connector-manager/constants"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger struct {
	log *zap.Logger
}

func New(lvl string) (*Logger, error) {

	Log, err := load(lvl)

	if err != nil {
		err = errors.New("Error during Logger Initialisation " + err.Error())
	}

	return Log, err

}

func load(level string) (*Logger, error) {

	log := Logger{}
	lvl := getLogLevel(level)

	if err := log.build(lvl); err != nil {
		return &log, err
	}

	return &log, nil

}

func (l *Logger) build(lvl zapcore.Level) error {

	cfg := zap.Config{
		Encoding:    "json",
		Level:       zap.NewAtomicLevelAt(lvl),
		OutputPaths: []string{"stdout"},
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey:  "message",
			LevelKey:    "level",
			EncodeLevel: zapcore.CapitalLevelEncoder,
			TimeKey:     "time",
			EncodeTime:  zapcore.EpochMillisTimeEncoder,
		},
	}

	logger, err := cfg.Build()

	if err != nil {
		return err
	}

	defer logger.Sync()
	l.log = logger

	return nil

}

func (l *Logger) Debug(message string, fields ...zap.Field) {

	go l.log.Debug(message, fields...)

}

func (l *Logger) Info(message string, fields ...zap.Field) {

	go l.log.Info(message, fields...)

}

func (l *Logger) Warn(message string, fields ...zap.Field) {

	go l.log.Warn(message, fields...)

}

func (l *Logger) Fatal(message string, fields ...zap.Field) {

	go l.log.Fatal(message, fields...)

}

func (l *Logger) Panic(message string, fields ...zap.Field) {

	go l.log.Panic(message, fields...)

}

func getLogLevel(level string) zapcore.Level {

	var lvl zapcore.Level

	switch level {

	case constants.WARN_LOG, constants.WARNING_LOG:
		lvl = zap.WarnLevel
	case constants.FATAL_LOG:
		lvl = zap.FatalLevel
	case constants.INFO_LOG:
		lvl = zap.InfoLevel
	case constants.DEBUG_LOG:
		lvl = zap.DebugLevel
	case constants.PANIC_LOG:
		lvl = zap.PanicLevel
	default:
		lvl = zap.InfoLevel

	}

	return lvl

}
