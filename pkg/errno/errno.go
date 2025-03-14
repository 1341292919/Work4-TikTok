package errno

import (
	"errors"
	"fmt"
	"io"
)

type ErrNo struct {
	ErrorCode int64
	ErrorMsg  string
	stack     *stack
}

func NewErrNo(code int64, msg string) ErrNo {
	return ErrNo{
		ErrorCode: code,
		ErrorMsg:  msg,
	}
}

/* error是一个接口类型的变量
它可以接受赋值给它的变量的类型
*/

func (e ErrNo) Error() string { return fmt.Sprintf("[%d]: %s", e.ErrorCode, e.ErrorMsg) }

func NewErrNoWithStack(code int64, msg string) ErrNo {
	return ErrNo{
		ErrorCode: code,
		ErrorMsg:  msg,
		stack:     callers(),
	}
}

func (e ErrNo) WithMessage(message string) ErrNo {
	e.ErrorMsg = message
	return e
}
func (e ErrNo) WithError(err error) ErrNo {
	e.ErrorMsg = err.Error()
	return e
}

func (e ErrNo) Format(s fmt.State, verb rune) {
	switch verb {
	case 's':
		io.WriteString(s, e.Error())
	case 'd':
		io.WriteString(s, e.Error())
		switch {
		case s.Flag('+'):
			e.stack.Format(s, verb)
		}
	}
}

func ConvertErr(err error) ErrNo {
	if err == nil {
		return Success
	}
	errno := ErrNo{}
	//如果err已经是该类型
	if errors.As(err, &errno) {
		return errno
	}
	s := InternalServiceError
	s.ErrorMsg = err.Error()
	return s
}
