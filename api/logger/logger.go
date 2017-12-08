package logger

import (
 "github.com/goinggo/tracelog"
)

func init() {
	tracelog.Start(tracelog.LevelTrace)
}

var Trace = tracelog.Trace
var Info = tracelog.Info
var Warn = tracelog.Warning
var Error = tracelog.Error
