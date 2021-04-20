package main

import (
	"log"
	"os"
)

type Logging struct {
	log.Logger
	Level	int
}

func (l Logging) Trace(args... interface{}) {
	if l.Level == 0 {
		args = append([]interface{}{"TRACE "}, args...)
		log.Print(args...)
	}
}

func (l Logging) Debug(args... interface{}) {
	if l.Level < 2 {
		args = append([]interface{}{"DEBUG "}, args...)
		log.Print(args...)
	}
}

func (l Logging) Info(args... interface{}) {
	if l.Level < 3 {
		args = append([]interface{}{"INFO "}, args...)
		log.Print(args...)
	}
}

func (l Logging) Warn(args ...interface{}) {
	if l.Level < 4 {
		args = append([]interface{}{"WARN "}, args...)
		log.Print(args...)
	}
}

func (l Logging) Error(args ...interface{}) {
	if l.Level < 5 {
		args = append([]interface{}{"ERROR "}, args...)
		log.Print(args...)
	}
}

func (l Logging) Fatal(args ...interface{}) {
	if l.Level < 6 {
		args = append([]interface{}{"FATAL "}, args...)
		log.Print(args...)
		os.Exit(1)
	}
}
