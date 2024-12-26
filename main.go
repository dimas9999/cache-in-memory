package cacheinmemory

type CacheInMemory struct {
	data map[string]interface{}
}

func New() CacheInMemory {
	return CacheInMemory{make(map[string]interface{})}
}

func (c *CacheInMemory) Set(key string, value interface{}) {
	c.data[key] = value
}

func (c CacheInMemory) Get(key string) interface{} {
	return c.data[key]
}

func (c CacheInMemory) Delete(key string) {
	delete(c.data, key)
}
