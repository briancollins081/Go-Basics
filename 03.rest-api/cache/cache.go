package cache

import (
	"net/http"
	"strings"
	"sync"
)

type response struct {
	header http.Header
	code   int
	body   []byte
}

type memCache struct {
	lock sync.RWMutex
	data map[string]response
}

var (
	cache = memCache{data: map[string]response{}}
)

func set(resource string, response *response) {
	cache.lock.Lock()
	if response == nil {
		delete(cache.data, resource)
	} else {
		cache.data[resource] = *response
	}
	cache.lock.Unlock()
}

func get(resource string) *response {
	cache.lock.RLock()
	res, ok := cache.data[resource]
	cache.lock.RUnlock()
	if ok {
		return &res
	}
	return nil
}

func copyHeader(src, dst http.Header) {
	for key, list := range src {
		for _, value := range list {
			dst.Add(key, value)
		}
	}
}

// MakeResource returns a valid resource
func MakeResource(r *http.Request) string {
	if r == nil {
		return ""
	}
	return strings.TrimSuffix(r.RequestURI, "/")
}

// Clear removes all the cache entries
func Clear() {
	cache.lock.Lock()
	cache.data = map[string]response{}
	cache.lock.Unlock()
}

// Drop removes a single cache entry
func Drop(res string) {
	set(res, nil)
}

func Serve(w http.ResponseWriter, r *http.Request) bool {
	if r == nil || r == nil {
		return false
	}
	if r.Header.Get("Cache-Control") == "no-chache" {
		return false
	}
	resp := get(MakeResource(r))
	if resp == nil {
		return false
	}

	copyHeader(resp.header, w.Header())
	w.WriteHeader(resp.code)
	if r.Method != http.MethodHead {
		w.Write(resp.body)
	}
	return true
}
