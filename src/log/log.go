package log

import (
	"fmt"
	"os"
	"time"

	"webserver/src/errors"
)

const debug int = 1
const info 	int = 2
const warn 	int = 3
const err 	int = 4

var _log_lv_names = [4]string{"debug","info","warn","error"}

type Log interface {
	Debug(string)			(int, errors.Error)
	Info(string)			(int, errors.Error)
	Warn(string)			(int, errors.Error)
	Err(string)				(int, errors.Error)
	printLog(int, string)	(int, errors.Error)
}

type Logger struct {
	logLv 	int			// The Min Log Level to Trigger printLog()
	name	string		// The Name of the logger (Optional)	
	out	   	*os.File	// Where to Output DEBUG and INFO messages to
	err		*os.File	// Where to Output WARN and ERROR messages to
}

func New(lv int, name string) Log {
	return &Logger{lv, name, os.Stdout, os.Stderr}
}

func getTimeStamp (now time.Time) string {
	ret := fmt.Sprintf("%d-%02d-%02dT%02d:%02d:%02d.%.3d", now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second(), now.Nanosecond()  / 1_000_000)

	return ret 
}

func (l *Logger) printLog(lv int, msg string) (int, errors.Error) {
	if l.logLv > lv{
		return 0, nil
	}

	if lv > err {
		return 0, errors.New("Log LV Can't be > 4")
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
	if len(l.name) > 0 {
		_log_name = fmt.Sprintf("[%s]", l.name)
	}

	ret, err := fmt.Fprintf(_output, "[%s]%s[%s] --> %s\n", timestamp, _log_name, lv_name, msg)

	return ret, err
}

func (l* Logger) Debug	(msg string) 	(int, errors.Error)	{ return l.printLog(debug, msg) }
func (l* Logger) Info	(msg string) 	(int, errors.Error)	{ return l.printLog(info, msg) 	}
func (l* Logger) Warn	(msg string) 	(int, errors.Error)	{ return l.printLog(warn, msg) 	}
func (l* Logger) Err	(msg string) 	(int, errors.Error)	{ return l.printLog(err, msg) 	}