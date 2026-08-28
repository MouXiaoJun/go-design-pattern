package _4_prototype

import (
	"encoding/json"
	"time"
)

// Keywords 原型注册表：登记一批创建成本高的对象，按需复制。
var Keywords map[string]*Keyword

// Keyword 模拟一个创建成本高的对象（真实场景可能来自慢 IO / 重计算）。
type Keyword struct {
	Word    string
	Visit   int
	Updated *time.Time
}

// NewKeyword 创建一个原型对象。注意：字段是导出的，深拷贝才能通过 JSON 复制到。
func NewKeyword(word string, visit int) *Keyword {
	now := time.Now()
	return &Keyword{Word: word, Visit: visit, Updated: &now}
}

// Clone 深拷贝：JSON 序列化实现。只能拷贝导出字段，未导出字段会被忽略。
func (k *Keyword) Clone() *Keyword {
	var newKeyword Keyword
	b, _ := json.Marshal(k)
	json.Unmarshal(b, &newKeyword)
	return &newKeyword
}

// ShallowCopy 浅拷贝：复制值本身，但嵌套指针（Updated）仍与原对象共享。
func (k *Keyword) ShallowCopy() *Keyword {
	c := *k
	return &c
}

// CloneFrom 从注册表复制原型（原型模式的标准用法）。
func CloneFrom(name string) *Keyword {
	p, ok := Keywords[name]
	if !ok {
		return nil
	}
	return p.Clone()
}
