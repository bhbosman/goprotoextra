package goprotoextra

import (
	"github.com/bhbosman/goCommsDefinitions"
)

type BaseMessageWrapper struct {
	toNext goCommsDefinitions.IAdder
}

func (self *BaseMessageWrapper) Adder() goCommsDefinitions.IAdder {
	return self.toNext
}

func (self *BaseMessageWrapper) SetNext(toNext goCommsDefinitions.IAdder) {
	if self != nil {
		self.toNext = toNext
	}
}

func (self *BaseMessageWrapper) ToNext(any interface{}) {
	if self != nil && self.toNext != nil {
		self.toNext.Add(any)
	}
}
