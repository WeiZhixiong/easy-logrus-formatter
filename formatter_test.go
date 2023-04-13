package formatter

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"strings"
	"testing"
	"time"
)

var (
	testCases = []struct {
		message string
		level   logrus.Level
		time    time.Time
		field   logrus.Fields
	}{
		{"Test Message1", logrus.DebugLevel, time.Now(), nil},
		{"Test Message2", logrus.InfoLevel, time.Now(), nil},
		{"Test Message3", logrus.WarnLevel, time.Now(), nil},
		{"Test Message1", logrus.DebugLevel, time.Now(), map[string]interface{}{"key1": "value1"}},
		{"Test Message2", logrus.InfoLevel, time.Now(), map[string]interface{}{"key1": "value1"}},
		{"Test Message3", logrus.WarnLevel, time.Now(), map[string]interface{}{"key1": "value1"}},
	}
)

func TestFormatterDefaultFormat(t *testing.T) {
	f := &Formatter{}
	l := logrus.New()
	for _, tt := range testCases {
		e := logrus.NewEntry(l)
		e = e.WithFields(tt.field)
		e.Message = tt.message
		e.Level = tt.level
		e.Time = tt.time
		b, _ := f.Format(e)

		expected := strings.Join([]string{e.Time.Format(f.TimeFormat), strings.ToUpper(e.Level.String()), e.Message}, f.Delimiter)
		if e.Data != nil {
			for k, v := range e.Data {
				expected += fmt.Sprintf("%s%s=%s", f.Delimiter, k, v)
			}
		}
		expected += "\n"

		if string(b) != expected {
			t.Errorf("formatting expected result was %q instead of %q", string(b), expected)
		}
	}
}

func TestFormatterCustomFormat(t *testing.T) {
	f := &Formatter{
		Delimiter:  ",",
		LogFields:  []LogField{level, logTime, msg},
		TimeFormat: time.RFC3339Nano,
	}
	l := logrus.New()
	for _, tt := range testCases {
		e := logrus.NewEntry(l)
		e = e.WithFields(tt.field)
		e.Message = tt.message
		e.Level = tt.level
		e.Time = tt.time
		b, _ := f.Format(e)

		expected := strings.Join([]string{strings.ToUpper(e.Level.String()), e.Time.Format(f.TimeFormat), e.Message}, f.Delimiter)
		if e.Data != nil {
			for k, v := range e.Data {
				expected += fmt.Sprintf("%s%s=%s", f.Delimiter, k, v)
			}
		}
		expected += "\n"

		if string(b) != expected {
			t.Errorf("formatting expected result was %q instead of %q", string(b), expected)
		}
	}
}
