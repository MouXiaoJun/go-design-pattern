package _4_prototype

import (
	"testing"
	"time"
)

func TestCloneDeepCopy(t *testing.T) {
	now := time.Now()
	original := &Keyword{Word: "golang", Visit: 10, Updated: &now}

	clone := original.Clone()

	// 深拷贝：改克隆体的 Updated 不应影响原型
	clone.Updated = &[]time.Time{now.Add(time.Hour)}[0]
	clone.Visit = 99

	if original.Visit != 10 {
		t.Fatalf("deep copy failed: original.Visit = %d, want 10", original.Visit)
	}
	if original.Updated.Equal(*clone.Updated) {
		t.Fatalf("deep copy failed: Updated pointer shared")
	}
}

func TestShallowCopySharesPointer(t *testing.T) {
	now := time.Now()
	original := &Keyword{Word: "golang", Visit: 10, Updated: &now}

	shallow := original.ShallowCopy()
	if shallow.Updated != original.Updated {
		t.Fatalf("shallow copy should share Updated pointer")
	}
}

func TestCloneFromRegistry(t *testing.T) {
	Keywords = map[string]*Keyword{
		"hot": NewKeyword("热词", 100),
	}
	defer func() { Keywords = nil }()

	clone := CloneFrom("hot")
	if clone == nil {
		t.Fatalf("CloneFrom should return a clone")
	}
	if clone.Word != "热词" {
		t.Fatalf("CloneFrom word = %q, want 热词", clone.Word)
	}

	if got := CloneFrom("missing"); got != nil {
		t.Fatalf("CloneFrom(missing) should return nil")
	}
}
