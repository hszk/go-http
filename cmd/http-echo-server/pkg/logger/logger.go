package logger

import (
	"fmt"
	// I/Oプリミティブへの基本的なインタフェースを提供
	"io"
)

type Logger interface {
	Log(...interface{})
}

type logger struct {
	out io.Writer
}

func (l *logger) Log(args ...interface{}) {
	l.out.Write([]byte(fmt.Sprint(args...)))
}

// New makes new logger instance.
func New(w io.Writer) Logger {
	return &logger{out: w}
}
