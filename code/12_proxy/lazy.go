package _12_proxy

import (
	"fmt"
	"sync"
)

// Image 是昂贵的真实对象，加载成本高。
type Image struct {
	filename string
}

func loadImage(filename string) *Image {
	// 模拟从磁盘/网络加载大对象。
	return &Image{filename: filename}
}

func (i *Image) draw() string {
	return fmt.Sprintf("显示图片 %s", i.filename)
}

// LazyImage 懒加载代理：只有第一次调用才真正创建 Image。
type LazyImage struct {
	filename string
	once     sync.Once
	real     *Image
}

func NewLazyImage(filename string) *LazyImage {
	return &LazyImage{filename: filename}
}

func (l *LazyImage) draw() string {
	l.once.Do(func() {
		l.real = loadImage(l.filename)
	})
	return l.real.draw()
}
