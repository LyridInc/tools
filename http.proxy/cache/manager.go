package cache

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"sync"
	"time"
)

type ProxyMap struct {
	Source string `json:"source"`
	Target string `json:"target"`
}

type ProxyMapCache struct {
	Cache map[string]string

	mu sync.Mutex
}

var instance *ProxyMapCache
var once sync.Once

func GetInstance() *ProxyMapCache {
	once.Do(func() {
		instance = &ProxyMapCache{}
	})
	return instance
}

func (manager *ProxyMapCache) Init() {
	manager.getUpdate()
	go manager.runUpdateThread()
}

func (manager *ProxyMapCache) getUpdate() {

	update := make(map[string]string)

	var raw_map []byte = nil

	if os.Getenv("USE_MAP") == "FILE" {
		f, err := os.ReadFile(os.Getenv("MAP_FILE"))
		if err != nil {
			log.Println(err)
		} else {
			raw_map = f
		}
	} else if os.Getenv("USE_MAP") == "URL" {
		resp, err := http.Get(os.Getenv("MAP_URL"))
		if err != nil {
			log.Println(err)
		} else {
			//We Read the response body on the line below.
			raw_map, err = io.ReadAll(resp.Body)
			if err != nil {
				log.Println(err)
			}
		}
	} else {
		log.Println("No compatible map")
	}

	if raw_map != nil {
		var proxy_map []ProxyMap
		err := json.Unmarshal(raw_map, &proxy_map)

		if err != nil {
			log.Println(err)
		} else {
			for _, proxyMap := range proxy_map {
				update[proxyMap.Source] = proxyMap.Target
			}
			manager.mu.Lock()
			manager.Cache = update
			manager.mu.Unlock()
		}
	}
}

func (manager *ProxyMapCache) runUpdateThread() {

	for {
		time.Sleep(time.Second * 30)
		manager.getUpdate()
		// pull current values from the internet or local
	}
}

func (cache *ProxyMapCache) GetCache() map[string]string {
	cache.mu.Lock()
	defer cache.mu.Unlock()

	values := make(map[string]string)

	for s, s2 := range cache.Cache {
		values[s] = s2
	}

	return values
}
