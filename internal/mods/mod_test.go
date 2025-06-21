package mods

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestModStructFields(t *testing.T) {
	m := Mod{
		HashKey:     "hash123",
		Name:        "Test Mod",
		ModId:       "mod/123",
		SortedKey:   "Test Mod",
		Dependencies: []string{"dep1", "dep2"},
	}

	if m.HashKey != "hash123" {
		t.Errorf("Expected HashKey 'hash123', got '%s'", m.HashKey)
	}
	if m.Name != "Test Mod" {
		t.Errorf("Expected Name 'Test Mod', got '%s'", m.Name)
	}
	if m.ModId != "mod/123" {
		t.Errorf("Expected ModId 'mod/123', got '%s'", m.ModId)
	}
	if m.SortedKey != "Test Mod" {
		t.Errorf("Expected SortedKey 'Test Mod', got '%s'", m.SortedKey)
	}
	if len(m.Dependencies) != 2 || m.Dependencies[0] != "dep1" || m.Dependencies[1] != "dep2" {
		t.Errorf("Dependencies not set or incorrect: %v", m.Dependencies)
	}
}

func TestModZeroValue(t *testing.T) {
	var m Mod
	if m.HashKey != "" || m.Name != "" || m.ModId != "" || m.SortedKey != "" || m.Dependencies != nil {
		t.Errorf("Zero value Mod fields not as expected: %+v", m)
	}
}

func TestModJSONMarshaling(t *testing.T) {
	m := Mod{
		HashKey:     "h",
		Name:        "N",
		ModId:       "id",
		SortedKey:   "S",
		Dependencies: []string{"a", "b"},
	}
	b, err := json.Marshal(m)
	if err != nil {
		t.Fatalf("Failed to marshal Mod: %v", err)
	}
	var m2 Mod
	if err := json.Unmarshal(b, &m2); err != nil {
		t.Fatalf("Failed to unmarshal Mod: %v", err)
	}
	if !reflect.DeepEqual(m, m2) {
		t.Errorf("Expected %v, got %v after marshal/unmarshal", m, m2)
	}
}

func TestModFieldMutability(t *testing.T) {
	m := Mod{}
	m.HashKey = "abc"
	m.Name = "def"
	m.ModId = "ghi"
	m.SortedKey = "jkl"
	m.Dependencies = []string{"x"}
	if m.HashKey != "abc" || m.Name != "def" || m.ModId != "ghi" || m.SortedKey != "jkl" || len(m.Dependencies) != 1 || m.Dependencies[0] != "x" {
		t.Errorf("Mod fields not mutable as expected: %+v", m)
	}
}
