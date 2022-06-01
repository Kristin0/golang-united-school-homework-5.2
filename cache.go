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
	return Cache{}
}

func (c Cache) Get(key string) (string, bool) {
	return c.m[key].v, c.b
}

func (c *Cache) Put(key, value string) {
	if thisKey, ok := c.m[key]; ok {   // https://webapplicationconsultant.com/go-lang/cannot-assign-to-struct-field-in-map/
		thisKey.v = value
		c.m[key] = thisKey 
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
	if thisKey, ok := c.m[key]; ok {   // https://webapplicationconsultant.com/go-lang/cannot-assign-to-struct-field-in-map/
		thisKey.v = value
		c.m[key] = thisKey 
	}
	if thisTime, ok := c.m[key]; ok {   // https://webapplicationconsultant.com/go-lang/cannot-assign-to-struct-field-in-map/
		thisTime.t = deadline
		c.m[key] = thisTime 
	}


	
}
