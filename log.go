/*
 * Copyright 2011 Nan Deng
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

package log

import (
	"io"
	"log"
)

const (
	LOGLEVEL_FATAL = iota
	LOGLEVEL_ALERT
	LOGLEVEL_ERROR
	LOGLEVEL_WARN
	LOGLEVEL_CONFIG
	LOGLEVEL_INFO
	LOGLEVEL_DEBUG
	NR_LOGLEVELS
)

type nullWriter struct{}

func (f *nullWriter) Write(p []byte) (int, error) {
	return len(p), nil
}

type Logger interface {
	Debug(v ...interface{})
	Debugf(format string, v ...interface{})
	Info(v ...interface{})
	Infof(format string, v ...interface{})
	Config(v ...interface{})
	Configf(format string, v ...interface{})
	Warn(v ...interface{})
	Warnf(format string, v ...interface{})
	Error(v ...interface{})
	Errorf(format string, v ...interface{})
	Alert(v ...interface{})
	Alertf(format string, v ...interface{})
	Fatal(v ...interface{})
	Fatalf(format string, v ...interface{})
}

type logger struct {
	logLevel int
	loggers  []*log.Logger
	prefix   string
	writer   io.Writer
}

func (l *logger) Debug(v ...interface{}) {
	l.loggers[LOGLEVEL_DEBUG].Print(v...)
}

func (l *logger) Debugf(format string, v ...interface{}) {
	l.loggers[LOGLEVEL_DEBUG].Printf(format, v...)
}

func (l *logger) Info(v ...interface{}) {
	l.loggers[LOGLEVEL_INFO].Print(v...)
}

func (l *logger) Infof(format string, v ...interface{}) {
	l.loggers[LOGLEVEL_INFO].Printf(format, v...)
}

func (l *logger) Config(v ...interface{}) {
	l.loggers[LOGLEVEL_CONFIG].Print(v...)
}

func (l *logger) Configf(format string, v ...interface{}) {
	l.loggers[LOGLEVEL_CONFIG].Printf(format, v...)
}

func (l *logger) Warn(v ...interface{}) {
	l.loggers[LOGLEVEL_WARN].Print(v...)
}

func (l *logger) Warnf(format string, v ...interface{}) {
	l.loggers[LOGLEVEL_WARN].Printf(format, v...)
}

func (l *logger) Error(v ...interface{}) {
	l.loggers[LOGLEVEL_ERROR].Print(v...)
}

func (l *logger) Errorf(format string, v ...interface{}) {
	l.loggers[LOGLEVEL_ERROR].Printf(format, v...)
}

func (l *logger) Alert(v ...interface{}) {
	l.loggers[LOGLEVEL_ALERT].Print(v...)
}

func (l *logger) Alertf(format string, v ...interface{}) {
	l.loggers[LOGLEVEL_ALERT].Printf(format, v...)
}

func (l *logger) Fatal(v ...interface{}) {
	l.loggers[LOGLEVEL_FATAL].Fatal(v...)
}

func (l *logger) Fatalf(format string, v ...interface{}) {
	l.loggers[LOGLEVEL_FATAL].Fatalf(format, v...)
}

var logLevelToName map[int]string

func init() {
	logLevelToName = make(map[int]string, NR_LOGLEVELS)
	logLevelToName[LOGLEVEL_DEBUG] = "[Debug]"
	logLevelToName[LOGLEVEL_INFO] = "[Info]"
	logLevelToName[LOGLEVEL_CONFIG] = "[Config]"
	logLevelToName[LOGLEVEL_WARN] = "[Warning]"
	logLevelToName[LOGLEVEL_ERROR] = "[Error]"
	logLevelToName[LOGLEVEL_ALERT] = "[Alert]"
	logLevelToName[LOGLEVEL_FATAL] = "[Fatal]"
}

func NewLogger(writer io.Writer, prefix string, logLevel int) Logger {
	ret := new(logger)
	ret.loggers = make([]*log.Logger, NR_LOGLEVELS)
	if writer == nil {
		ret.writer = &nullWriter{}
	} else {
		ret.writer = writer
	}
	ret.prefix = prefix
	ret.SetLogLevel(logLevel)
	return ret
}

func (l *logger) SetLogLevel(logLevel int) {
	if logLevel > LOGLEVEL_DEBUG {
		logLevel = LOGLEVEL_DEBUG
	}
	l.logLevel = logLevel
	for i := 0; i <= logLevel; i++ {
		l.loggers[i] = log.New(l.writer, l.prefix+logLevelToName[i]+" ", log.LstdFlags)
	}
	nullwriter := &nullWriter{}
	for i := logLevel + 1; i < NR_LOGLEVELS; i++ {
		l.loggers[i] = log.New(nullwriter, l.prefix+logLevelToName[i]+" ", log.LstdFlags)
	}
}

