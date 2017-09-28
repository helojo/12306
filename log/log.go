package log

import (
	"fmt"
	"time"
)

var debug bool = true

type LogLevel string

const (
	DEBUG   = "D"
	INFO    = "I"
	WARNING = "W"
	ERROR   = "E"
)

type LogType string

const (
	LOGIN     = "登录"
	PASSENGER = "获取用户信息"
	OTHER     = "其它"
	ORDER     = "订票"
	CHECK     = "查票"
)

func SetDebug(d bool) {
	debug = d
}
func MyLoginLogI(format string, a ...interface{}) {
	MyLog(INFO, LOGIN, format, a...)
}
func MyLoginLogW(format string, a ...interface{}) {
	MyLog(WARNING, LOGIN, format, a...)
}
func MyLoginLogE(format string, a ...interface{}) {
	MyLog(ERROR, LOGIN, format, a...)
}

func MyLogDebug(format string, a ...interface{}) {
	MyLog(DEBUG, OTHER, format, a...)
}
func MyLogInfo(format string, a ...interface{}) {
	MyLog(INFO, OTHER, format, a...)
}

// order log
func MyOrderLogI(format string, a ...interface{}) {
	MyLog(INFO, ORDER, format, a...)
}
func MyOrderLogW(format string, a ...interface{}) {
	MyLog(WARNING, ORDER, format, a...)
}
func MyOrderLogE(format string, a ...interface{}) {
	MyLog(ERROR, ORDER, format, a...)
}
func MyOrderLogD(format string, a ...interface{}) {
	MyLog(DEBUG, ORDER, format, a...)
}

// check log
func MyCheckLogI(format string, a ...interface{}) {
	MyLog(INFO, CHECK, format, a...)
}
func MyCheckLogW(format string, a ...interface{}) {
	MyLog(WARNING, CHECK, format, a...)
}
func MyCheckLogE(format string, a ...interface{}) {
	MyLog(ERROR, CHECK, format, a...)
}
func MyCheckLogD(format string, a ...interface{}) {
	MyLog(DEBUG, CHECK, format, a...)
}

// other log
func MyLogI(format string, a ...interface{}) {
	MyLog(INFO, OTHER, format, a...)
}
func MyLogW(format string, a ...interface{}) {
	MyLog(WARNING, OTHER, format, a...)
}
func MyLogE(format string, a ...interface{}) {
	MyLog(ERROR, OTHER, format, a...)
}
func MyLogD(format string, a ...interface{}) {
	MyLog(DEBUG, OTHER, format, a...)
}

func MyLog(l LogLevel, t LogType, format string, a ...interface{}) {
	if l == DEBUG && debug == false {
		return
	}
	format = string("[%v][%s][%s]:") + format + string("\n")
	arg := make([]interface{}, 3, len(a)+3)
	arg[0] = time.Now().Format("2006-01-02 15:04:05.999")
	arg[1] = t
	arg[2] = l
	if len(a) > 0 {
		arg = append(arg, a...)
	}
	fmt.Printf(format, arg...)
}
