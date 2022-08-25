package event

import (
	"context"
	"sync"
)

type Event string

func newBroker[T any](e Event) *broker[T] {
	return &broker[T]{
		e:           e,
		subscribers: []func(context.Context, T){},
		wg:          &sync.WaitGroup{},
	}
}

type broker[T any] struct {
	e           Event
	subscribers []func(context.Context, T)

	wg *sync.WaitGroup
}

func (b *broker[T]) Publish(ctx context.Context, event T) {
	for _, fn := range b.subscribers {
		b.wg.Add(1)
		go func(callSubscriber func(context.Context, T)) {
			defer b.wg.Done()
			callSubscriber(ctx, event)
		}(fn)
	}
}

func (b *broker[T]) Subscribe(fn func(context.Context, T)) {
	b.subscribers = append(b.subscribers, fn)
}

func (b *broker[T]) Wait(ctx context.Context) {
	b.wg.Wait()
}
