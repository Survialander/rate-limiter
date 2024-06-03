package service

import (
	"context"
	"encoding/json"
	"errors"
	"log"

	"github.com/Survialander/rate-limitter/configs"
	"github.com/Survialander/rate-limitter/internal/interfaces"
)

type CustomToken struct {
	Token                string
	RequestLimit         int
	TimeoutTimeInSeconds int
}

type LimitterService struct {
	limit                int
	timeoutTimeInSeconds int
	customTokens         []CustomToken
	provider             interfaces.LimiterCacheProvider
}

func GetLimitterService(config configs.Config, provider interfaces.LimiterCacheProvider) *LimitterService {
	var customTokens []CustomToken
	err := json.Unmarshal([]byte(config.CustomTokens), &customTokens)
	if err != nil {
		log.Print("unable to load custom tokens")
		customTokens = []CustomToken{}
	}

	return &LimitterService{
		limit:                config.RequestLimit,
		timeoutTimeInSeconds: config.TimeoutTimeInSeconds,
		customTokens:         customTokens,
		provider:             provider,
	}
}

func (l *LimitterService) CheckRequestLimit(key string) bool {
	var identifier string
	var blockTime int
	var requestLimit int

	token, err := l.getTokenData(key)

	if err != nil {
		identifier = key
		blockTime = l.timeoutTimeInSeconds
		requestLimit = l.limit
	} else {
		identifier = token.Token
		blockTime = token.TimeoutTimeInSeconds
		requestLimit = token.RequestLimit
	}

	blocked, _ := l.provider.BlockCheck(context.TODO(), identifier)

	if blocked {
		return true
	}

	keyRequests, _ := l.provider.GetLimitInfo(context.TODO(), identifier)

	if keyRequests >= requestLimit {
		l.provider.Block(context.TODO(), identifier, blockTime)

		return true
	}

	l.provider.SetLimitInfo(context.TODO(), identifier)

	return false
}

func (l *LimitterService) getTokenData(token string) (CustomToken, error) {
	for _, customToken := range l.customTokens {
		if customToken.Token == token {
			return customToken, nil
		}
	}

	return CustomToken{}, errors.New("invalid token")
}
