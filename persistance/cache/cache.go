package cache

type ReplicationCache []string

var cache ReplicationCache

func GetCache() ReplicationCache {
	return cache
}

func WriteToCache(key string) bool {
	if cacheContains(key) {
		return false
	}

	cache = append(cache, key)
	return true
}

func cacheContains(key string) bool {
	for _, k := range cache {
		if k == key {
			return true
		}
	}

	return false
}
