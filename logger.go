package akinformer

import (
	"fmt"
	"io"

	"golang.org/x/net/context"

	"os"

	gaelog "google.golang.org/appengine/log"
)

// Log is the actual logger to be used
type Log struct {
	ctx context.Context
}

// Fprintf satisifies the Logger interface
func (l Log) Fprintf(out io.Writer, format string, v ...interface{}) {
	fmt.Fprintf(out, format, v)
}

// LoggerWithContext returns a new logger instance with a given context
func LoggerWithContext(c context.Context) *Log {
	return &Log{ctx: c}
}

// Errorf writes the error log to the gae log if context is present, stderr otherwise
func (l Log) Errorf(format string, v ...interface{}) {
	if l.ctx != nil {
		gaelog.Errorf(l.ctx, format, v)
	}
	l.Fprintf(os.Stderr, format+"\n", v)
}

// Infof writes the info log to the gae log if context is present, stdout otherwise
func (l Log) Infof(format string, v ...interface{}) {
	if l.ctx != nil {
		gaelog.Infof(l.ctx, format, v)
	}
	l.Fprintf(os.Stdout, format+"\n", v)
}
