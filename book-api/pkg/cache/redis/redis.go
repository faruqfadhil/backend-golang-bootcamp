package redis

import (
	"book-api/pkg/cache"

	"github.com/gomodule/redigo/redis"
)

type redisCacheEngine struct {
	cli *redis.Pool
}

func NewRedisCacheEngine(cli *redis.Pool) cache.Cache {
	return &redisCacheEngine{
		cli: cli,
	}
}

func (r *redisCacheEngine) Get(key string) ([]byte, error) {
	conn := r.cli.Get()
	defer conn.Close()

	data, err := redis.Bytes(conn.Do("GET", key))
	if err != nil {
		if err == redis.ErrNil {
			return nil, nil
		}
		return nil, err
	}

	return data, nil
}

func (r *redisCacheEngine) Set(key string, payload []byte, ttl int) error {
	conn := r.cli.Get()
	defer conn.Close()

	if ttl == 0 {
		_, err := conn.Do("SET", key, payload)
		if err != nil {
			return err
		}
	} else {
		_, err := conn.Do("SETEX", key, ttl, payload)
		if err != nil {
			return err
		}
	}

	return nil
}

func (r *redisCacheEngine) Del(keys ...string) error {
	conn := r.cli.Get()
	defer conn.Close()

	var ikeys []interface{}
	for _, e := range keys {
		ikeys = append(ikeys, e)
	}
	_, err := conn.Do("DEL", ikeys...)
	if err != nil {
		return err
	}

	return nil
}
