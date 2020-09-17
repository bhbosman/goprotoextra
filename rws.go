package goprotoextra

import (
	"context"
	"errors"
	"io"
)

type IReadWriterSize interface {
	io.ReadWriter
	Size() int
}

type IReaderSize interface {
	io.Reader
	Size() int
}

type ReadWriterSize IReadWriterSize
type ReaderSize IReaderSize

type IMessageWrapper interface {
	CancelCtx() context.Context
	CancelFunc() context.CancelFunc
	ToReactor(inline bool, any interface{}) error
	ToConnection(rws ReadWriterSize) error
	Message() interface{}
}
type BaseMessageWrapper struct {
	cancelCtx    context.Context
	cancelFunc   context.CancelFunc
	toReactor    ToReactorFunc
	toConnection func(rws ReadWriterSize) error
}
type ToReactorFunc func(inline bool, any interface{}) error
type ToConnectionFunc func(rws ReadWriterSize) error
type ConnectionReactorHandler func(i interface{})
type ErrorStateFunc func() (bool, error)

func NewBaseMessageWrapper(
	cancelCtx context.Context,
	cancelFunc context.CancelFunc,
	toReactor ToReactorFunc,
	toConnection ToConnectionFunc) BaseMessageWrapper {
	return BaseMessageWrapper{
		cancelCtx:    cancelCtx,
		cancelFunc:   cancelFunc,
		toReactor:    toReactor,
		toConnection: toConnection,
	}
}

var NullValueError = errors.New("NullValueError")

func (self *BaseMessageWrapper) ToReactor(inline bool, any interface{}) error {
	if self != nil && self.toReactor != nil {
		if self.cancelCtx != nil {
			err := self.cancelCtx.Err()
			if err != nil {
				return err
			}
		}
		return self.toReactor(inline, any)
	}
	return NullValueError
}

func (self *BaseMessageWrapper) ToConnection(rws ReadWriterSize) error {
	if self != nil && self.toConnection != nil {
		if self.cancelCtx != nil {
			err := self.cancelCtx.Err()
			if err != nil {
				return err
			}
		}
		return self.toConnection(rws)
	}
	return NullValueError
}

func (self *BaseMessageWrapper) CancelCtx() context.Context {
	if self != nil && self.cancelCtx != nil {
		return self.cancelCtx
	}
	return nil
}

func (self *BaseMessageWrapper) CancelFunc() context.CancelFunc {
	if self != nil && self.cancelFunc != nil {
		return self.cancelFunc
	}
	return nil
}

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
