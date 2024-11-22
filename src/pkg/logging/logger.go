package logging

import "github.com/Mahdi-ak/golang-carshop-api/src/config"

type Logger interface {
	Init()

	Info(cat Category, sub SubCategory, msg string, extra map[ExtraKey]interface{})
	Infof(template string, args ...interface{})

	Debug(cat Category, sub SubCategory, msg string, extra map[ExtraKey]interface{})
	Debugf(template string, args ...interface{})

	Warn(cat Category, sub SubCategory, msg string, extra map[ExtraKey]interface{})
	Warnf(template string, args ...interface{})

	fatal(err error, cat Category, sub SubCategory, msg string, extra map[ExtraKey]interface{})
	fatalf(err error, template string, args ...interface{})

	Error(err error, cat Category, sub SubCategory, msg string, extra map[ExtraKey]interface{})
	Errorf(err error, template string, args ...interface{})
}

func NewLogger(cfg *config.Config) Logger {
	return nil
}
