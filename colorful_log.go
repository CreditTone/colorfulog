package log

import (
	"fmt"
	"github.com/wsxiaoys/terminal/color"
	"io"
	"os"
	"runtime"
	"sync"
	"time"
)

const (
	Ldate = 1 << iota
	Ltime
	Lmicroseconds
	Llongfile
	Lshortfile
	LUTC
	LstdFlags = Ldate | Ltime
)

type Logger struct {
	mu     sync.Mutex
	prefix string
	flag   int
	out    io.Writer
}

func init() {
	SetFlags(Lshortfile | LstdFlags)
}

func New(out io.Writer, prefix string, flag int) *Logger {
	return &Logger{out: out, prefix: prefix, flag: flag}
}

// SetOutput sets the output destination for the logger.
func (l *Logger) SetOutput(w io.Writer) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.out = w
}

var std = New(os.Stderr, "", LstdFlags)

func itoa(buf *[]byte, i int, wid int) {
	var b [20]byte
	bp := len(b) - 1
	for i >= 10 || wid > 1 {
		wid--
		q := i / 10
		b[bp] = byte('0' + i - q*10)
		bp--
		i = q
	}
	// i < 10
	b[bp] = byte('0' + i)
	*buf = append(*buf, b[bp:]...)
}

func (l *Logger) formatHeader(buf *[]byte, t time.Time, file string, line int) {
	*buf = append(*buf, l.prefix...)
	if l.flag&(Ldate|Ltime|Lmicroseconds) != 0 {
		if l.flag&LUTC != 0 {
			t = t.UTC()
		}
		if l.flag&Ldate != 0 {
			year, month, day := t.Date()
			itoa(buf, year, 4)
			*buf = append(*buf, '/')
			itoa(buf, int(month), 2)
			*buf = append(*buf, '/')
			itoa(buf, day, 2)
			*buf = append(*buf, ' ')
		}
		if l.flag&(Ltime|Lmicroseconds) != 0 {
			hour, min, sec := t.Clock()
			itoa(buf, hour, 2)
			*buf = append(*buf, ':')
			itoa(buf, min, 2)
			*buf = append(*buf, ':')
			itoa(buf, sec, 2)
			if l.flag&Lmicroseconds != 0 {
				*buf = append(*buf, '.')
				itoa(buf, t.Nanosecond()/1e3, 6)
			}
			*buf = append(*buf, ' ')
		}
	}
	if l.flag&(Lshortfile|Llongfile) != 0 {
		if l.flag&Lshortfile != 0 {
			short := file
			for i := len(file) - 1; i > 0; i-- {
				if file[i] == '/' {
					short = file[i+1:]
					break
				}
			}
			file = short
		}
		*buf = append(*buf, file...)
		*buf = append(*buf, ':')
		itoa(buf, line, -1)
		*buf = append(*buf, ": "...)
	}
}

func (l *Logger) outputColorful(colorDef, prefixFlag string, calldepth int, s string) error {
	//s = prefixFlag + s
	now := time.Now() // get this early.
	var file string
	var line int
	l.mu.Lock()
	defer l.mu.Unlock()
	if l.flag&(Lshortfile|Llongfile) != 0 {
		// Release lock while getting caller info - it's expensive.
		l.mu.Unlock()
		var ok bool
		_, file, line, ok = runtime.Caller(calldepth)
		if !ok {
			file = "???"
			line = 0
		}
		l.mu.Lock()
	}
	var buff []byte
	if len(prefixFlag) > 0 {
		buff = append(buff, prefixFlag...)
		buff = append(buff, " "...)
	}
	l.formatHeader(&buff, now, file, line)
	buff = append(buff, s...)
	if len(s) == 0 || s[len(s)-1] != '\n' {
		buff = append(buff, '\n')
	}
	if colorDef == "red" {
		buff = []byte(color.Sprintf("@r%v", string(buff)))
	} else if colorDef == "yellow" {
		buff = []byte(color.Sprintf("@y%v", string(buff)))
	}
	_, err := l.out.Write(buff)
	return err
}

func (l *Logger) Flags() int {
	l.mu.Lock()
	defer l.mu.Unlock()
	return l.flag
}

func (l *Logger) SetFlags(flag int) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.flag = flag
}

func (l *Logger) Prefix() string {
	l.mu.Lock()
	defer l.mu.Unlock()
	return l.prefix
}

func (l *Logger) SetPrefix(prefix string) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.prefix = prefix
}

func (l *Logger) Writer() io.Writer {
	l.mu.Lock()
	defer l.mu.Unlock()
	return l.out
}

func SetOutput(w io.Writer) {
	std.mu.Lock()
	defer std.mu.Unlock()
	std.out = w
}

func Flags() int {
	return std.Flags()
}

func SetFlags(flag int) {
	std.SetFlags(flag)
}

func Prefix() string {
	return std.Prefix()
}

func SetPrefix(prefix string) {
	std.SetPrefix(prefix)
}

func Writer() io.Writer {
	return std.Writer()
}

func Info(v ...interface{}) {
	std.outputColorful("", "[INFO]", 2, fmt.Sprintln(v...))
}

func Infof(format string, v ...interface{}) {
	std.outputColorful("", "[INFO]", 2, fmt.Sprintf(format, v...))
}

func Warn(v ...interface{}) {
	std.outputColorful("yellow", "[WARN]", 2, fmt.Sprintln(v...))
}

func Warnf(format string, v ...interface{}) {
	std.outputColorful("yellow", "[WARN]", 2, fmt.Sprintf(format, v...))
}

func Error(v ...interface{}) {
	std.outputColorful("red", "[ERROR]", 2, fmt.Sprintln(v...))
}

func Errorf(format string, v ...interface{}) {
	std.outputColorful("red", "[ERROR]", 2, fmt.Sprintf(format, v...))
}

func Fatal(v ...interface{}) {
	std.outputColorful("red", "[FATAL]", 2, fmt.Sprintln(v...))
	os.Exit(1)
}

func Fatalf(format string, v ...interface{}) {
	std.outputColorful("red", "[FATAL]", 2, fmt.Sprintf(format, v...))
	os.Exit(1)
}
