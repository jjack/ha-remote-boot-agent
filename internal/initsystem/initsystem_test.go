package initsystem

import (
	"context"
	"testing"
)

type mockInitSystem struct {
	name   string
	active bool
}

func (m *mockInitSystem) Name() string                      { return m.name }
func (m *mockInitSystem) IsActive(ctx context.Context) bool { return m.active }

func TestRegistry(t *testing.T) {
	reg := NewRegistry()
	reg.Register("mock", func() InitSystem { return &mockInitSystem{name: "mock", active: true} })

	sys := reg.Get("mock")
	if sys == nil || sys.Name() != "mock" {
		t.Fatal("failed to get mock init system")
	}

	none := reg.Get("nonexistent")
	if none != nil {
		t.Fatal("expected nil for nonexistent init system")
	}

	detected, err := reg.Detect(context.Background())
	if err != nil || detected == nil || detected.Name() != "mock" {
		t.Fatalf("failed to detect mock init system: %v", err)
	}
}

func TestRegistry_DetectFail(t *testing.T) {
	reg := NewRegistry()
	reg.Register("mock", func() InitSystem { return &mockInitSystem{name: "mock", active: false} })

	_, err := reg.Detect(context.Background())
	if err == nil {
		t.Fatal("expected error when no active init system is detected")
	}
}
