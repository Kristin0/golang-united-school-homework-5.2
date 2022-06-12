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

// func (c Cache) clear() Cache {
// 	for _, key := range c.Keys() {
// 		if time.Now().After(c.m[key].t) {
// 			delete(c.m, key)
// 		}
// 	}
// 	return c
// }

func (c Cache) Get(key string) (string, bool) {
	if time.Now().After(c.m[key].t){
		delete(c.m, key)
		return "", false
	}
	if _, ok := c.m[key]; ok   {
		return c.m[key].v, c.b
	}else {
		return "", false
	}
	
}

func (c *Cache) Put(key, value string) {
	if _, ok := c.m[key]; ok {
		c.m[key] = myType{v: value, t: time.Unix(1<<63-62135596801, 999999999)}
	}else{
		c.m[key] = myType{v: value, t: time.Unix(1<<63-62135596801, 999999999)}
		c.b = true
	}
}

func (c Cache) Keys() []string {
	var s[]string
	for key, _ := range c.m {
		if time.Now().After(c.m[key].t){
			delete(c.m, key)
		}
		s = append(s, key)
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
