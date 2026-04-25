package initsystem

import (
	"context"
	"os"
	"path/filepath"
	"testing"
)

func TestSystemd_IsActive(t *testing.T) {
	sys := NewSystemd()
	if sys.Name() != "systemd" {
		t.Errorf("expected name 'systemd', got %s", sys.Name())
	}

	tempDir := t.TempDir()
	fakeSystemdDir := filepath.Join(tempDir, "system")
	if err := os.MkdirAll(fakeSystemdDir, 0o755); err != nil {
		t.Fatalf("failed to create temp dir: %v", err)
	}

	oldSystemdDir := systemdDir
	defer func() { systemdDir = oldSystemdDir }()

	systemdDir = fakeSystemdDir
	if !sys.IsActive(context.Background()) {
		t.Error("expected systemd to be active with mock directory")
	}

	systemdDir = filepath.Join(tempDir, "nonexistent")
	if sys.IsActive(context.Background()) {
		t.Error("expected systemd to be inactive with nonexistent directory")
	}
}
