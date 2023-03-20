package main

import (
	"sync"
)

type syncMap struct {
	m  map[string]interface{}
	mu sync.RWMutex
}

func New() *syncMap {
	return &syncMap{
		m:  make(map[string]interface{}),
		mu: sync.RWMutex{},
	}
}

func (s *syncMap) Load(key string) (interface{}, bool) {
	s.mu.RLock() // блокируем на чтение, то есть если кто то будет записывать то mutex залочится
	val, ok := s.m[key]
	s.mu.RUnlock()
	return val, ok
}

func (s *syncMap) Save(key string, val interface{}) {
	s.mu.Lock()
	s.m[key] = val
	s.mu.Unlock()
}
