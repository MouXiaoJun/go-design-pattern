package _17_mediator

import "fmt"

// Mediator 中介者接口：统一协调 Colleague 之间的交互，让对象之间不直接引用。
type Mediator interface {
	Register(user *User)
	Send(from *User, to string, msg string)
	Broadcast(from *User, msg string)
}

// ChatRoom 是 Mediator 的一个具体实现：聊天室。
type ChatRoom struct {
	users map[string]*User
}

func NewChatRoom() *ChatRoom {
	return &ChatRoom{users: make(map[string]*User)}
}

func (r *ChatRoom) Register(user *User) {
	r.users[user.name] = user
}

// Send 点对点转发：发送方只依赖中介者，不依赖具体接收者。
func (r *ChatRoom) Send(from *User, to string, msg string) {
	if target, ok := r.users[to]; ok {
		target.receive(from, msg)
	}
}

// Broadcast 广播给除发送者之外的所有用户。
func (r *ChatRoom) Broadcast(from *User, msg string) {
	for name, user := range r.users {
		if name != from.name {
			user.receive(from, msg)
		}
	}
}

// User 是同事对象：只持有对中介者的引用，不直接引用其他用户。
type User struct {
	name     string
	room     Mediator
	received []string
}

func NewUser(name string, room Mediator) *User {
	u := &User{name: name, room: room}
	room.Register(u)
	return u
}

func (u *User) Name() string { return u.name }

// Send 点对点发送消息。
func (u *User) Send(to string, msg string) {
	u.room.Send(u, to, msg)
}

// Broadcast 向聊天室广播消息。
func (u *User) Broadcast(msg string) {
	u.room.Broadcast(u, msg)
}

func (u *User) receive(from *User, msg string) {
	u.received = append(u.received, fmt.Sprintf("%s -> %s: %s", from.name, u.name, msg))
}

// Messages 返回该用户收到的消息，便于测试与观察。
func (u *User) Messages() []string { return u.received }
