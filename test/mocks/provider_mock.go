package mocks

import (
	"context"

	"github.com/stretchr/testify/mock"
)

type PromiderMock struct {
	mock.Mock
}

func (m *PromiderMock) GetLimitInfo(ctx context.Context, key string) (int, error) {
	args := m.Called(ctx, key)
	return args.Int(0), args.Error(1)
}

func (m *PromiderMock) SetLimitInfo(ctx context.Context, key string) error {
	args := m.Called(ctx, key)
	return args.Error(0)
}

func (m *PromiderMock) Block(ctx context.Context, key string, timeout int) error {
	args := m.Called(ctx, key, timeout)
	return args.Error(0)
}

func (m *PromiderMock) BlockCheck(ctx context.Context, key string) (bool, error) {
	args := m.Called(ctx, key)
	return args.Bool(0), args.Error(1)
}
