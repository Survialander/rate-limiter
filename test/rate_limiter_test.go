package test

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/Survialander/rate-limitter/configs"
	"github.com/Survialander/rate-limitter/internal/service"
	"github.com/Survialander/rate-limitter/test/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestShouldAllowWhenDoenstHaveBlockedKey(t *testing.T) {
	key := "192.0.0.68"
	provider := &mocks.PromiderMock{}
	config := configs.Config{
		RequestLimit:         5,
		TimeoutTimeInSeconds: 10,
	}

	provider.On("BlockCheck", context.TODO(), key).Return(false, nil)
	provider.On("GetLimitInfo", context.TODO(), key).Return(1, nil)
	provider.On("SetLimitInfo", context.TODO(), key).Return(nil)

	service := service.GetLimitterService(config, provider)
	result := service.CheckRequestLimit(key)

	assert.Equal(t, false, result)
	provider.AssertNumberOfCalls(t, "BlockCheck", 1)
	provider.AssertNumberOfCalls(t, "GetLimitInfo", 1)
	provider.AssertNumberOfCalls(t, "SetLimitInfo", 1)
	provider.AssertNumberOfCalls(t, "Block", 0)
}

func TestShouldBlockWhenHasABlockedKey(t *testing.T) {
	key := "192.0.0.68"
	provider := &mocks.PromiderMock{}
	config := configs.Config{
		RequestLimit:         5,
		TimeoutTimeInSeconds: 10,
	}

	provider.On("BlockCheck", context.TODO(), key).Return(true, nil)

	service := service.GetLimitterService(config, provider)
	result := service.CheckRequestLimit(key)

	assert.Equal(t, true, result)
	provider.AssertNumberOfCalls(t, "BlockCheck", 1)
	provider.AssertNumberOfCalls(t, "GetLimitInfo", 0)
	provider.AssertNumberOfCalls(t, "SetLimitInfo", 0)
	provider.AssertNumberOfCalls(t, "Block", 0)
}

func TestShouldBlockWhenHasReachedTheLimit(t *testing.T) {
	key := "192.0.0.68"
	provider := &mocks.PromiderMock{}
	config := configs.Config{
		RequestLimit:         5,
		TimeoutTimeInSeconds: 10,
	}

	provider.On("BlockCheck", context.TODO(), key).Return(false, nil)
	provider.On("GetLimitInfo", context.TODO(), key).Return(5, nil)
	provider.On("Block", mock.Anything, mock.Anything, mock.Anything).Return(nil)

	service := service.GetLimitterService(config, provider)
	result := service.CheckRequestLimit(key)

	assert.Equal(t, true, result)
	provider.AssertCalled(t, "Block", context.TODO(), key, config.TimeoutTimeInSeconds)
	provider.AssertNumberOfCalls(t, "BlockCheck", 1)
	provider.AssertNumberOfCalls(t, "GetLimitInfo", 1)
	provider.AssertNumberOfCalls(t, "Block", 1)
	provider.AssertNumberOfCalls(t, "SetLimitInfo", 0)
}

func TestShouldUseValidTokenCustomConfig(t *testing.T) {
	key := "abcd"
	provider := &mocks.PromiderMock{}
	tokeString, _ := json.Marshal([]service.CustomToken{
		{
			Token:                "abcd",
			RequestLimit:         30,
			TimeoutTimeInSeconds: 40,
		},
	})
	config := configs.Config{
		RequestLimit:         5,
		TimeoutTimeInSeconds: 10,
		CustomTokens:         string(tokeString),
	}

	provider.On("BlockCheck", mock.Anything, mock.Anything).Return(false, nil)
	provider.On("GetLimitInfo", mock.Anything, mock.Anything).Return(25, nil)
	provider.On("SetLimitInfo", mock.Anything, mock.Anything).Return(nil)

	service := service.GetLimitterService(config, provider)
	result := service.CheckRequestLimit(key)

	assert.Equal(t, false, result)
	provider.AssertCalled(t, "BlockCheck", context.TODO(), "abcd")
	provider.AssertCalled(t, "SetLimitInfo", context.TODO(), "abcd")
	provider.AssertCalled(t, "GetLimitInfo", context.TODO(), "abcd")
	provider.AssertCalled(t, "SetLimitInfo", context.TODO(), "abcd")
}

func TestShouldBlockCustomTokenIfItReachTheLimit(t *testing.T) {
	key := "abcd"
	provider := &mocks.PromiderMock{}
	tokeString, _ := json.Marshal([]service.CustomToken{
		{
			Token:                "abcd",
			RequestLimit:         30,
			TimeoutTimeInSeconds: 40,
		},
	})
	config := configs.Config{
		RequestLimit:         5,
		TimeoutTimeInSeconds: 10,
		CustomTokens:         string(tokeString),
	}

	provider.On("BlockCheck", mock.Anything, mock.Anything).Return(false, nil)
	provider.On("GetLimitInfo", mock.Anything, mock.Anything).Return(30, nil)
	provider.On("SetLimitInfo", mock.Anything, mock.Anything).Return(nil)
	provider.On("Block", mock.Anything, mock.Anything, mock.Anything).Return(nil)

	service := service.GetLimitterService(config, provider)
	result := service.CheckRequestLimit(key)

	assert.Equal(t, true, result)
	provider.AssertCalled(t, "BlockCheck", context.TODO(), "abcd")
	provider.AssertCalled(t, "GetLimitInfo", context.TODO(), "abcd")
	provider.AssertCalled(t, "Block", context.TODO(), "abcd", 40)
}
