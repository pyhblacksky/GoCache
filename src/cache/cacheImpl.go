package cache

import "sync"

//结构： map + 锁保护
type SimpleCache struct {
	c     map[string][]byte
	mutex sync.RWMutex
	Stat
}

//Set方法
func (c *SimpleCache) Set(k string, v []byte) error {
	c.mutex.Lock()         //锁
	defer c.mutex.Unlock() //保证mutex对象锁住后的自动释放
	tmp, exist := c.c[k]
	if exist {
		c.del(k, tmp)
	}
	c.c[k] = v
	c.add(k, v)
	return nil
}

//Get方法
func (c *SimpleCache) Get(k string) ([]byte, error) {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	return c.c[k], nil
}

//Del方法
func (c *SimpleCache) Del(k string) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	v, exist := c.c[k]
	if exist {
		delete(c.c, k)
		c.del(k, v)
	}
	return nil
}

//获取状态
func (c *SimpleCache) GetStat() Stat {
	return c.Stat
}

//在内存中新建新缓存
func newInMemoryCache() *SimpleCache {
	return &SimpleCache{make(map[string][]byte), sync.RWMutex{}, Stat{}}
}
