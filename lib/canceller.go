package lib

import "context"

type Canceller struct {
	Ctx    context.Context
	Cancel func()
}

func NewCanceller() Canceller {
	ctx, cancel := context.WithCancel(context.Background())
	return Canceller{ctx, cancel}
}
