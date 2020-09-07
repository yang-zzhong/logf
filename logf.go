package logf

// this package implement a file logger that auto split into deferent file by time format
import (
	"log"
	"os"
	"sync"
	"time"
)

type Logf struct {
	Path   string
	Format string
	Prefix string

	m        sync.Mutex
	ins      *log.Logger
	filename string
}

var (
	std    *Logf
	format string
	path   string
	prefix string
)

func init() {
	std = nil
	format = "2006-01-02"
	prefix = ""
	path = "./"
}

func (logf *Logf) EnsurePathExist() {
	if !exists(logf.Path) {
		if err := os.MkdirAll(logf.Path, os.ModePerm); err != nil {
			panic(err)
		}
	}
}

func (logf *Logf) Ins() *log.Logger {
	logf.m.Lock()
	defer logf.m.Unlock()
	filename := logf.Path + logf.Prefix + time.Now().Format(logf.Format) + ".log"
	if filename == logf.filename {
		return logf.ins
	}
	logf.filename = filename
	file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}

	logf.ins = log.New(file, "", log.LstdFlags)

	return logf.ins
}

func SetFilenamePrefix(p string) {
	prefix = p
}

func SetFormat(f string) {
	format = f
}

func SetPath(p string) {
	path = p
}

func Fatal(v ...interface{}) {
	getStdIns().Fatal(v...)
}

func Fatalf(format string, v ...interface{}) {
	getStdIns().Fatalf(format, v...)
}

func Fatalln(v ...interface{}) {
	getStdIns().Fatalln(v...)
}

func Output(calldepth int, s string) error {
	return getStdIns().Output(calldepth, s)
}

func Panic(v ...interface{}) {
	getStdIns().Panic(v...)
}

func Panicf(format string, v ...interface{}) {
	getStdIns().Panicf(format, v...)
}

func Panicln(v ...interface{}) {
	getStdIns().Panicln(v...)
}

func Prefix() string {
	return getStdIns().Prefix()
}

func Print(v ...interface{}) {
	getStdIns().Print(v...)
}

func Printf(format string, v ...interface{}) {
	getStdIns().Printf(format, v...)
}

func Println(v ...interface{}) {
	getStdIns().Println(v...)
}

func SetFlags(flag int) {
	getStdIns().SetFlags(flag)
}

func Flags() int {
	return getStdIns().Flags()
}

func SetPrefix(prefix string) {
	getStdIns().SetPrefix(prefix)
}

func getStdIns() *log.Logger {
	if std == nil {
		std = &Logf{Path: path, Format: format, Prefix: prefix}
		std.EnsurePathExist()
	}
	return std.Ins()
}

func exists(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		return os.IsExist(err)
	}
	return true
}
