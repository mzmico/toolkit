package errors

import (
	"bytes"
	"fmt"
	"runtime"
)

type Context struct {
	Traces []*StackTrace
}

func (m *Context) String() string {

	w := bytes.NewBuffer(nil)

	for _, trace := range m.Traces {
		fmt.Fprintf(
			w,
			"+-- %s:[ %s:%d ]\n",
			trace.Function,
			trace.File,
			trace.Line)
	}

	return w.String()
}

type StackTrace struct {
	Function string
	File     string
	Line     int
}

func trace() (stack []*StackTrace) {

	var (
		SKIP int = 4
	)

	callers := make([]uintptr, 30)

	n := runtime.Callers(SKIP, callers)

	stack = make([]*StackTrace, 0)

	for index := 0; index < n; index++ {

		pc := callers[index]
		trace := &StackTrace{}

		runtime.Caller(index)

		fp := runtime.FuncForPC(pc)

		if fp != nil {
			trace.File, trace.Line = fp.FileLine(pc)
			trace.Function = fp.Name()
		}

		stack = append(stack, trace)
	}

	return
}

func NewContext() *Context {

	return &Context{
		Traces: trace(),
	}
}
