package main

import (
	"errors"
	"io/ioutil"
	"os"
	"time"

	json "github.com/json-iterator/go"
)

// Setup setups the cache
func (c *Cache) Setup(fname string, timeout time.Duration) error {
	c.cacheFile = fname
	c.timeout = timeout
	_, err := os.Create(c.cacheFile)
	if err != nil {
		return err
	}
	return nil
}

// Write write into the cache
func (c *Cache) Write(in []byte) error {
	jsonString, _ := json.Marshal(in)
	wErr := ioutil.WriteFile(c.cacheFile, jsonString, 0777)
	if wErr != nil {
		return wErr
	}
	return nil
}

// Read from the cache
func (c *Cache) Read() ([]interface{}, error) {
	var content []interface{}
	cache, err := ioutil.ReadFile(c.cacheFile)
	if err != nil {
		return nil, err
	}
	jsonErr := json.Unmarshal(cache, &content)
	if jsonErr != nil {
		return nil, jsonErr
	}
	c.size = len(content)
	return content, nil
}

// Expired checks if the cache is still valid
func (c *Cache) Expired(t time.Time) bool {
	return time.Now().Sub(t) > c.timeout
}

// Age gives back the cache age
func (c *Cache) Age() time.Time {
	fileinfo, _ := os.Stat(c.cacheFile)
	return fileinfo.ModTime()
}

// Reset the cache
func (c *Cache) Reset() error {
	rmErr := os.Remove(c.cacheFile)
	if rmErr != nil {
		return rmErr
	}
	return nil
}

// Exists check if there's a cache
func (c *Cache) Exists() bool {
	if _, err := os.Stat(c.cacheFile); err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return false
		}
	}
	return true
}

// Size returns the cache size
func (c *Cache) Size() int {
	return c.size
}
