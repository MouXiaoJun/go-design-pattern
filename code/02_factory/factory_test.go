package _2_factory

import (
	"testing"
)

// ---- 简单工厂 ----

func TestNewConfigParser(t *testing.T) {
	cases := map[string]bool{
		"json":  true,
		"toml":  true,
		"yaml":  true,
		"xml":   false,
		"":      false,
	}
	for kind, want := range cases {
		if got := NewConfigParser(kind); (got != nil) != want {
			t.Fatalf("NewConfigParser(%q) != nil = %v, want %v", kind, got != nil, want)
		}
	}
}

// ---- 工厂方法 ----

func TestNewJsonConfigFactory(t *testing.T) {
	jsonFactory := NewJsonConfigFactory("json")
	if jsonFactory == nil {
		t.Fatalf("json factory should not be nil")
	}
	if _, ok := jsonFactory.createParser().(*jsonConfigParser); !ok {
		t.Fatalf("json factory should create jsonConfigParser")
	}

	yamlFactory := NewJsonConfigFactory("yaml")
	if _, ok := yamlFactory.createParser().(*yamlConfigParser); !ok {
		t.Fatalf("yaml factory should create yamlConfigParser")
	}

	if got := NewJsonConfigFactory("xml"); got != nil {
		t.Fatalf("unknown kind should return nil factory")
	}
}

// ---- DI 容器 ----

type serviceA struct{ Value string }
type serviceB struct{ Value string }

type consumer struct {
	A *serviceA `inject:"a"`
	B *serviceB `inject:"b"`
}

func TestContainerInject(t *testing.T) {
	c := NewContainer()
	c.RegisterBean("a", &serviceA{Value: "A"})
	c.RegisterBean("b", &serviceB{Value: "B"})

	con := &consumer{}
	if err := c.InjectBean(con); err != nil {
		t.Fatalf("InjectBean error: %v", err)
	}
	if con.A == nil || con.A.Value != "A" {
		t.Fatalf("A not injected: %+v", con.A)
	}
	if con.B == nil || con.B.Value != "B" {
		t.Fatalf("B not injected: %+v", con.B)
	}
}

func TestContainerErrors(t *testing.T) {
	c := NewContainer()

	// 非指针
	if err := c.InjectBean(consumer{}); err == nil {
		t.Fatalf("injecting non-pointer should error")
	}

	// 缺少依赖
	type missing struct {
		A *serviceA `inject:"a"`
	}
	if err := c.InjectBean(&missing{}); err == nil {
		t.Fatalf("missing bean should error")
	}

	// 类型不匹配
	c.RegisterBean("x", "not-a-service")
	type mismatch struct {
		A *serviceA `inject:"x"`
	}
	if err := c.InjectBean(&mismatch{}); err == nil {
		t.Fatalf("type mismatch should error")
	}
}

func TestContainerGetBean(t *testing.T) {
	c := NewContainer()
	if _, err := c.GetBean("nope"); err == nil {
		t.Fatalf("GetBean(missing) should error")
	}
	c.RegisterBean("ok", 42)
	v, err := c.GetBean("ok")
	if err != nil || v != 42 {
		t.Fatalf("GetBean(ok) = %v, %v; want 42, nil", v, err)
	}
}
