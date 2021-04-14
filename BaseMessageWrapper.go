package goprotoextra

import "context"

type BaseMessageWrapper struct {
	cancelCtx    context.Context
	cancelFunc   context.CancelFunc
	toReactor    ToReactorFunc
	toConnection func(rws ReadWriterSize) error
}

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
