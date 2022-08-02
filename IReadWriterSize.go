package goprotoextra

import "io"

type IReadWriterSize interface {
	io.ReadWriter
	Size() int
	ReadTypeCode() (uint32, error)
}
