package config

import (
	"os"
	"path/filepath"
	"testing"
)

func TestFindStellarisPath_FindsExisting(t *testing.T) {
	dir := t.TempDir()
	modsRegistry := "test_registry.mod"
	filePath := filepath.Join(dir, modsRegistry)
	if err := os.WriteFile(filePath, []byte("dummy"), 0644); err != nil {
		t.Fatalf("failed to create test registry: %v", err)
	}
	cwd, _ := os.Getwd()
	defer os.Chdir(cwd)
	os.Chdir(dir)
	found, err := FindStellarisPath(modsRegistry)
	if err != nil {
		t.Errorf("expected to find registry, got error: %v", err)
	}
	if found != "." && found != dir {
		t.Errorf("expected '.' or temp dir, got %q", found)
	}
}

func TestFindStellarisPath_NotFound(t *testing.T) {
	modsRegistry := "nonexistent_registry.mod"
	_, err := FindStellarisPath(modsRegistry)
	if err == nil {
		t.Error("expected error for missing registry, got nil")
	}
}

func TestConfigStructFields(t *testing.T) {
	cfg := Config{
		SettingsPath: "/tmp/settings",
		ModsRegistry: "foo.mod",
		BakExt:       ".bak",
	}
	if cfg.SettingsPath != "/tmp/settings" || cfg.ModsRegistry != "foo.mod" || cfg.BakExt != ".bak" {
		t.Errorf("Config struct fields not set correctly: %+v", cfg)
	}
}