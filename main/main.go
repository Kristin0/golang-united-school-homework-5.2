package main
import "cache"
import "fmt"

func main() {
	c := cache.NewCache()
	fmt.Println(c)
	c.Put("key1", "value1")
	c.Put("key1","value2")
	c.Put("key2","value2")
	//fmt.Println(c.Keys())
	fmt.Println(c.Get("key2"))


	// b := cache.NewCache()
	// b.Put("key1", "value2")
	// fmt.Println(b.Get("key1"))
	
}