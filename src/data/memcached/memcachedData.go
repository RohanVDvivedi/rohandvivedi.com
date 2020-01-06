package memcached

import (
	"github.com/bradfitz/gomemcache/memcache"
)

var memcacheClient *memcache.Client= nil;

func Initialize() {
	memcacheClient = memcache.New("localhost:11211");
}

func Get(key string) ([]byte) {
	itemG, err := memcacheClient.Get(key);
	if(err != nil && err != memcache.ErrCacheMiss) {
		panic(err.Error());
	} else if (err == memcache.ErrCacheMiss) {
		return nil;
	} else {
		return itemG.Value;
	}
}

// Expiration = 0, means no expiration
func Set(key string, Value []byte, Expiration int32) {
	err := memcacheClient.Set(&memcache.Item{
		Key: key,
		Value: Value,
		Expiration: Expiration,
	});
	if(err != nil) {
		panic(err.Error());
	}
}

// returns true, if the element was deleted
func Del(key string) bool {
	err := memcacheClient.Delete(key);
	if(err != nil && err != memcache.ErrCacheMiss) {
		panic(err.Error());
	} else if (err == memcache.ErrCacheMiss) {
		return false;
	} else {
		return true;
	}
}