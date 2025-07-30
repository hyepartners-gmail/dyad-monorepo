package config

import "testing"

func TestLoadEnvFallback(t *testing.T) {
	t.Setenv("PORT", "9999")
	cfg := loadFromEnv()

	if cfg.Port != "9999" {
		t.Errorf("expected 9999, got %s", cfg.Port)
	}
}
