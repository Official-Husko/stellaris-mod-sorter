package mods

import (
	"fmt"
	"sort"
	"strings"

	prettylog "stellaris-mod-sorter-go/internal/utils"
)

// tweakModOrder swaps mods if one's SortedKey is a prefix of the next's SortedKey.
func TweakModOrder(arr []*Mod) []*Mod {
	for i := len(arr) - 1; i > 0; i-- {
		j := i - 1
		if len(arr[j].SortedKey) > 0 && len(arr[i].SortedKey) > 0 &&
			len(arr[j].SortedKey) >= len(arr[i].SortedKey) &&
			arr[j].SortedKey[:len(arr[i].SortedKey)] == arr[i].SortedKey {
			arr[j], arr[i] = arr[i], arr[j]
		}
	}
	if len(arr) > 0 {
		return arr
	}
	prettylog.PrintPretty("TweakModOrder", "no mod found", prettylog.LogFatal)
	return nil
}

// getHashFromName returns the hash key for a given mod name.
func GetHashFromName(data map[string]map[string]interface{}, name string) (string, bool) {
	for h, d := range data {
		if displayName, ok := d["displayName"].(string); ok && displayName == name {
			return h, true
		}
	}
	return "", false
}

// getIndexFromHash returns the index of the mod with the given hashKey using a map for O(1) lookup.
func GetIndexFromHash(modList []*Mod, h, name string) int {
	modMap := make(map[string]int, len(modList))
	for i, mod := range modList {
		modMap[mod.HashKey] = i
	}
	if idx, ok := modMap[h]; ok {
		return idx
	}
	prettylog.PrintPretty("GetIndexFromHash", fmt.Sprintf("Hashkey not found in game_data.json %s", h), prettylog.LogError)
	return -1
}

// sortAfterTags merges allTags with modList according to tag rules.
func SortAfterTags(allTags map[string][]string, modList []*Mod) []*Mod {
	output := []string{}
	addAfter := []string{}

	// Remove duplicates, keep last occurrence
	rmvDupes := func(dupes []string) []string {
		finalList := []string{}
		seen := map[string]bool{}
		for i := len(dupes) - 1; i >= 0; i-- {
			d := dupes[i]
			if !seen[d] {
				finalList = append([]string{d}, finalList...)
				seen[d] = true
			}
		}
		return finalList
	}

	reorderModList := func(name string) {
		for i, mod := range modList {
			if mod.SortedKey == name {
				modList = append(modList[:i], modList[i+1:]...)
				modList = append(modList, mod)
				break
			}
		}
	}

	insertPairToModList := func(name, name2 string) {
		var comp *Mod
		for i, mod := range modList {
			if mod.SortedKey == name2 {
				comp = mod
				modList = append(modList[:i], modList[i+1:]...)
				break
			}
		}
		if comp != nil {
			for i, mod := range modList {
				if mod.SortedKey == name {
					modList = append(modList[:i+1], append([]*Mod{comp}, modList[i+1:]...)...)
					break
				}
			}
		}
	}

	for _, o := range []string{"OST", "Music", "Sound", "Graphics"} {
		if mods, ok := allTags[o]; ok {
			output = append(output, mods...)
			delete(allTags, o)
		}
	}
	for _, o := range []string{"AI", "Utilities", "Fixes"} {
		if mods, ok := allTags[o]; ok {
			addAfter = append(addAfter, mods...)
			delete(allTags, o)
		}
	}
	if mods, ok := allTags["Patch"]; ok {
		for _, x := range mods {
			found := false
			for _, y := range addAfter {
				if x == y {
					found = true
					break
				}
			}
			if !found {
				addAfter = append(addAfter, x)
			}
		}
		delete(allTags, "Patch")
	}

	for _, mods := range allTags {
		if len(mods) == 1 {
			continue
		}
		if len(mods) == 2 {
			insertPairToModList(mods[0], mods[1])
			continue
		}
		output = append(output, mods...)
	}

	output = append(output, addAfter...)
	output = rmvDupes(output)

	for _, name := range output {
		reorderModList(name)
	}

	return modList
}

// sortAfterDependencies reorders modList based on dependencies for a given mod.
func SortAfterDependencies(modList []*Mod, dependencies []string, order int, name string, idList []string, data map[string]map[string]interface{}) []*Mod {
	for _, n := range dependencies {
		i, found := GetHashFromName(data, n)
		if !found {
			if contains(idList, modList[order].ModId) {
				prettylog.PrintPretty("SortAfterDependencies", fmt.Sprintf("Fail dependencie: %s not found for %s in mods_registry", n, name), prettylog.LogWarning)
			}
			continue
		}
		idx := GetIndexFromHash(modList, i, name)
		if idx > order {
			prettylog.PrintPretty("SortAfterDependencies", fmt.Sprintf("FIX dependencie: %s - %d is lower than %d - %s", name, order, idx, n), prettylog.LogInfo)
			item := modList[order]
			modList = append(modList[:order], modList[order+1:]...)
			// Insert before idx (which may have shifted)
			if idx > len(modList) {
				idx = len(modList)
			}
			modList = append(modList[:idx], append([]*Mod{item}, modList[idx:]...)...)
			order = idx
		}
	}
	return modList
}

// sortDependencies applies dependency sorting to all mods in modList.
func SortDependencies(modList []*Mod, idList []string, data map[string]map[string]interface{}) []*Mod {
	for idx, mod := range modList {
		if len(mod.Dependencies) > 0 {
			modList = SortAfterDependencies(modList, mod.Dependencies, idx, mod.SortedKey, idList, data)
		}
	}
	return modList
}

// specialOrder applies custom ordering for specific mods if no dependency is present.
func SpecialOrder(modList []*Mod) []*Mod {
	specialNames := []string{"UI Overhaul Dynamic", "Dark UI", "Dark U1"}
	specialList := make([]struct{ idx int; mod *Mod }, 0)
	for _, specialName := range specialNames {
		for i, mod := range modList {
			if containsSpecial(mod.Name, specialName) {
				specialList = append(specialList, struct{ idx int; mod *Mod }{i, mod})
			}
		}
	}
	if len(specialList) > 1 {
		c := specialList[0].idx
		specialList = specialList[1:]
		for _, cmpMod := range specialList {
			ix := cmpMod.idx
			if c > ix {
				for i, mod := range modList {
					if mod.Name == cmpMod.mod.Name {
						cmp := modList[i]
						modList = append(modList[:i], modList[i+1:]...)
						prettylog.PrintPretty("SpecialOrder", fmt.Sprintf("Special order %s after %s", cmp.SortedKey, modList[c-1].SortedKey), prettylog.LogInfo)
						if c > len(modList) {
							c = len(modList)
						}
						modList = append(modList[:c], append([]*Mod{cmp}, modList[c:]...)...)
						c++
						break
					}
				}
			} else {
				c = ix
			}
		}
	}
	return modList
}

// ContainsSpecial checks if a substring is in a string (case-sensitive) using strings.Contains for performance.
func containsSpecial(s, substr string) bool {
	return strings.Contains(s, substr)
}

// GetModList converts the registry data to a slice of *Mod, sorted by SortedKey descending.
func GetModList(data map[string]map[string]interface{}) []*Mod {
	modList := []*Mod{}
	keyToMod := make(map[string]*Mod, len(data))
	for key, d := range data {
		modId := ""
		if v, ok := d["gameRegistryId"].(string); ok && v != "" {
			modId = v
		} else if v, ok := d["steamId"].(string); ok && v != "" {
			modId = v
		}
		name, _ := d["displayName"].(string)
		if modId == "" || name == "" {
			continue
		}
		mod := &Mod{
			HashKey:   key,
			Name:      name,
			ModId:     modId,
			SortedKey: name,
		}
		modList = append(modList, mod)
		keyToMod[key] = mod
	}
	// Sort by SortedKey descending
	sort.Slice(modList, func(i, j int) bool {
		return modList[i].SortedKey > modList[j].SortedKey
	})
	return modList
}
