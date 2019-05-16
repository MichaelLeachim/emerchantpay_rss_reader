// @@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
// @ Copyright (c) Michael Leahcim                                                      @
// @ You can find additional information regarding licensing of this work in LICENSE.md @
// @ You must not remove this notice, or any other, from this software.                 @
// @ All rights reserved.                                                               @
// @@@@@@ At 2019-05-16 22:57 <thereisnodotcollective@gmail.com> @@@@@@@@@@@@@@@@@@@@@@@@

package emerchantpay_rss_reader

import (
	logrus "github.com/sirupsen/logrus"
	"testing"
)

type Logger interface {
	Warn(args ...interface{})
	Error(args ...interface{})
}

type MockLogger struct{}

func newMockLogger() Logger {
	return MockLogger{}
}

func (m MockLogger) Error(args ...interface{}) {
	return
}
func (m MockLogger) Warn(args ...interface{}) {
	return
}

type TestLogger struct {
	t *testing.T
}

func newTestLogger(t *testing.T) Logger {
	return TestLogger{t: t}
}

func (m TestLogger) Error(args ...interface{}) {
	m.t.Log(args...)
	return
}

func (m TestLogger) Warn(args ...interface{}) {
	m.t.Log(args...)

}

type LogrusLogger struct {
}

func newLogrusLogger() Logger {
	return LogrusLogger{}
}

func (m LogrusLogger) Error(args ...interface{}) {
	logrus.Error(args...)

}

func (m LogrusLogger) Warn(args ...interface{}) {
	logrus.Warn(args...)
}
