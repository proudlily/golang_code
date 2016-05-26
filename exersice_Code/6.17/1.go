package main

import "fmt"

type cache struct {
	Data map[string]interface{}
}

func (c *cache) Set(key string, value interface{}) {
	c.Data[key] = value
}

func (c *cache) Get(key string) (value interface{}) {
	value = c.Data[key]
	return
}

func main() {
	ca := cache{Data: map[string]interface{}{}}
	ca.Data["key"] = "aaa"
	ca.Set("key", "sasasasa")
	value := ca.Get("key")
	fmt.Println(value)
}
