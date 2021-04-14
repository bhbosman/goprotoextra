package goprotoextra

import "io"

type IReadWriterSize interface {
	io.ReadWriter
	Size() int
}
