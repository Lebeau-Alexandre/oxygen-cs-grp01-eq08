package tests

import (
	"os"
	"testing"

	"github.com/Lebeau-Alexandre/oxygen-cs-grp01-eq08/config"
)

func TestMain(m *testing.M) {
	config.InitConfig()

	exitCode := m.Run()
	defer func() {
		config.Clear()
	}()
	os.Exit(exitCode)
}
