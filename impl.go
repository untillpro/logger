/*
 * Copyright (c) 2020-present unTill Pro, Ltd. and Contributors
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */

package logger

import (
	"context"
	"fmt"
	"runtime"
	"strings"
	"sync/atomic"
	"time"
)

// TLogLevel s.e.
type TLogLevel int32

// Log Levels
const (
	LogLevelNone = TLogLevel(iota)
	LogLevelError
	LogLevelWarning
	LogLevelInfo
	LogLevelVerbose
)

// LogLevel s.e.
var globalLogLevel = LogLevelInfo

// SetLogLevel s.e.
func SetLogLevel(logLevel TLogLevel) {
	atomic.StoreInt32((*int32)(&globalLogLevel), int32(logLevel))
}

// IsEnabled s.e.
func IsEnabled(logLevel TLogLevel) bool {
	curLogLevel := TLogLevel(atomic.LoadInt32((*int32)(&globalLogLevel)))
	return curLogLevel >= logLevel
}

// IsVerbose s.e.
func IsVerbose(ctx context.Context) bool {
	return IsEnabled(LogLevelVerbose)
}

func print(msgType string, args ...interface{}) {
	var funcName string
	pc, _, _, ok := runtime.Caller(2)
	details := runtime.FuncForPC(pc)
	if ok && details != nil {
		elems := strings.Split(details.Name(), "/")
		if len(elems) > 1 {
			funcName = elems[len(elems)-1]
		} else {
			funcName = details.Name()
		}
	} else {
		funcName = ""
	}
	t := time.Now()
	fmt.Print(t.Format("2006/01/02 1504"))
	fmt.Print(" "+msgType+": ", funcName)
	if len(args) > 0 {
		fmt.Print(": ")
		fmt.Println(args...)
	}
}

// Error s.e.
func Error(ctx context.Context, args ...interface{}) {
	print("*** ERROR  ", args...)

}

// Warning s.e.
func Warning(ctx context.Context, args ...interface{}) {
	if IsEnabled(LogLevelWarning) {
		print("*** WARNING", args...)
	}
}

// Info s.e.
func Info(ctx context.Context, args ...interface{}) {
	if IsEnabled(LogLevelInfo) {
		print("*** INFO   ", args...)
	}
}

// Verbose s.e.
func Verbose(ctx context.Context, args ...interface{}) {
	if IsVerbose(ctx) {
		print("--- DEBUG  ", args...)
	}
}