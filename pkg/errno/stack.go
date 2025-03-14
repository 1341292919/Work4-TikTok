package errno

import (
	"fmt"
	"runtime"
)

type Frame uintptr //栈帧
type stack []uintptr

func (s *stack) Format(st fmt.State, verb rune) {
	//nolint
	switch verb {
	case 'v':
		switch {
		case st.Flag('+'):
			for _, pc := range *s {
				f := Frame(pc)
				_, _ = fmt.Fprintf(st, "\n%+v", f)
			}
		}
	}
}

var skip = 3

// 为什么 skip = 3？
// runtime.Callers 会从调用栈的顶部开始捕获栈帧。
// 跳过的栈帧包括：
// runtime.Callers 自身的调用。
// callers 函数的调用。
// 调用 callers 函数的上一层函数。
func callers() *stack {
	const depth = 32
	var pcs [depth]uintptr
	n := runtime.Callers(skip, pcs[:])
	var st stack = pcs[0:n]
	return &st
}
