package goprotoextra

import "context"

type IMessageWrapper interface {
	CancelCtx() context.Context
	CancelFunc() context.CancelFunc
	ToReactor(inline bool, any interface{}) error
	ToConnection(rws ReadWriterSize) error
	Message() interface{}
}
