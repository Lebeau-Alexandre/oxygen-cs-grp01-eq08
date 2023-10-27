package tests

import (
	"os"
	"testing"

	"github.com/Lebeau-Alexandre/oxygen-cs-grp01-eq08/config"
	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func TestMain(m *testing.M) {
	config.InitConfig()
	gin.SetMode(gin.TestMode)

	exitCode := m.Run()
	defer func() {
		config.Clear()
	}()
	os.Exit(exitCode)
}
