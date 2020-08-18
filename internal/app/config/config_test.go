package config_test

import (
	"testing"

	"github.com/cyruzin/bexs_challenge/internal/app/config"
)

func TestLoadConfigSuccess(t *testing.T) {
	config := config.Load()

	if config.EnvMode == "" {
		t.Error("empty env mode")
	}
}

func TestLoadConfigFail(t *testing.T) {
	config := config.Load()

	config.EnvMode = ""

	if config.EnvMode != "" {
		t.Error("not empty env mode")
	}
}
