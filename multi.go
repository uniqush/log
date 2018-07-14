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

type multiLogger struct {
	loggers []Logger
}

func MultiLogger(loggers ...Logger) Logger {
	return &multiLogger{loggers}
}

func (self *multiLogger) Debug(v ...interface{}) {
	for _, l := range self.loggers {
		if l != nil {
			l.Debug(v...)
		}
	}
}

func (self *multiLogger) Debugf(format string, v ...interface{}) {
	for _, l := range self.loggers {
		if l != nil {
			l.Debugf(format, v...)
		}
	}
}

func (self *multiLogger) Info(v ...interface{}) {
	for _, l := range self.loggers {
		if l != nil {
			l.Info(v...)
		}
	}
}

func (self *multiLogger) Infof(format string, v ...interface{}) {
	for _, l := range self.loggers {
		if l != nil {
			l.Infof(format, v...)
		}
	}
}

func (self *multiLogger) Config(v ...interface{}) {
	for _, l := range self.loggers {
		if l != nil {
			l.Config(v...)
		}
	}
}

func (self *multiLogger) Configf(format string, v ...interface{}) {
	for _, l := range self.loggers {
		if l != nil {
			l.Configf(format, v...)
		}
	}
}

func (self *multiLogger) Warn(v ...interface{}) {
	for _, l := range self.loggers {
		if l != nil {
			l.Warn(v...)
		}
	}
}

func (self *multiLogger) Warnf(format string, v ...interface{}) {
	for _, l := range self.loggers {
		if l != nil {
			l.Warnf(format, v...)
		}
	}
}

func (self *multiLogger) Error(v ...interface{}) {
	for _, l := range self.loggers {
		if l != nil {
			l.Error(v...)
		}
	}
}

func (self *multiLogger) Errorf(format string, v ...interface{}) {
	for _, l := range self.loggers {
		if l != nil {
			l.Errorf(format, v...)
		}
	}
}

func (self *multiLogger) Alert(v ...interface{}) {
	for _, l := range self.loggers {
		if l != nil {
			l.Alert(v...)
		}
	}
}

func (self *multiLogger) Alertf(format string, v ...interface{}) {
	for _, l := range self.loggers {
		if l != nil {
			l.Alertf(format, v...)
		}
	}
}

func (self *multiLogger) Fatal(v ...interface{}) {
	for _, l := range self.loggers {
		if l != nil {
			l.Fatal(v...)
		}
	}
}

func (self *multiLogger) Fatalf(format string, v ...interface{}) {
	for _, l := range self.loggers {
		if l != nil {
			l.Fatalf(format, v...)
		}
	}
}
