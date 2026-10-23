package main

import (
	"fmt"

	"github.com/cristiantebaldi/singleton-go/cache"
	"github.com/cristiantebaldi/singleton-go/locator"
)

func main() {
	singletonCache1 := cache.GetInstance()
	singletonCache2 := cache.GetInstance()

	singletonCache1.Set("user", "João")

	fmt.Println(singletonCache1.Get("user"))
	fmt.Println(singletonCache2.Get("user"))

	sl := locator.GetInstance()
	locator.Register[cache.Cache](sl, "cacheService", cache.NewMemoryCache())

	cache1, _ := locator.Get[cache.Cache](sl, "cacheService")
	cache2, _ := locator.Get[cache.Cache](sl, "cacheService")

	cache1.Set("user", "João")

	v1, _ := cache1.Get("user")
	v2, _ := cache2.Get("user")
	fmt.Println(v1, v2)
}
