package main

import (
	"github.com/brianvoe/gofakeit/v7"
	"github.com/stretchr/testify/assert"
	"sync"
	"testing"
)

func TestSetAndGet(t *testing.T) {
	ds := NewDataStorage()

	tt := make([]struct {
		Key   string `fake:"{sentence:2}"`
		Value string `fake:"{sentence:5}"`
	}, 100)
	gofakeit.Slice(&tt)

	for _, tc := range tt {
		ds.Set(tc.Key, tc.Value)
		got, err := ds.Get(tc.Key)
		assert.NoError(t, err)
		assert.Equal(t, tc.Value, got)
	}
}

func TestGetNoExistKey(t *testing.T) {
	ds := NewDataStorage()
	_, err := ds.Get("hello")
	assert.Error(t, err)
}

func TestConcurrent(t *testing.T) {
	ds := NewDataStorage()
	wg := sync.WaitGroup{}
	tt := make([]struct {
		Key   string `fake:"{sentence:2}"`
		Value string `fake:"{sentence:5}"`
	}, 100)
	gofakeit.Slice(&tt)

	for _, tc := range tt {
		wg.Add(1)
		go func() {
			defer wg.Done()
			ds.Set(tc.Key, tc.Value)
		}()
	}

	for _, tc := range tt {
		wg.Add(1)
		go func() {
			defer wg.Done()
			value, err := ds.Get(tc.Key)
			assert.NoError(t, err)
			assert.Equal(t, tc.Value, value)
		}()
	}

	wg.Wait()
}
