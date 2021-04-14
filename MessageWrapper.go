package goprotoextra

import "context"

type MessageWrapper struct {
	BaseMessageWrapper
	Data interface{}
}

func NewMessageWrapper(
	cancelCtx context.Context,
	cancelFunc context.CancelFunc,
	toReactor ToReactorFunc,
	toConnection ToConnectionFunc,
	data interface{}) *MessageWrapper {
	return &MessageWrapper{
		BaseMessageWrapper: NewBaseMessageWrapper(
			cancelCtx,
			cancelFunc,
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
