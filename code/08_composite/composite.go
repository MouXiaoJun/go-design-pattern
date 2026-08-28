package _8_composite

import "strings"

// Node 文件系统统一抽象：叶子与容器对调用方透明
type Node interface {
	Name() string
	Size() int
	Print(prefix string) string
}

// File 叶子节点
type File struct {
	name string
	size int
}

func NewFile(name string, size int) *File {
	return &File{name: name, size: size}
}

func (f *File) Name() string { return f.name }

func (f *File) Size() int { return f.size }

func (f *File) Print(prefix string) string {
	return prefix + f.name
}

// Directory 容器节点：递归管理子节点
type Directory struct {
	name     string
	children []Node
}

func NewDirectory(name string) *Directory {
	return &Directory{name: name}
}

func (d *Directory) Add(child Node) {
	d.children = append(d.children, child)
}

func (d *Directory) Name() string { return d.name }

func (d *Directory) Size() int {
	total := 0
	for _, child := range d.children {
		total += child.Size()
	}
	return total
}

func (d *Directory) Print(prefix string) string {
	var b strings.Builder
	b.WriteString(prefix + d.name + "/")
	for _, child := range d.children {
		b.WriteString("\n" + child.Print(prefix+"  "))
	}
	return b.String()
}
