package goprotoextra

import (
	"github.com/reactivex/rxgo/v2"
)

type BaseMessageWrapper struct {
	toReactor    rxgo.NextFunc
	toConnection rxgo.NextFunc
}

func (self *BaseMessageWrapper) ToReactor(any interface{}) {
	if self != nil && self.toReactor != nil {
		self.toReactor(any)
	}
}

func (self *BaseMessageWrapper) ToConnection(any interface{}) {
	if self != nil && self.toConnection != nil {
		self.toConnection(any)
	}
}
