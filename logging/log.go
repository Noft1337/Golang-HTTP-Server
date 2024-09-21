package logging

import (
	"fmt"
	"os"
	"runtime"
	"strings"
	"time"
)

const debug int = 1
const info 	int = 2
const warn 	int = 3
const err 	int = 4

var _log_lv_names = [4]string{"debug","info","warn","error"}

const LOG_LV int = 2
const LOG_NAME string = "xhttp"

type Log interface {
	Debug(string)			int
	Info(string)			int
	Warn(string)			int
	Err(string)				int
	printLog(int, string)	int
}

type Logger struct {
	LogLv 	int			// The Min Log Level to Trigger printLog()
	Name	string		// The Name of the logger (Optional)	
	out	   	*os.File	// Where to Output DEBUG and INFO messages to
	err		*os.File	// Where to Output WARN and ERROR messages to
}

func NewLogger(lv int, name string) *Logger {
	return &Logger{lv, name, os.Stdout, os.Stderr}
}

func getTimeStamp (now time.Time) string {
	ret := fmt.Sprintf("%d-%02d-%02dT%02d:%02d:%02d.%.3d", now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second(), now.Nanosecond()  / 1_000_000)

	return ret 
}

func GetPackageName(temp interface{}) (string, string) {
	pc, _, _, ok := runtime.Caller(3)
	if !ok {
		return "unknown", "unknown"
	}

	names := strings.Split(runtime.FuncForPC(pc).Name(), ".")
	
	lastSlash := strings.LastIndex(names[0], "/")
	pkg := names[0][lastSlash + 1:]
	foo := names[1]	


	return pkg, foo
}

func (l *Logger) printLog(lv int, msg string, args ...interface{}) int {
	packageName, funcName := GetPackageName(msg)
	
	if l.LogLv > lv{
		return 0
	}

	if lv > err {
		return 0
	}

	var _output *os.File
	if lv < warn {
		_output = l.out
	} else {
		_output = l.err
	}

	var ret int
	lv_name := _log_lv_names[lv - 1]

	_current_time := time.Now()
	timestamp := getTimeStamp(_current_time)

	_log_name := ""
	if len(l.Name) > 0 {
		_log_name = fmt.Sprintf("[%s]", l.Name)
	}

	finalMsg := fmt.Sprintf(msg, args...)
	ret, _ = fmt.Fprintf(_output, "[%s]%s[%s]>(%s::%s) --> %s\n", timestamp, _log_name, lv_name, packageName, funcName, finalMsg)

	return ret
}

func (l* Logger) Debug	(msg string, args ...interface{}) 	int	{ return l.printLog(debug, msg, args...) 	}
func (l* Logger) Info	(msg string, args ...interface{}) 	int	{ return l.printLog(info, msg, args...) 	}
func (l* Logger) Warn	(msg string, args ...interface{}) 	int	{ return l.printLog(warn, msg, args...) 	}
func (l* Logger) Err	(msg string, args ...interface{}) 	int	{ return l.printLog(err, msg, args...) 	}

