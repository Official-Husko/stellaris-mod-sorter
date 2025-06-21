package config

import (
	"os"
	"path/filepath"
)

type Config struct {
	SettingsPath string
	ModsRegistry string
	BakExt       string
}

// FindStellarisPath tries to locate the Stellaris settings directory.
func FindStellarisPath(modsRegistry string) (string, error) {
	candidates := []string{
		".",
		"..",
		filepath.Join(os.Getenv("HOME"), "Documents", "Paradox Interactive", "Stellaris"),
		filepath.Join(os.Getenv("HOME"), ".local", "share", "Paradox Interactive", "Stellaris"),
	}
	for _, s := range candidates {
		if _, err := os.Stat(filepath.Join(s, modsRegistry)); err == nil {
			return s, nil
		}
	}
	return "", os.ErrNotExist
}
