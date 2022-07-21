package goprotoextra

import (
	"errors"
)

type ReadWriterSize IReadWriterSize
type ReaderSize IReaderSize

type ToReactorFunc func(inline bool, any interface{}) error
type ToConnectionFunc func(rws ReadWriterSize) error

//type ConnectionReactorHandler func(i interface{})
//type ErrorStateFunc func() (bool, error)

func NewBaseMessageWrapper() BaseMessageWrapper {
	return BaseMessageWrapper{}
}

var NullValueError = errors.New("NullValueError")
