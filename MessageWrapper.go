package goprotoextra

type MessageWrapper struct {
	BaseMessageWrapper
	Data interface{}
}

func NewMessageWrapper(data interface{}) *MessageWrapper {
	return &MessageWrapper{
		BaseMessageWrapper: NewBaseMessageWrapper(),
		Data:               data,
	}
}

func (self *MessageWrapper) Message() interface{} {
	return self.Data
}

func (self *MessageWrapper) messageWrapper() IMessageWrapper {
	return self
}
