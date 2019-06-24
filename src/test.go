package main

import (
	"./cache"
	"fmt"
)

/**
 * @Author: pyh
 * @Date: 2019/6/24 15:11
 * @Function:
 */

func main() {
	//测试
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
