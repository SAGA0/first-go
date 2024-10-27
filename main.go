package main

import (
	simplecache "first_go/simple_cache"
	"fmt"
)

func main() {

	cache := simplecache.NewCache()

	cache.Set("Возраст", 25)
	cache.Set("Auto", "Red")

	if age, exist := cache.Get("Возраст"); exist {
		fmt.Println("Возраст", age)
	} else {
		fmt.Println("Возраста нет")
	}

	cache.Delete("Auto")

	if _, exist := cache.Get("Auto"); !exist {
		fmt.Println("Авто не найдено")
	}

}
