package registry

import (
	"fmt"
	"sync"
)

var (
	mu       sync.RWMutex
	couriers = make(map[string]interface{})
)

func Register(name string, trackerInterface interface{}) {
	mu.Lock()
	defer mu.Unlock()
	couriers[name] = trackerInterface
}

func PrintRegistry() {
	mu.RLock()
	defer mu.RUnlock()
	fmt.Println(couriers)
}
