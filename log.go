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

// one copy form log
const (
	Ldate = 1 << iota
	Ltime
	Lmicroseconds
	Llongfile
	Lshortfile
	LstdFlags = Ldate | Ltime
)

const (
	_                = 0 // skip zero for clear flag argument of log.New
	LOGLEVEL_FATAL   = -iota
	LOGLEVEL_RECOVER // for recover panic
	LOGLEVEL_ALERT
	LOGLEVEL_ERROR
	LOGLEVEL_WARN
	LOGLEVEL_CONFIG
	LOGLEVEL_INFO
	LOGLEVEL_DEBUG
	NR_LOGLEVELS
	EQUAL_LEVEL     // equal level mode
	NONE_LEVEL_NAME // dont write default level name
	DONT_EXIT       // no os.Exit
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
	Write([]byte) (int, error)
	Recover(v ...interface{})
	Recoverf(format string, v ...interface{})
}

type logger struct {
	level         int
	log           *log.Logger
	writer        io.Writer
	prefix        string
	flags         int
	equal         bool
	noneLevelName bool
	dontExit      bool
}

func (l *logger) ok(level int) bool {
	level = -level
	if l.equal && level == l.level || level <= l.level {
		if !l.noneLevelName {
			l.writer.Write([]byte(logLevelToName[level]))
		}
		return true
	}
	return false
}

func (l *logger) Write(p []byte) (int, error) {
	return l.writer.Write(p)
}

func (l *logger) Debug(v ...interface{}) {
	if l.ok(LOGLEVEL_DEBUG) {
		l.log.Print(v...)
	}
}

func (l *logger) Debugf(format string, v ...interface{}) {
	if l.ok(LOGLEVEL_DEBUG) {
		l.log.Printf(format, v...)
	}
}

func (l *logger) Info(v ...interface{}) {
	if l.ok(LOGLEVEL_INFO) {
		l.log.Print(v...)
	}
}

func (l *logger) Infof(format string, v ...interface{}) {
	if l.ok(LOGLEVEL_INFO) {
		l.log.Printf(format, v...)
	}
}

func (l *logger) Config(v ...interface{}) {
	if l.ok(LOGLEVEL_CONFIG) {
		l.log.Print(v...)
	}
}

func (l *logger) Configf(format string, v ...interface{}) {
	if l.ok(LOGLEVEL_CONFIG) {
		l.log.Printf(format, v...)
	}
}

func (l *logger) Warn(v ...interface{}) {
	if l.ok(LOGLEVEL_WARN) {
		l.log.Print(v...)
	}
}

func (l *logger) Warnf(format string, v ...interface{}) {
	if l.ok(LOGLEVEL_WARN) {
		l.log.Printf(format, v...)
	}
}

func (l *logger) Error(v ...interface{}) {
	if l.ok(LOGLEVEL_ERROR) {
		l.log.Print(v...)
	}
}

func (l *logger) Errorf(format string, v ...interface{}) {
	if l.ok(LOGLEVEL_ERROR) {
		l.log.Printf(format, v...)
	}
}

func (l *logger) Alert(v ...interface{}) {
	if l.ok(LOGLEVEL_ALERT) {
		l.log.Print(v...)
	}
}

func (l *logger) Alertf(format string, v ...interface{}) {
	if l.ok(LOGLEVEL_ALERT) {
		l.log.Printf(format, v...)
	}
}

func (l *logger) Recover(v ...interface{}) {
	if l.ok(LOGLEVEL_RECOVER) {
		l.log.Print(v...)
	}
}

func (l *logger) Recoverf(format string, v ...interface{}) {
	if l.ok(LOGLEVEL_RECOVER) {
		l.log.Printf(format, v...)
	}
}

func (l *logger) Fatal(v ...interface{}) {
	if l.ok(LOGLEVEL_FATAL) {
		if l.dontExit {
			l.log.Print(v...)
		} else {
			l.log.Fatal(v...)
		}
	}
}

func (l *logger) Fatalf(format string, v ...interface{}) {
	if l.ok(LOGLEVEL_FATAL) {
		if l.dontExit {
			l.log.Printf(format, v...)
		} else {
			l.log.Fatalf(format, v...)
		}
	}
}

var logLevelToName [-NR_LOGLEVELS]string

func init() {
	logLevelToName[-LOGLEVEL_DEBUG] = "[Debug]"
	logLevelToName[-LOGLEVEL_INFO] = "[Info]"
	logLevelToName[-LOGLEVEL_CONFIG] = "[Config]"
	logLevelToName[-LOGLEVEL_WARN] = "[Warning]"
	logLevelToName[-LOGLEVEL_ERROR] = "[Error]"
	logLevelToName[-LOGLEVEL_ALERT] = "[Alert]"
	logLevelToName[-LOGLEVEL_FATAL] = "[Fatal]"
	logLevelToName[-LOGLEVEL_RECOVER] = "[Recover]"
}

func NewLogger(writer io.Writer, prefix string, flags ...int) Logger {
	ret := new(logger)
	hasflags := false
	for _, flag := range flags {
		if flag == EQUAL_LEVEL {
			ret.equal = true
			continue
		}
		if flag == NONE_LEVEL_NAME {
			ret.noneLevelName = true
			continue
		}
		if flag == DONT_EXIT {
			ret.dontExit = true
			continue
		}
		if flag >= 0 {
			hasflags = true
			ret.flags = ret.flags | flag
		} else {
			ret.level = -flag
		}
	}
	// defaults to LstdFlags.
	if !hasflags {
		ret.flags = LstdFlags
	}

	if ret.level >= -NR_LOGLEVELS {
		ret.level = -NR_LOGLEVELS - 1
	}

	ret.prefix = prefix
	if writer == nil {
		ret.writer = &nullWriter{}
	} else {
		ret.writer = writer
	}
	ret.log = log.New(writer, prefix+" ", ret.flags)
	return ret
}
