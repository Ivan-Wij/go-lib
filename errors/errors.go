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
	sources := make([]string, len(e.Traces))
	funcs := make([]string, len(e.Traces))
	for i := range e.Traces {
		sources[i] = fmt.Sprintf("\"%s\"", e.Traces[i].Source)
		funcs[i] = fmt.Sprintf("\"%s\"", e.Traces[i].Func)
	}
	return fmt.Sprintf("{\"Error\": \"%s\", \"Code\": \"%s\", \"Sources\": %v, \"Funcs\": %v}", e.Err.Error(), e.Code, sources, funcs)
}

func new(err error) (newErr Err) {
	switch v := err.(type) {
	case Err:
		newErr = v
	default:
		newErr.Err = v
		newErr.Func = getFuncName(3)
		newErr.Source = getSource(3)
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
	if caller == nil {
		return ""
	}
	callerName := caller.Name()
	return callerName[strings.LastIndex(callerName, "/")+1:]
}

// original error will be put on top, current trace will be added last
func AddTrace(err error, errs ...Err) Err {
	newErr := new(err)

	newErr.Traces = append(newErr.Traces, Err{
		Func:   getFuncName(2),
		Source: getSource(2),
	})
	return newErr
}
