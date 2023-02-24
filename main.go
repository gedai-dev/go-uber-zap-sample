package main

import (
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {

	// This constant values will be used in zap.Config bellow
	const (
		timeKey          string = "Timestamp"
		nameKey          string = "Key"
		levelKey         string = "Severity"
		encoding         string = "json"
		CallerKey        string = "Caller"
		messageKey       string = "Body"
		outputPaths      string = "stdout"
		stacktraceKey    string = "Trace"
		errorOutputPaths string = "stderr"
	)

	// This Variable will be used to write logs stdout
	var logger *zap.SugaredLogger

	// Structuring our log
	zapConfig := &zap.Config{
		Development: false,
		Level:       zap.NewAtomicLevelAt(zap.InfoLevel),
		Encoding:    encoding,
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey:     messageKey,
			LevelKey:       levelKey,
			TimeKey:        timeKey,
			NameKey:        nameKey,
			CallerKey:      CallerKey,
			FunctionKey:    zapcore.OmitKey,
			StacktraceKey:  stacktraceKey,
			SkipLineEnding: false,
			LineEnding:     zapcore.DefaultLineEnding,
			EncodeLevel:    zapcore.CapitalLevelEncoder,
			EncodeTime:     zapcore.RFC3339TimeEncoder,
			EncodeDuration: zapcore.MillisDurationEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder,
			EncodeName:     zapcore.FullNameEncoder,
		},
		OutputPaths:      []string{outputPaths},
		ErrorOutputPaths: []string{errorOutputPaths},
		InitialFields: map[string]interface{}{
			"Attributes": map[string]interface{}{
				"service.name":    "gedai",
				"service.version": "v1.0.0",
				"time":            time.Now().UTC(),
			},
			"Annotations": map[string]interface{}{
				"team":     "team-gedai",
				"contact":  "gedai-contact",
				"handbook": "http://handbook.io",
			},
		},
	}

	// 	Config offers a declarative way to construct a logger. It doesn't do anything that can't be done with New, Options, and the various zapcore.WriteSyncer and zapcore.Core wrappers, but it's a simpler way to toggle common options.
	// Note that Config intentionally supports only the most common options. More unusual logging setups (logging to network connections or message queues, splitting output between multiple files, etc.) are possible, but require direct use of the zapcore package. For sample code, see the package-level BasicConfiguration and AdvancedConfiguration examples.
	// For an example showing runtime log level changes, see the documentation for AtomicLevel.
	_logger, err := zapConfig.Build(zap.AddCaller(), zap.AddCallerSkip(1))
	if err != nil {
		panic(err)
	}

	logger = _logger.Sugar()

	logger.Info("Hi there! I`m a log kind of INFO ...")

	logger.Warn("Hey pay attention! I`m a log kind of WARN :o ...")

	logger.Warn("OH NO! I`m a log kind of ERR :o ...")

}
