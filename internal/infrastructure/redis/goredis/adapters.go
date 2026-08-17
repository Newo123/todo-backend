package goredis

import (
	"errors"

	redispool "github.com/Newo123/todo-backend/internal/infrastructure/redis"
	"github.com/redis/go-redis/v9"
)

type goredisStringCmd struct {
	*redis.StringCmd
}

func (c goredisStringCmd) Bytes() ([]byte, error) {
	data, err := c.StringCmd.Bytes()
	if err != nil {
		return nil, mapErrors(err)
	}

	return data, nil
}

type goredisStatusCmd struct {
	*redis.StatusCmd
}

type goredisIntCmd struct {
	*redis.IntCmd
}

func mapErrors(err error) error {
	if errors.Is(err, redispool.NotFound) {
		return redispool.NotFound
	}

	return err
}
