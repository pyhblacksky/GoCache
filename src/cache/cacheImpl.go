package cache

import "sync"

//结构： map + 锁保护
type SimpleCache struct {
	c 		map[string][]byte
	mutex 	sync.RWMutex
	Stat
}

//Set方法
func (c *SimpleCache) Set(k string, v []byte) error{
	c.mutex.Lock()		//锁
	defer c.mutex.Unlock()
	tmp, exist := c.c[k]
	if exist{
		c.del(k, tmp)
	}
	c.c[k] = v
	c.add(k, v)
	return nil
}

