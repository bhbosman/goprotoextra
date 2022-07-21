package goprotoextra

import (
	"github.com/reactivex/rxgo/v2"
)

type BaseMessageWrapper struct {
	toNext rxgo.NextFunc
}

func (self *BaseMessageWrapper) SetNext(toNext rxgo.NextFunc) {
	if self != nil {
		self.toNext = toNext
	}
}

func (self *BaseMessageWrapper) ToNext(any interface{}) {
	if self != nil && self.toNext != nil {
		self.toNext(any)
	}
}
