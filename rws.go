package goprotoextra

import (
	"errors"
	"github.com/reactivex/rxgo/v2"
)

type ReadWriterSize IReadWriterSize
type ReaderSize IReaderSize

type ToReactorFunc func(inline bool, any interface{}) error
type ToConnectionFunc func(rws ReadWriterSize) error

//type ConnectionReactorHandler func(i interface{})
//type ErrorStateFunc func() (bool, error)

func NewBaseMessageWrapper(
	toReactor rxgo.NextFunc,
	toConnection rxgo.NextFunc,
) BaseMessageWrapper {
	return BaseMessageWrapper{
		toReactor:    toReactor,
		toConnection: toConnection,
	}
}

var NullValueError = errors.New("NullValueError")
