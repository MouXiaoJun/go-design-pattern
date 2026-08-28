package _19_observer

// Stock 主题：维护价格状态与订阅者列表，价格变化时通知订阅者。
type Stock struct {
	name  string
	price float64
	subs  map[string]chan float64
}

func NewStock(name string) *Stock {
	return &Stock{name: name, subs: make(map[string]chan float64)}
}

func (s *Stock) Name() string   { return s.name }
func (s *Stock) Price() float64 { return s.price }

// Subscribe 返回一个只读通道，订阅者从通道消费价格更新。
// 通道带一个缓冲位，保证订阅者暂时来不及消费时广播仍能继续。
func (s *Stock) Subscribe(id string) <-chan float64 {
	ch := make(chan float64, 1)
	s.subs[id] = ch
	return ch
}

// Unsubscribe 取消订阅并关闭通道，通知订阅者停止消费。
func (s *Stock) Unsubscribe(id string) {
	if ch, ok := s.subs[id]; ok {
		close(ch)
		delete(s.subs, id)
	}
}

// Update 更新价格并广播给所有订阅者。
// 采用非阻塞发送：订阅者消费不及时时丢弃本次更新，避免拖慢甚至阻塞广播。
func (s *Stock) Update(price float64) {
	s.price = price
	for _, ch := range s.subs {
		select {
		case ch <- price:
		default:
		}
	}
}
