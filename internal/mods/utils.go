package mods

import (
	"archive/zip"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func contains(slice []string, s string) bool {
	for _, v := range slice {
		if v == s {
			return true
		}
	}
	return false
}

func containsStr(s, substr string) bool {
	return strings.Contains(s, substr)
}

func splitLines(s string) []string {
	return strings.Split(s, "\n")
}

func trimQuotes(s string) string {
	s = strings.TrimSpace(s)
	s = strings.Trim(s, "\"")
	return s
}

func isDir(path string) bool {
	info, err := os.Stat(path)
	return err == nil && info.IsDir()
}

func fileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

func extractZip(archivePath, dirPath string) error {
	r, err := zip.OpenReader(archivePath)
	if err != nil {
		return err
	}
	defer r.Close()
	for _, f := range r.File {
		fpath := filepath.Join(dirPath, f.Name)
		if f.FileInfo().IsDir() {
			os.MkdirAll(fpath, os.ModePerm)
			continue
		}
		if err := os.MkdirAll(filepath.Dir(fpath), os.ModePerm); err != nil {
			return err
		}
		outFile, err := os.OpenFile(fpath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
		if err != nil {
			return err
		}
		rc, err := f.Open()
		if err != nil {
			outFile.Close()
			return err
		}
		_, err = io.Copy(outFile, rc)
		outFile.Close()
		rc.Close()
		if err != nil {
			return err
		}
	}
	return nil
}

// getModHashKeys returns a slice of HashKeys from the mod list.
func GetModHashKeys(modList []*Mod) []string {
	result := make([]string, len(modList))
	for i, mod := range modList {
		result[i] = mod.HashKey
	}
	return result
}

// sliceToSet converts a slice of strings to a set for O(1) lookups.
func sliceToSet(slice []string) map[string]struct{} {
	set := make(map[string]struct{}, len(slice))
	for _, v := range slice {
		set[v] = struct{}{}
	}
	return set
}

// GetModIdsReversed returns a reversed slice of ModIds from modList that are present in idList.
func GetModIdsReversed(modList []*Mod, idList []string) []string {
	idSet := sliceToSet(idList)
	result := []string{}
	for i := len(modList) - 1; i >= 0; i-- {
		if _, ok := idSet[modList[i].ModId]; ok {
			result = append(result, modList[i].ModId)
		}
	}
	return result
}

// Contains checks if a string is present in a slice.
func Contains(slice []string, s string) bool {
	for _, v := range slice {
		if v == s {
			return true
		}
	}
	return false
}
