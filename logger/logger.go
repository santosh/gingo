package logger

import "go.uber.org/zap"

var Log *logger
var err error

type logger struct {
	*zap.Logger
}

func (l *logger) init() error {
	// if config.IsProdEnv() {
	// 	l.Logger, err = zap.NewProduction()
	// } else {
	// 	l.Logger, err = zap.NewDevelopment()
	// }

	l.Logger, err = zap.NewDevelopment()

	l.Logger = l.Logger.With(
		zap.String("service", "gingo"),
		zap.String("version", "0.1.0"),
	)

	return err
}

func init() {
	Log = &logger{}
	err = Log.init()
	if err != nil {
		panic(err)
	}
}
