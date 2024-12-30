package storage

import (
	"errors"
	"sync"
)

var Store = struct {
	sync.RWMutex
	M map[string]string
}{M: make(map[string]string)}

var ErrorNoSuchKey = errors.New("no such key")

func Put(key, value string) error {
	Store.Lock()
	Store.M[key] = value
	Store.Unlock()

	return nil
}

func Get(key string) (string, error) {
	Store.RLock()
	value, ok := Store.M[key]
	Store.RUnlock()

	if !ok {
		return "", ErrorNoSuchKey
	}

	return value, nil
}

func Delete(key string) error {
	Store.Lock()
	delete(Store.M, key)
	Store.Unlock()

	return nil
}
