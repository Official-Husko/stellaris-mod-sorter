package mods

import (
	"bytes"
	"encoding/json"
	"os"
	"path/filepath"
	"reflect"
	"testing"
)

func TestLoadJsonOrder_WriteJsonOrder_RoundTrip(t *testing.T) {
	dir := t.TempDir()
	file := filepath.Join(dir, "test.json")
	bak := ".bak"
	data := map[string]interface{}{"foo": "bar", "num": 42.0}
	content, _ := json.MarshalIndent(data, "", "  ")
	os.WriteFile(file, content, 0644)

	loaded, path := LoadJsonOrder(dir, "test.json", bak)
	if !reflect.DeepEqual(loaded, data) {
		t.Errorf("Loaded data does not match: got %v, want %v", loaded, data)
	}
	if path != file {
		t.Errorf("Returned file path incorrect: got %s, want %s", path, file)
	}

	// Test WriteJsonOrder (should backup and overwrite)
	data2 := map[string]interface{}{"foo": "baz"}
	WriteJsonOrder(data2, file, bak)
	if _, err := os.Stat(file+bak); err != nil {
		t.Error("Expected backup file to exist after WriteJsonOrder")
	}
	loaded2, _ := LoadJsonOrder(dir, "test.json", bak)
	if loaded2["foo"] != "baz" {
		t.Errorf("Expected 'foo' to be 'baz', got %v", loaded2["foo"])
	}
}

func TestLoadJsonOrder_FileNotExist(t *testing.T) {
	dir := t.TempDir()
	loaded, path := LoadJsonOrder(dir, "nope.json", ".bak")
	if len(loaded) != 0 {
		t.Error("Expected empty map for missing file")
	}
	if path != filepath.Join(dir, "nope.json") {
		t.Errorf("Expected path to be %s, got %s", filepath.Join(dir, "nope.json"), path)
	}
}

func TestWriteJsonOrder_Backup(t *testing.T) {
	dir := t.TempDir()
	file := filepath.Join(dir, "test.json")
	bak := ".bak"
	os.WriteFile(file, []byte(`{"foo": "bar"}`), 0644)
	WriteJsonOrder(map[string]interface{}{"foo": "baz"}, file, bak)
	if _, err := os.Stat(file + bak); err != nil {
		t.Error("Expected backup file to exist")
	}
}

func TestDecodeJSON(t *testing.T) {
	data := map[string]interface{}{"a": 1.0, "b": "two"}
	buf := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(data)
	var out map[string]interface{}
	if err := DecodeJSON(buf, &out); err != nil {
		t.Fatalf("DecodeJSON failed: %v", err)
	}
	if !reflect.DeepEqual(data, out) {
		t.Errorf("Decoded data mismatch: got %v, want %v", out, data)
	}
}

func TestDecodeJSON_Invalid(t *testing.T) {
	buf := bytes.NewBufferString("not json")
	var out map[string]interface{}
	if err := DecodeJSON(buf, &out); err == nil {
		t.Error("Expected error for invalid JSON")
	}
}

func TestValidateExampleRegistry(t *testing.T) {
	schemaPath := filepath.Join("..", "..", "mods_registry.schema.json")
	jsonPath := filepath.Join("..", "..", "example_registry.json")
	err := ValidateJSONSchema(jsonPath, schemaPath)
	if err != nil {
		t.Errorf("example_registry.json is NOT valid: %v", err)
	} else {
		t.Log("example_registry.json is valid against mods_registry.schema.json")
	}
}
