package _16_iterator

// Iterator 提供顺序遍历，不暴露底层结构。
type Iterator interface {
	Next() bool
	Value() string
}

// SliceIterator 基于下标遍历切片。
type SliceIterator struct {
	data []string
	idx  int
}

func (it *SliceIterator) Next() bool {
	if it.idx >= len(it.data) {
		return false
	}
	it.idx++
	return true
}

func (it *SliceIterator) Value() string {
	return it.data[it.idx-1]
}

// Collection 聚合对象。
type Collection struct {
	data []string
}

func NewCollection(items ...string) *Collection {
	return &Collection{data: items}
}

// Iterator 返回新的遍历器，每次从头开始。
func (c *Collection) Iterator() Iterator {
	return &SliceIterator{data: c.data}
}

// Iter 用 channel 遍历，消费方可用 for range 读取（Go 惯用的替代方案）。
func (c *Collection) Iter() <-chan string {
	ch := make(chan string)
	go func() {
		defer close(ch)
		for _, s := range c.data {
			ch <- s
		}
	}()
	return ch
}
