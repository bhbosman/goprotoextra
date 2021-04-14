package goprotoextra

import (
	"context"
	"errors"
)

type ReadWriterSize IReadWriterSize
type ReaderSize IReaderSize

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
