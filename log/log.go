// Copyright 2016 Travelydia, Inc. All rights reserved.

package log

import (
	"fmt"
	"runtime"
	"strings"

	"github.com/Sirupsen/logrus"
)

const StackTraceSkip = 2

var Logger = logrus.New()

type Fields logrus.Fields

func SetLogLevel(level logrus.Level) {
	Logger.Level = level
}

func SetLogFormatter(formatter logrus.Formatter) {
	Logger.Formatter = formatter
}

// Info logs a message at level Info on the standard logger.
func Info(args ...interface{}) {
	if Logger.Level >= logrus.InfoLevel {
		logger := Logger.WithFields(logrus.Fields{})
		logger.Data["file"] = getFile(StackTraceSkip)
		logger.Info(args...)
	}
}

func Debug(args ...interface{}) {
	if Logger.Level >= logrus.DebugLevel {
		logger := Logger.WithFields(logrus.Fields{})
		logger.Data["file"] = getFile(StackTraceSkip)
		logger.Debug(args)
	}
}

// Warn logs a message at level Warn on the standard logger.
func Warn(args ...interface{}) {
	if Logger.Level >= logrus.WarnLevel {
		logger := Logger.WithFields(logrus.Fields{})
		logger.Data["file"] = getFile(StackTraceSkip)
		logger.Warn(args...)
	}
}

func Error(args ...interface{}) {
	if Logger.Level >= logrus.ErrorLevel {
		logger := Logger.WithFields(logrus.Fields{})
		logger.Data["file"] = getFile(StackTraceSkip)
		logger.Error(args...)
	}
}

// Fatal logs a message at level Fatal on the standard logger.
func Fatal(args ...interface{}) {
	if Logger.Level >= logrus.FatalLevel {
		logger := Logger.WithFields(logrus.Fields{})
		logger.Data["file"] = getFile(StackTraceSkip)
		logger.Fatal(args...)
	}
}

// Panic logs a message at level Panic on the standard logger.
func Panic(args ...interface{}) {
	if Logger.Level >= logrus.PanicLevel {
		logger := Logger.WithFields(logrus.Fields{})
		logger.Data["file"] = getFile(StackTraceSkip)
		logger.Panic(args...)
	}
}

func getFile(skip int) string {
	_, file, line, _ := runtime.Caller(skip)
	return fmt.Sprintf("%s:%d", file[strings.LastIndex(file, "/")+1:], line)
}
