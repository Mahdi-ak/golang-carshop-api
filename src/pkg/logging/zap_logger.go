package logging

import (
	"github.com/Mahdi-ak/golang-carshop-api/src/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var zapSinLogger *zap.SugaredLogger

type zapLogger struct {
	cfg    *config.Config
	logger *zap.SugaredLogger
}

var logLevelMap = map[string]zapcore.Level{
	"debug": zap.DebugLevel,
	"info":  zap.InfoLevel,
	"warn":  zap.WarnLevel,
	"error": zap.ErrorLevel,
	"fatal": zap.FatalLevel,
}

// Init initializes the zap logger with the configuration settings.

func newZapLogger(cfg *config.Config) *zapLogger {
	logger := &zapLogger{
		cfg: cfg,
	}
	logger.Init()
	return logger

}

func (l *zapLogger) GetLogLevel() zapcore.Level {

	level, exists := logLevelMap[l.cfg.Logger.Level]
	if !exists {
		level = zap.DebugLevel
	}
	return level

}

func (l *zapLogger) Init() {
	once.Do(func() {

		w := zapcore.AddSync(&lumberjack.Logger{
			// rotate log files every 50MB and keep a max of 10 days
			Filename:   l.cfg.Logger.FilePath,
			MaxSize:    50,    // megabytes
			MaxAge:     10,    // days
			LocalTime:  false, // use UTC time
			MaxBackups: 10,    // keep 10 backup logs
			Compress:   true,
		})

		config := zap.NewProductionEncoderConfig()
		config.EncodeTime = zapcore.ISO8601TimeEncoder

		core := zapcore.NewCore(
			zapcore.NewJSONEncoder(config),
			w,
			l.GetLogLevel(),
		)

		logger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1),
			zap.AddStacktrace(zapcore.ErrorLevel)).Sugar()

		zapSinLogger = logger.With("AppName", "MyApp", "LoggerName", "Zaplog")
	})

	l.logger = zapSinLogger
}

func (l *zapLogger) Debug(cat Category, sub SubCategory, msg string, extra map[ExtraKey]interface{}) {
	params := prepareLogKeys(extra, cat, sub)

	l.logger.Debugw(msg, params...)
}

func (l *zapLogger) Debugf(template string, args ...interface{}) {
	l.logger.Debugf(template, args)
}

func (l *zapLogger) Info(cat Category, sub SubCategory, msg string, extra map[ExtraKey]interface{}) {
	params := prepareLogKeys(extra, cat, sub)

	l.logger.Infow(msg, params...)
}

func (l *zapLogger) Infof(template string, args ...interface{}) {
	l.logger.Infof(template, args)
}
func (l *zapLogger) Warn(cat Category, sub SubCategory, msg string, extra map[ExtraKey]interface{}) {
	params := prepareLogKeys(extra, cat, sub)

	l.logger.Warnw(msg, params...)
}

func (l *zapLogger) Warnf(template string, args ...interface{}) {
	l.logger.Warnf(template, args)
}

func (l *zapLogger) Error(cat Category, sub SubCategory, msg string, extra map[ExtraKey]interface{}) {
	params := prepareLogKeys(extra, cat, sub)

	l.logger.Errorw(msg, params...)
}

func (l *zapLogger) Errorf(template string, args ...interface{}) {
	l.logger.Errorf(template, args)
}
func (l *zapLogger) Fatal(cat Category, sub SubCategory, msg string, extra map[ExtraKey]interface{}) {
	params := prepareLogKeys(extra, cat, sub)

	l.logger.Fatalw(msg, params...)
}

func (l *zapLogger) Fatalf(template string, args ...interface{}) {
	l.logger.Fatalf(template, args)
}

func prepareLogKeys(extra map[ExtraKey]interface{}, cat Category, sub SubCategory) []interface{} {
	if extra == nil {
		extra = make(map[ExtraKey]interface{}, 0)
	}
	extra["Category"] = cat
	extra["SubCategory"] = sub
	params := mapToZapParams(extra)
	return params
}
