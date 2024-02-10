package client

import (
	"strconv"

	"github.com/mastik5h/LLD/cache/models"
	"github.com/mastik5h/LLD/cache/services"
)

func CreateCache(cache_size string, eviction_policy string, default_ttl string) (models.ICache[string, string], string) {
	cacheSize, err := strconv.Atoi(cache_size)
	if err != nil {
		cacheSize = 10
	}
	var evictionPolicy models.EvictionPolicyType
	switch eviction_policy {
	case "LRU":
		evictionPolicy = models.LRU_EvicitionPolicyType
	case "LFU":
		evictionPolicy = models.LFU_EvicitionPolicyType
	default:
		evictionPolicy = models.LRU_EvicitionPolicyType
	}
	ttl, err := strconv.Atoi(default_ttl)
	if err != nil {
		ttl = 60
	}
	cc, err := services.InitializeCacheClient[string, string](cacheSize, evictionPolicy, ttl)
	if err != nil {
		return nil, err.Error()
	}
	return cc, ""
}

func SetEntry(cache models.ICache[string, string], key string, value string, time_to_live string) string {
	ttl := new(int)
	if time_to_live == "" {
		v, err := strconv.Atoi(time_to_live)
		if err == nil {
			*ttl = v
		} else {
			ttl = nil
		}
	}
	err := cache.SetEntry(key, value, ttl)
	if err != nil {
		return err.Error()
	}
	return "Success!!"
}

func GetEntry(cache models.ICache[string, string], key string) (string, string) {
	v, err := cache.GetEntry(key)
	if err != nil {
		return "", err.Error()
	}
	return v, ""
}
