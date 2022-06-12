package cache

import "time"

type myType struct {
	t time.Time
	v string
}
type Cache struct {
	m map[string] myType
	b bool
	
}

func NewCache() Cache {
	return Cache{m: make(map[string]myType), b: false}
}

func (c Cache) clear() Cache {
	for _, key := range c.Keys() {
		if time.Now().After(c.m[key].t) {
			delete(c.m, key)
		}
	}
	return c
}

func (c Cache) Get(key string) (string, bool) {
	c = c.clear()
	if _, ok := c.m[key]; ok   {
		return c.m[key].v, c.b
	}else {
		return c.m[key].v, false
	}
	
}

func (c *Cache) Put(key, value string) {
	time := time.Now()
	if _, ok := c.m[key]; ok {
		c.m[key] = myType{v: value, t: time}
	}else{
		c.m[key] = myType{v: value, t: time}
		c.b = true
	}
}

func (c Cache) Keys() []string {
	var s[]string
	for k, _ := range c.m {
		s = append(s, k)
	}
	return s
}

func (c *Cache) PutTill(key, value string, deadline time.Time) {
	if _, ok := c.m[key]; ok {
		c.m[key] = myType{v: value, t: deadline}
	}else{
		c.m[key] = myType{v: value, t: deadline}
		c.b = true
	}


	
}
