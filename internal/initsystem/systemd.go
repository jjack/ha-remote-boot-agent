package initsystem

import (
	"context"
	"os"
)

const systemdName = "systemd"

// systemdDir is defined as a variable so it can be easily mocked in tests
var systemdDir = "/run/systemd/system"

type Systemd struct{}

func NewSystemd() InitSystem {
	return &Systemd{}
}

func (s *Systemd) Name() string {
	return systemdName
}

func (s *Systemd) IsActive(ctx context.Context) bool {
	fi, err := os.Stat(systemdDir)
	return err == nil && fi.IsDir()
}
