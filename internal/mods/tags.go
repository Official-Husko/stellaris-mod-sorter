package mods

import (
	"io/ioutil"
	"path/filepath"
	"strings"
)

// parseDependencies extracts dependencies from descriptor content.
func parseDependencies(desc string) []string {
	var deps []string
	inBlock := false
	for _, line := range splitLines(desc) {
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "dependencies") {
			inBlock = true
			continue
		}
		if inBlock {
			if strings.Contains(line, "}") {
				break
			}
			if line != "" {
				deps = append(deps, trimQuotes(line))
			}
		}
	}
	return deps
}

// parseTags extracts tags from descriptor content.
func parseTags(desc string) []string {
	var tags []string
	inBlock := false
	for _, line := range splitLines(desc) {
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "tags") {
			inBlock = true
			continue
		}
		if inBlock {
			if strings.Contains(line, "}") {
				break
			}
			if line != "" {
				tags = append(tags, trimQuotes(line))
			}
		}
	}
	return tags
}

// CheckDependencies parses dependencies from descriptor content and updates the mod.
func CheckDependencies(descContent []string, mod *Mod) {
	for _, desc := range descContent {
		if strings.Contains(desc, "dependencies") {
			deps := parseDependencies(desc)
			if len(deps) > 0 {
				mod.Dependencies = deps
			}
		}
	}
}

// CheckTags parses tags from descriptor content and updates allTags.
func CheckTags(descContent []string, mod *Mod, allTags map[string][]string) {
	for _, desc := range descContent {
		if strings.Contains(desc, "tags") {
			tags := parseTags(desc)
			for _, t := range tags {
				if !contains(allTags[t], mod.SortedKey) {
					allTags[t] = append(allTags[t], mod.SortedKey)
				}
			}
		}
	}
}

// GetModDescription processes mods, extracting tags and dependencies from descriptor files.
func GetModDescription(modList []*Mod, data map[string]map[string]interface{}, allTags map[string][]string, settingPath string) {
	for _, mod := range modList {
		d := data[mod.HashKey]
		dirPath, _ := d["dirPath"].(string)
		archivePath, _ := d["archivePath"].(string)
		descriptor := []string{}
		if dirPath == "" || !isDir(dirPath) {
			continue
		}
		descFile := filepath.Join(dirPath, "descriptor.mod")
		if fileExists(descFile) {
			if content, err := ioutil.ReadFile(descFile); err == nil {
				descriptor = append(descriptor, string(content))
			}
		}
		if archivePath != "" && len(descriptor) == 0 && fileExists(archivePath) {
			if err := extractZip(archivePath, dirPath); err == nil {
				if fileExists(descFile) {
					if content, err := ioutil.ReadFile(descFile); err == nil {
						descriptor = append(descriptor, string(content))
					}
				}
			}
		}
		if len(descriptor) == 0 {
			continue
		}
		modFile := filepath.Join(settingPath, "mod", mod.ModId)
		if fileExists(modFile) {
			if content, err := ioutil.ReadFile(modFile); err == nil {
				descriptor = append(descriptor, string(content))
			}
		}
		if len(descriptor) > 0 {
			CheckTags(descriptor, mod, allTags)
			CheckDependencies(descriptor, mod)
		}
	}
}
