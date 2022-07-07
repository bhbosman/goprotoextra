package goprotoextra

import "github.com/reactivex/rxgo/v2"

type MessageWrapper struct {
	BaseMessageWrapper
	Data interface{}
}

func NewMessageWrapper(
	toReactor rxgo.NextFunc,
	toConnection rxgo.NextFunc,
	data interface{}) *MessageWrapper {
	return &MessageWrapper{
		BaseMessageWrapper: NewBaseMessageWrapper(
			toReactor,
			toConnection),
		Data: data,
	}
}

func (self *MessageWrapper) Message() interface{} {
	return self.Data
}

func (self *MessageWrapper) messageWrapper() IMessageWrapper {
	return self
}
