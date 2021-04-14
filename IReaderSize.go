package goprotoextra

import "io"

type IReaderSize interface {
	io.Reader
	Size() int
}
