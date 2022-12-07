package cache

type FIFO struct {
	data []*Entry
}

func NewFIFO(capacity int) *FIFO {
	return &FIFO{
		data: make([]*Entry, capacity),
	}
}

func (c *FIFO) Get(key int) *Entry {
	for _, v := range c.data {
		if v.Key == key {
			return v
		}
	}

	return nil
}

func (c *FIFO) Put(key, value int) {
	c.data = append(c.data[1:], &Entry{
		Key:   key,
		Value: value,
	})
}

func (c *FIFO) Data() []*Entry {
	return c.data
}
