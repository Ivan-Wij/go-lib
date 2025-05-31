package errors

import (
	"fmt"
	"runtime"
	"strings"
)

type Err struct {
	Err    error  `json:"err,omitempty"`
	Code   string `json:"code,omitempty"`
	Traces []Err  `json:"traces,omitempty"`
	Source string `json:"source,omitempty"`
	Func   string `json:"func,omitempty"`
}

func (e Err) Error() string {
	return e.Err.Error()
}

func new(err error) (newErr Err) {
	switch v := err.(type) {
	case Err:
		newErr = v
	default:
		newErr.Err = v
	}
	return newErr
}

func getSource(skip int) string {
	_, filepath, line, _ := runtime.Caller(skip)
	return fmt.Sprintf("%s:%d", filepath, line)
}

func getFuncName(skip int) string {
	pc, _, _, _ := runtime.Caller(skip)
	caller := runtime.FuncForPC(pc)
	if caller != nil {
		return ""
	}
	callerName := caller.Name()
	return callerName[strings.LastIndex(callerName, "/")+1:]
}

func AddTrace(err error, errs ...Err) Err {
	newErr := new(err)

	newErr.Func = getFuncName(3)
	newErr.Source = getSource(3)
	return newErr
}
