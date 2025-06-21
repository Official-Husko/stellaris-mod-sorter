package mods

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"reflect"
	"testing"
)

func TestCheckTags(t *testing.T) {
	descContent := []string{
		`tags={
	"UI"
	"Overhaul"
}`,
	}
	allTags := make(map[string][]string)
	mod := &Mod{SortedKey: "mod1"}
	CheckTags(descContent, mod, allTags)
	if !reflect.DeepEqual(allTags["UI"], []string{"mod1"}) || !reflect.DeepEqual(allTags["Overhaul"], []string{"mod1"}) {
		t.Errorf("CheckTags failed: got %v", allTags)
	}
}

func TestCheckDependencies(t *testing.T) {
	descContent := []string{
		`dependencies={
	"modA"
	"modB"
}`,
	}
	mod := &Mod{ModId: "mod1"}
	CheckDependencies(descContent, mod)
	if !reflect.DeepEqual(mod.Dependencies, []string{"modA", "modB"}) {
		t.Errorf("CheckDependencies failed: got %v", mod.Dependencies)
	}
}

func TestGetModDescription(t *testing.T) {
	dir, err := ioutil.TempDir("", "modtest")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(dir)
	descPath := filepath.Join(dir, "descriptor.mod")
	desc := `tags={
	"UI"
}
dependencies={
	"modA"
}`
	ioutil.WriteFile(descPath, []byte(desc), 0644)
	modList := []*Mod{{ModId: "mod1", HashKey: "h1", SortedKey: "mod1"}}
	data := map[string]map[string]interface{}{
		"h1": {"dirPath": dir, "archivePath": ""},
	}
	allTags := make(map[string][]string)
	GetModDescription(modList, data, allTags, dir)
	if !reflect.DeepEqual(allTags["UI"], []string{"mod1"}) {
		t.Errorf("GetModDescription tags failed: got %v", allTags)
	}
	if !reflect.DeepEqual(modList[0].Dependencies, []string{"modA"}) {
		t.Errorf("GetModDescription dependencies failed: got %v", modList[0].Dependencies)
	}
}
