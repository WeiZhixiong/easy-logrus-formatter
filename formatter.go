package formatter

import (
	"bytes"
	"fmt"
	"github.com/sirupsen/logrus"
	"strings"
	"time"
)

// LogField is a type of log field.
type LogField string

const (
	logTime = LogField("time")
	level   = LogField("level")
	msg     = LogField("msg")
)

var (
	defaultDelimiter  = "||"
	defaultLogFields  = []LogField{logTime, level, msg}
	defaultTimeFormat = time.RFC3339
)

// Formatter implements logrus.Formatter
type Formatter struct {
	// Delimiter is a delimiter to be used between log fields.
	Delimiter string
	// LogFields is a list of fields to be logged.
	LogFields []LogField
	// TimeFormat is a time format to be logged.
	TimeFormat string
}

func (f *Formatter) Format(entry *logrus.Entry) ([]byte, error) {
	f.setDefaultValues()

	var b *bytes.Buffer
	if entry.Buffer != nil {
		b = entry.Buffer
	} else {
		b = &bytes.Buffer{}
	}

	for i, field := range f.LogFields {
		if i > 0 {
			b.WriteString(f.Delimiter)
		}
		switch field {
		case logTime:
			b.WriteString(entry.Time.Format(f.TimeFormat))
		case level:
			b.WriteString(strings.ToUpper(entry.Level.String()))
		case msg:
			b.WriteString(entry.Message)
			if entry.Data != nil {
				for k, v := range entry.Data {
					fmt.Fprintf(b, "%s%s=%v", f.Delimiter, k, v)
				}
			}
		}

	}

	b.WriteByte('\n')
	return b.Bytes(), nil
}

func (f *Formatter) setDefaultValues() {
	if f.Delimiter == "" {
		f.Delimiter = defaultDelimiter
	}
	if f.LogFields == nil || len(f.LogFields) == 0 {
		f.LogFields = defaultLogFields
	}
	if f.TimeFormat == "" {
		f.TimeFormat = defaultTimeFormat
	}
}
