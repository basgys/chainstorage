package internal

import (
	"context"
	"errors"
)

type nopImpl struct{}

var _ DLQ = (*nopImpl)(nil)

func NewNop() DLQ {
	return &nopImpl{}
}

func (q *nopImpl) SendMessage(_ context.Context, _ *Message) error {
	return nil
}

func (q *nopImpl) ResendMessage(_ context.Context, _ *Message) error {
	return nil
}

func (q *nopImpl) ReceiveMessage(_ context.Context) (*Message, error) {
	return nil, errors.New("not implemented")
}

func (q *nopImpl) DeleteMessage(_ context.Context, _ *Message) error {
	return nil
}
