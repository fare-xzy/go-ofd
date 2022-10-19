package log

import (
	"fmt"
	"github.com/rs/zerolog"
)

func Assert(condition bool, msg string) {
	if !condition {
		Panicf("Assert failed: %s", msg)
	}
}

func Debugf(format string, args ...any) {
	logFinally(logger.Debug(), format, args...)
}

func Infof(format string, args ...any) {
	logFinally(logger.Info(), format, args...)
}

func Warnf(format string, args ...any) {
	logFinally(logger.Warn(), format, args...)
}

func Errorf(format string, args ...any) {
	logFinally(logger.Error(), format, args...)
}

func Panicf(format string, args ...any) {
	logFinally(logger.Panic(), format, args...)
}

func PanicError(err error) {
	Panicf(err.Error())
}

func PanicFinal(a ...any) {
	str := fmt.Sprintln(a)
	logger.Panic().Msg(str)
}

func logFinally(event *zerolog.Event, format string, args ...any) {
	str := fmt.Sprintf(format, args...)
	event.Msg(str)
}
