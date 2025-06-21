package mods

import (
	"archive/zip"
	"os"
	"path/filepath"
	"testing"
)

func TestGetModHashKeys(t *testing.T) {
	mods := []*Mod{
		{HashKey: "a"},
		{HashKey: "b"},
		{HashKey: "c"},
	}
	expected := []string{"a", "b", "c"}
	result := GetModHashKeys(mods)
	if len(result) != len(expected) {
		t.Fatalf("expected length %d, got %d", len(expected), len(result))
	}
	for i := range expected {
		if result[i] != expected[i] {
			t.Errorf("expected %q at %d, got %q", expected[i], i, result[i])
		}
	}
}

func TestGetModIdsReversed(t *testing.T) {
	mods := []*Mod{
		{ModId: "a"},
		{ModId: "b"},
		{ModId: "c"},
	}
	idList := []string{"a", "b", "c", "d"}
	expected := []string{"c", "b", "a"}
	result := GetModIdsReversed(mods, idList)
	if len(result) != len(expected) {
		t.Fatalf("expected length %d, got %d", len(expected), len(result))
	}
	for i := range expected {
		if result[i] != expected[i] {
			t.Errorf("expected %q at %d, got %q", expected[i], i, result[i])
		}
	}

	// Test with idList missing some mods
	idList = []string{"b", "c"}
	expected = []string{"c", "b"}
	result = GetModIdsReversed(mods, idList)
	if len(result) != len(expected) {
		t.Fatalf("expected length %d, got %d", len(expected), len(result))
	}
	for i := range expected {
		if result[i] != expected[i] {
			t.Errorf("expected %q at %d, got %q", expected[i], i, result[i])
		}
	}

	// Test with empty idList
	idList = []string{}
	expected = []string{}
	result = GetModIdsReversed(mods, idList)
	if len(result) != 0 {
		t.Errorf("expected empty result, got %v", result)
	}
}

func TestContains(t *testing.T) {
	slice := []string{"foo", "bar", "baz"}
	if !Contains(slice, "foo") {
		t.Error("expected Contains to find 'foo'")
	}
	if Contains(slice, "qux") {
		t.Error("expected Contains to not find 'qux'")
	}
	if Contains([]string{}, "foo") {
		t.Error("expected Contains to not find anything in empty slice")
	}
}

func TestContains_EmptySlice(t *testing.T) {
	if contains([]string{}, "a") {
		t.Error("Expected not to find anything in empty slice")
	}
}

func TestContainsStr(t *testing.T) {
	if !containsStr("foobar", "foo") {
		t.Error("Expected to find 'foo' in 'foobar'")
	}
	if containsStr("bar", "foo") {
		t.Error("Did not expect to find 'foo' in 'bar'")
	}
}

func TestContainsStr_Unicode(t *testing.T) {
	if !containsStr("héllo世界", "世界") {
		t.Error("Expected to find unicode substring")
	}
}

func TestSplitLines(t *testing.T) {
	lines := splitLines("a\nb\nc")
	if len(lines) != 3 || lines[0] != "a" || lines[2] != "c" {
		t.Error("splitLines did not split correctly")
	}
}

func TestSplitLines_Empty(t *testing.T) {
	lines := splitLines("")
	if len(lines) != 1 || lines[0] != "" {
		t.Error("splitLines should return one empty string for empty input")
	}
}

func TestTrimQuotes(t *testing.T) {
	if trimQuotes("  \"hello\"  ") != "hello" {
		t.Error("trimQuotes did not trim quotes and spaces correctly")
	}
}

func TestTrimQuotes_NoQuotes(t *testing.T) {
	if trimQuotes("hello") != "hello" {
		t.Error("trimQuotes should not alter string without quotes")
	}
}

func TestIsDirAndFileExists(t *testing.T) {
	dir := t.TempDir()
	if !isDir(dir) {
		t.Error("Expected temp dir to be a directory")
	}
	file := filepath.Join(dir, "testfile.txt")
	os.WriteFile(file, []byte("data"), 0644)
	if !fileExists(file) {
		t.Error("Expected file to exist")
	}
}

func TestIsDir_FileInsteadOfDir(t *testing.T) {
	dir := t.TempDir()
	file := filepath.Join(dir, "file.txt")
	os.WriteFile(file, []byte("x"), 0644)
	if isDir(file) {
		t.Error("isDir should be false for a file")
	}
}

func TestFileExists_NotExists(t *testing.T) {
	dir := t.TempDir()
	if fileExists(filepath.Join(dir, "nope.txt")) {
		t.Error("fileExists should be false for missing file")
	}
}

func TestExtractZip(t *testing.T) {
	dir := t.TempDir()
	zipPath := filepath.Join(dir, "test.zip")
	filePath := filepath.Join(dir, "file.txt")
	os.WriteFile(filePath, []byte("hello"), 0644)

	// Create a zip file
	zipFile, err := os.Create(zipPath)
	if err != nil {
		t.Fatalf("Failed to create zip: %v", err)
	}
	zw := NewTestZipWriter(zipFile)
	zw.AddFile("file.txt", []byte("hello"))
	zw.Close()
	zipFile.Close()

	extractDir := filepath.Join(dir, "extract")
	os.Mkdir(extractDir, 0755)
	if err := extractZip(zipPath, extractDir); err != nil {
		t.Fatalf("extractZip failed: %v", err)
	}
	if !fileExists(filepath.Join(extractDir, "file.txt")) {
		t.Error("Expected extracted file to exist")
	}
}

func TestExtractZip_Subfolder(t *testing.T) {
	dir := t.TempDir()
	zipPath := filepath.Join(dir, "test2.zip")
	zipFile, err := os.Create(zipPath)
	if err != nil {
		t.Fatalf("Failed to create zip: %v", err)
	}
	zw := NewTestZipWriter(zipFile)
	zw.AddFile("subdir/file.txt", []byte("data"))
	zw.Close()
	zipFile.Close()

	extractDir := filepath.Join(dir, "extract2")
	os.Mkdir(extractDir, 0755)
	if err := extractZip(zipPath, extractDir); err != nil {
		t.Fatalf("extractZip failed: %v", err)
	}
	if !fileExists(filepath.Join(extractDir, "subdir", "file.txt")) {
		t.Error("Expected extracted subfolder file to exist")
	}
}

// TestZipWriter is a helper for creating zip files in tests.
type TestZipWriter struct {
	zw *zip.Writer
}

func NewTestZipWriter(f *os.File) *TestZipWriter {
	return &TestZipWriter{zw: zip.NewWriter(f)}
}

func (t *TestZipWriter) AddFile(name string, data []byte) {
	w, _ := t.zw.Create(name)
	w.Write(data)
}

func (t *TestZipWriter) Close() {
	t.zw.Close()
}

func BenchmarkContains(b *testing.B) {
	slice := make([]string, 150000)
	for i := 0; i < 150000; i++ {
		slice[i] = "mod" + string(rune(i))
	}
	target := "mod149999"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Contains(slice, target)
	}
}

func BenchmarkSliceToSet(b *testing.B) {
	slice := make([]string, 150000)
	for i := 0; i < 150000; i++ {
		slice[i] = "mod" + string(rune(i))
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sliceToSet(slice)
	}
}

func BenchmarkSetLookup(b *testing.B) {
	slice := make([]string, 150000)
	for i := 0; i < 150000; i++ {
		slice[i] = "mod" + string(rune(i))
	}
	set := sliceToSet(slice)
	target := "mod149999"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = set[target]
	}
}
