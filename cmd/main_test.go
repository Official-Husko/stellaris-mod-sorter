package main

import (
	"os"
	"testing"
)

// TestMainIntegration runs the main function and checks for successful completion.
func TestMainIntegration(t *testing.T) {
	// Setup: create a temp directory structure and files to simulate Stellaris environment
	tempDir := t.TempDir()
	os.Setenv("HOME", tempDir) // For config.FindStellarisPath
	os.Chdir(tempDir)

	// Create dummy mods_registry.json
	modsRegistryPath := tempDir + "/mods_registry.json"
	modsRegistryContent := `{"mod1": {"dirPath": "` + tempDir + `", "archivePath": "", "modId": "mod1", "HashKey": "h1", "SortedKey": "mod1"}}`
	if err := os.WriteFile(modsRegistryPath, []byte(modsRegistryContent), 0644); err != nil {
		t.Fatalf("failed to write mods_registry.json: %v", err)
	}

	// Create dummy dlc_load.json and game_data.json
	dlcLoadPath := tempDir + "/dlc_load.json"
	gameDataPath := tempDir + "/game_data.json"
	dlcLoadContent := `{"enabled_mods": ["mod1"]}`
	gameDataContent := `{"modsOrder": ["h1"]}`
	os.WriteFile(dlcLoadPath, []byte(dlcLoadContent), 0644)
	os.WriteFile(gameDataPath, []byte(gameDataContent), 0644)

	// Create dummy descriptor.mod in tempDir
	descPath := tempDir + "/descriptor.mod"
	descContent := `tags={
	"UI"
}
dependencies={
	"modA"
}`
	os.WriteFile(descPath, []byte(descContent), 0644)

	// Redirect os.Args and run main
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()
	os.Args = []string{"cmd"}

	// Run main and check for panic (should not panic)
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("main panicked: %v", r)
		}
	}()
	main()
}
