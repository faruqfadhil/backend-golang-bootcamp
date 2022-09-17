package cache

type Cache interface {
	Get(key string) ([]byte, error)
	Set(key string, payload []byte, ttl int) error
	Del(keys ...string) error
}
