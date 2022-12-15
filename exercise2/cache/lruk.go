package cache

type LRUK struct {
	data []*Entry
	log  map[*Entry][]int
	t    int // time
	k    int // references
}

func NewLRUK(capacity int, k int) *LRUK {
	return &LRUK{
		data: make([]*Entry, capacity),
		log:  make(map[*Entry][]int),
		k:    k,
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

	if c.log[e] == nil {
		c.log[e] = make([]int, c.k)
	}

	if len(c.log[e]) < c.k {
		c.log[e] = append(c.log[e], c.t)
	} else {
		c.log[e] = append(c.log[e][1:], c.t)
	}

	c.t++

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

	var p int // position
	var m int // max delta

	p = len(c.data) - 1

	for i, v := range c.data {
		if len(c.log[v]) < c.k {
			continue
		}

		d := c.t - c.log[v][0]

		if d > m {
			m = d
			p = i
		}
	}

	c.data[p] = e
}

func (c *LRUK) Data() []*Entry {
	return c.data
}

func (c *LRUK) Log() [][2]int {
	log := make([][2]int, c.t)

	for e, t := range c.log {
		for _, v := range t {
			log[v] = [2]int{e.Key, e.Value}
		}
	}

	return log
}
