<!-- +
title: about logf - a auto spliting file log tool
urlid: logf-a-auto-spliting-file-log-tool
overview: logf is a auto spliting file log tool. it depends a time format to split log file automaticly that makes your log management easier
cate: go-tools
tags: #logf, #golang, #go, #log, #auto-spliting-file
lang: en
published_at: 2020-09-07T14:50:00Z
updated_at: 2020-09-07T14:50:00Z
+ --> 

# introduction

we demand different size of the log file. maybe split it as day, week or month, it always depends on period. this package provide a tool to do that. all you need to do is assigning a directory to save your log files, and tell them a time format used as indicator to split file. then use it as a standard golang log

# sample

```golang

import logf

// config log filename prefix
logf.SetFilenamePrefix("test-")

// split log file each month
logf.SetFormat("2006-01")

// save the log file in logs folder
SetPath("./logs/")

// log to file
logf.Println("you just set done your logf. so use it free")

```

# api

```golang

type Logf struct {
	Path   string
	Format string
	Prefix string

	m        sync.Mutex
	ins      *log.Logger
	filename string
}

func (logf *Logf) EnsurePathExist()

func (logf *Logf) Ins() *log.Logger

func SetFilenamePrefix(p string)

func SetFormat(f string)

func SetPath(p string)

func Fatal(v ...interface{})

func Fatalf(format string, v ...interface{})

func Fatalln(v ...interface{})

func Output(calldepth int, s string) error

func Panic(v ...interface{})

func Panicf(format string, v ...interface{})

func Panicln(v ...interface{})

func Prefix() string

func Print(v ...interface{})

func Printf(format string, v ...interface{})

func Println(v ...interface{})

func SetFlags(flag int)

func Flags() int

func SetPrefix(prefix string)

func getStdIns() *log.Logger

```
