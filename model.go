package main

import (
	"time"
)

// Cacher is the cache interface
type Cacher interface {
	Setup(string, time.Duration) error
	Write([]byte) error
	Read() ([]interface{}, error)
	Expired(time.Time) bool
	Age() time.Time
	Reset() error
	Exists() bool
	Size() int
}

// Cache is the cAPS cache
type Cache struct {
	cacheFile string
	timeout   time.Duration
	size      int
}
