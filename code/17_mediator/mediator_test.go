package _17_mediator

import "testing"

func TestChatRoomBroadcast(t *testing.T) {
	room := NewChatRoom()
	alice := NewUser("alice", room)
	bob := NewUser("bob", room)
	carol := NewUser("carol", room)

	alice.Broadcast("hello")

	if len(alice.Messages()) != 0 {
		t.Fatalf("发送者不应收到自己的消息: %v", alice.Messages())
	}
	if len(bob.Messages()) != 1 || len(carol.Messages()) != 1 {
		t.Fatalf("其他用户应各收到一条消息: bob=%v carol=%v", bob.Messages(), carol.Messages())
	}
}

func TestChatRoomSend(t *testing.T) {
	room := NewChatRoom()
	alice := NewUser("alice", room)
	bob := NewUser("bob", room)

	alice.Send("bob", "hi bob")

	if len(bob.Messages()) != 1 {
		t.Fatalf("bob 应收到一条消息: %v", bob.Messages())
	}
	if len(alice.Messages()) != 0 {
		t.Fatalf("alice 不应收到消息: %v", alice.Messages())
	}
}
