package cache

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
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
	resp, err := http.Get("https://opensheet.elk.sh/1hD-H-dSvBP4meyZvCnlEeLZJ3rRCeGhO38GBkgMNrlY/map")
	if err != nil {
		log.Println(err)
	} else {
		//We Read the response body on the line below.
		body, err := ioutil.ReadAll(resp.Body)

		if err != nil {
			log.Println(err)
		} else {
			var proxy_map []ProxyMap
			err = json.Unmarshal(body, &proxy_map)

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
