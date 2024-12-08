package main

import (
	"errors"
	"sync"
)

var (
	ErrKeyNotFound = errors.New("key not found")
)

type DataStorage interface {
	Set(key string, value string)
	Get(key string) (string, error)
}

type dataStorage struct {
	mutex sync.RWMutex
	data  map[string]string
}

func NewDataStorage() DataStorage {
	return &dataStorage{
		data:  make(map[string]string),
		mutex: sync.RWMutex{},
	}
}

func (ds *dataStorage) Set(key string, value string) {
	ds.mutex.Lock()
	defer ds.mutex.Unlock()

	ds.data[key] = value
}

func (ds *dataStorage) Get(key string) (string, error) {
	ds.mutex.RLock()
	defer ds.mutex.RUnlock()

	data, ok := ds.data[key]
	if !ok {
		return "", ErrKeyNotFound
	}
	return data, nil
}
