package mods

import (
	"testing"
)

func TestTweakModOrder(t *testing.T) {
	mods := []*Mod{
		{SortedKey: "abc"},
		{SortedKey: "ab"},
		{SortedKey: "a"},
	}
	result := TweakModOrder(mods)
	if len(result) != 3 {
		t.Errorf("Expected 3 mods, got %d", len(result))
	}
	if result[0].SortedKey != "a" {
		t.Errorf("Expected first mod to be 'a', got '%s'", result[0].SortedKey)
	}
}

func TestTweakModOrder_EmptyAndNil(t *testing.T) {
	if got := TweakModOrder([]*Mod{}); len(got) != 0 {
		t.Error("Expected empty slice for empty input")
	}
	defer func() {
		recover() // os.Exit will exit, but just in case
	}()
	// Should exit for nil input, but we can't test os.Exit easily
}

func TestGetHashFromName(t *testing.T) {
	data := map[string]map[string]interface{}{
		"h1": {"displayName": "foo"},
		"h2": {"displayName": "bar"},
	}
	h, found := GetHashFromName(data, "bar")
	if !found || h != "h2" {
		t.Errorf("Expected to find hash 'h2' for 'bar', got '%s'", h)
	}
	_, found = GetHashFromName(data, "baz")
	if found {
		t.Error("Did not expect to find hash for 'baz'")
	}
}

func TestGetHashFromName_DuplicateDisplayNames(t *testing.T) {
	data := map[string]map[string]interface{}{
		"h1": {"displayName": "foo"},
		"h2": {"displayName": "foo"},
	}
	h, found := GetHashFromName(data, "foo")
	if !found || (h != "h1" && h != "h2") {
		t.Errorf("Expected to find one of the hashes for duplicate displayName, got '%s'", h)
	}
}

func TestGetIndexFromHash(t *testing.T) {
	mods := []*Mod{{HashKey: "a"}, {HashKey: "b"}}
	idx := GetIndexFromHash(mods, "b", "test")
	if idx != 1 {
		t.Errorf("Expected index 1, got %d", idx)
	}
	idx = GetIndexFromHash(mods, "c", "test")
	if idx != -1 {
		t.Errorf("Expected -1 for missing hash, got %d", idx)
	}
}

func TestGetIndexFromHash_EmptyList(t *testing.T) {
	idx := GetIndexFromHash([]*Mod{}, "a", "test")
	if idx != -1 {
		t.Errorf("Expected -1 for empty mod list, got %d", idx)
	}
}

func TestSortAfterTags(t *testing.T) {
	allTags := map[string][]string{
		"OST": {"a"},
		"AI":  {"b"},
	}
	mods := []*Mod{{SortedKey: "a"}, {SortedKey: "b"}, {SortedKey: "c"}}
	result := SortAfterTags(allTags, mods)
	if len(result) != 3 {
		t.Errorf("Expected 3 mods, got %d", len(result))
	}
}

func TestSortAfterTags_OrderPreserved(t *testing.T) {
	allTags := map[string][]string{
		"OST": {"a", "b"},
	}
	mods := []*Mod{{SortedKey: "a"}, {SortedKey: "b"}, {SortedKey: "c"}}
	result := SortAfterTags(allTags, mods)
	if result[0].SortedKey != "a" || result[1].SortedKey != "b" {
		t.Errorf("Expected order a, b, ... got %v", []string{result[0].SortedKey, result[1].SortedKey})
	}
}

func TestSortAfterDependencies(t *testing.T) {
	mods := []*Mod{{HashKey: "a", ModId: "1"}, {HashKey: "b", ModId: "2"}}
	data := map[string]map[string]interface{}{
		"a": {"displayName": "A"},
		"b": {"displayName": "B"},
	}
	idList := []string{"1", "2"}
	result := SortAfterDependencies(mods, []string{"B"}, 0, "A", idList, data)
	if len(result) != 2 {
		t.Errorf("Expected 2 mods, got %d", len(result))
	}
}

func TestSortAfterDependencies_NoDependencies(t *testing.T) {
	mods := []*Mod{{HashKey: "a", ModId: "1"}, {HashKey: "b", ModId: "2"}}
	data := map[string]map[string]interface{}{
		"a": {"displayName": "A"},
		"b": {"displayName": "B"},
	}
	idList := []string{"1", "2"}
	result := SortAfterDependencies(mods, []string{}, 0, "A", idList, data)
	if len(result) != 2 {
		t.Errorf("Expected 2 mods, got %d", len(result))
	}
}

func TestSortDependencies(t *testing.T) {
	mods := []*Mod{{HashKey: "a", ModId: "1", Dependencies: []string{"B"}}, {HashKey: "b", ModId: "2"}}
	data := map[string]map[string]interface{}{
		"a": {"displayName": "A"},
		"b": {"displayName": "B"},
	}
	idList := []string{"1", "2"}
	result := SortDependencies(mods, idList, data)
	if len(result) != 2 {
		t.Errorf("Expected 2 mods, got %d", len(result))
	}
}

func TestSortDependencies_Empty(t *testing.T) {
	mods := []*Mod{}
	data := map[string]map[string]interface{}{}
	idList := []string{}
	result := SortDependencies(mods, idList, data)
	if len(result) != 0 {
		t.Errorf("Expected 0 mods, got %d", len(result))
	}
}

func TestSpecialOrder(t *testing.T) {
	mods := []*Mod{
		{Name: "UI Overhaul Dynamic", SortedKey: "UI Overhaul Dynamic"},
		{Name: "Dark UI", SortedKey: "Dark UI"},
		{Name: "Other", SortedKey: "Other"},
	}
	result := SpecialOrder(mods)
	if len(result) != 3 {
		t.Errorf("Expected 3 mods, got %d", len(result))
	}
}

func TestSpecialOrder_NoSpecials(t *testing.T) {
	mods := []*Mod{
		{Name: "Other", SortedKey: "Other"},
		{Name: "Another", SortedKey: "Another"},
	}
	result := SpecialOrder(mods)
	if len(result) != 2 {
		t.Errorf("Expected 2 mods, got %d", len(result))
	}
}

func TestContainsSpecial(t *testing.T) {
	if !containsSpecial("foobar", "foo") {
		t.Error("Expected containsSpecial to find substring")
	}
	if containsSpecial("bar", "foo") {
		t.Error("Did not expect containsSpecial to find substring")
	}
}

func TestGetModList(t *testing.T) {
	data := map[string]map[string]interface{}{
		"h1": {"displayName": "foo", "gameRegistryId": "id1"},
		"h2": {"displayName": "bar", "steamId": "id2"},
		"h3": {"displayName": "", "gameRegistryId": "id3"},
	}
	mods := GetModList(data)
	if len(mods) != 2 {
		t.Errorf("Expected 2 mods, got %d", len(mods))
	}
}

func TestGetModList_EmptyInput(t *testing.T) {
	mods := GetModList(map[string]map[string]interface{}{})
	if len(mods) != 0 {
		t.Errorf("Expected 0 mods, got %d", len(mods))
	}
}
