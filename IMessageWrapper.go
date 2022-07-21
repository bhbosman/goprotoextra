package goprotoextra

import "github.com/reactivex/rxgo/v2"

type IMessageWrapper interface {
	SetNext(toNext rxgo.NextFunc)
	ToNext(any interface{})
	Message() interface{}
}
