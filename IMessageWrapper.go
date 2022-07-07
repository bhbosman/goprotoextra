package goprotoextra

type IMessageWrapper interface {
	ToReactor(any interface{})
	ToConnection(any interface{})
	Message() interface{}
}
