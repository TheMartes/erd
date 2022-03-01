package cache

import (
	"errors"
	"log"
	"os"

	"github.com/themartes/erd/utils/file"
)

type ReplicationCache []string

var cachePath string = "/tmp/cache"
var cache ReplicationCache

func GetCacheFromMemory() ReplicationCache {
	return cache
}

func WriteToCache(key string) bool {
	if cacheContains(key) {
		return false
	}

	cache = append(cache, key)
	return true
}

func WriteToDisk() {
	file.WriteLines(cache, cachePath)
}

func PopFromDisk() {
	values, err := file.ReadLines(cachePath)

	// TODO if file doesn't exists load data from elastic

	if err != nil {
		log.Fatalf("Error while popping cache to memory: %s", err)
	}

	for _, k := range values {
		cache = append(cache, k)
	}
}

func ReplicationCacheExists() bool {
	_, err := os.Stat(cachePath)

	if errors.Is(err, os.ErrNotExist) {
		return false
	} else {
		log.Fatalf("Error while checking for existent cache. Error: %s", err)
	}

	return true
}

func FlushCache() {
	err := os.Remove(cachePath)

	if err != nil {
		log.Fatalf("Error while flushing cache: %s", err)
	}
}

func cacheContains(key string) bool {
	for _, k := range cache {
		if k == key {
			return true
		}
	}

	return false
}
