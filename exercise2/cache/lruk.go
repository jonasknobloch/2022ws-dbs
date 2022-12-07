package cache

// TODO implement (current implementation equals LRU-1)

type LRUK struct {
	data []*Entry
}

func NewLRUK(capacity int) *LRUK {
	return &LRUK{
		data: make([]*Entry, capacity),
	}
}

func (c *LRUK) Get(key int) *Entry {
	var p int
	var e *Entry

	for i, v := range c.data {
		if v.Key == key {
			e = v
			p = i

			break
		}
	}

	if e == nil {
		return nil
	}

	c.data = append(append(c.data[:p], c.data[p+1:]...), e)

	return e
}

func (c *LRUK) Put(key, value int) {
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

func (c *LRUK) Data() []*Entry {
	return c.data
}
