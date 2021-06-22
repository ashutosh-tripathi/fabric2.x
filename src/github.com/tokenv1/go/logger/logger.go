package logger

import (
	"fmt"
	"strings"
)

const DEBUG = 0
const INFO = 1
const NOTICE = 2
const WARNING = 3
const ERROR = 4
const FATAL = 5

type logger struct {
	loglevel int
}

func NewLogger(logLevel string) *logger {
	levelString := strings.ToUpper(logLevel)

	level := 0
	switch levelString {

	case "DEBUG":
		level = DEBUG
		break
	case "INFO":
		level = INFO
		break
	case "NOTICE":
		level = NOTICE
		break
	case "WARNING":
		level = WARNING
		break
	case "ERROR":
		level = ERROR
		break
	}
	x := logger{level}
	fmt.Printf("Logger: Logging level=%s %d\n", levelString, level)
	return &x
}

func (l logger) DEBUG(s string) {
	if l.loglevel <= DEBUG {
		fmt.Printf("DEBUG: %s\n", s)
	}

}
func (l logger) ERROR(s string) {

	if l.loglevel <= ERROR {
		fmt.Printf("ERROR: %s\n", s)
	}

}
func (l logger) INFO(s string) {

	if l.loglevel <= INFO {
		fmt.Printf("INFO: %s\n", s)
	}

}
func (l logger) WARNING(s string) {

	if l.loglevel <= WARNING {
		fmt.Printf("WARNING: %s\n", s)
	}

}
func (l logger) NOTICE(s string) {

	if l.loglevel <= NOTICE {
		fmt.Printf("NOTICE: %s\n", s)
	}

}
