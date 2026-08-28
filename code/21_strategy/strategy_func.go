package _21_strategy

import "sort"

// SortStrategy 排序策略，用函数类型代替接口，调用方直接传函数
type SortStrategy func([]int)

// Asc 升序排序
func Asc(data []int) {
	sort.Ints(data)
}

// Desc 降序排序
func Desc(data []int) {
	sort.Sort(sort.Reverse(sort.IntSlice(data)))
}

// Sorter 持有排序策略函数
type Sorter struct {
	strategy SortStrategy
}

func NewSorter(strategy SortStrategy) *Sorter {
	return &Sorter{strategy: strategy}
}

func (s *Sorter) Sort(data []int) {
	s.strategy(data)
}
