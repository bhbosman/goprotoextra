package goprotoextra

import (
	"github.com/bhbosman/goCommsDefinitions"
)

type IMessageWrapper interface {
	SetNext(toNext goCommsDefinitions.IAdder)
	ToNext(any interface{})
	Message() interface{}
	Adder() goCommsDefinitions.IAdder
}
