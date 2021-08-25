package logger

import (
	"context"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"honnef.co/go/tools/config"
	"os"
)

var Module = fx.Provide(NewLogger)

type Params struct {
	fx.In
	fx.Lifecycle
}

type ILogger interface {
	Debug(format string, v ...interface{})
	Info(format string, v ...interface{})
	Warning(format string, v ...interface{})
	Error(format string, v ...interface{})
}

// NewLogger constructs a new logger.
func NewLogger(params Params) ILogger {

	stdoutSyncer := zapcore.Lock(os.Stdout)
	core := zapcore.NewTee(
		zapcore.NewCore(
			zapcore.NewConsoleEncoder(zap.NewProductionEncoderConfig()),
			stdoutSyncer,
			zapcore.DebugLevel,
		),
	)

	log := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
	params.Lifecycle.Append(
		fx.Hook{
			OnStop: func(ctx context.Context) error {
				_ = log.Sync()
				return nil
			},
		},
	)

	return &logger{lg: log.Sugar()}
}

type logger struct {
	lg     *zap.SugaredLogger
	config config.Config
}

func (l *logger) Debug(format string, v ...interface{}) {
	l.lg.Debugf(format, v...)
}

func (l *logger) Info(format string, v ...interface{}) {
	l.lg.Infof(format, v...)
}

func (l *logger) Warning(format string, v ...interface{}) {
	l.lg.Warnf(format, v...)
}

func (l *logger) Error(format string, v ...interface{}) {
	l.lg.Errorf(format, v...)
}
