package main

import (
	"./cache"
	"fmt"
)

//主程序
func main() {

	c := cache.New("inmemory")
	k, v := "sola", []byte{'a', 'b', 'c', 'd', 'e', 'f'}

	c.Set(k, v)
	tmp, _ := c.Get(k)
	fmt.Println("Key:", k, " value: ", tmp)

	c.Del(k)
	tmp, _ = c.Get(k)
	fmt.Println("Key:", tmp, " value: ", tmp)

	fmt.Println("hello")
}
