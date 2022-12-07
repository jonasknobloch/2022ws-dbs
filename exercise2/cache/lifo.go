package cache

type LIFO struct {
	data []*Entry
}

func NewLIFO(capacity int) *LIFO {
	return &LIFO{
		data: make([]*Entry, capacity),
	}
}

func (c *LIFO) Get(key int) *Entry {
	for _, v := range c.data {
		if v.Key == key {
			return v
		}
	}

	return nil
}

func (c *LIFO) Put(key, value int) {
	e := &Entry{
		Key:   key,
		Value: value,
	}

	for i, v := range c.data {
		if v == nil {
			c.data[i] = e
			return
		}
	}

	c.data[len(c.data)-1] = e
}

func (c *LIFO) Data() []*Entry {
	return c.data
}
