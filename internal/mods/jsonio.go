package mods

import (
	"encoding/json"
	"io"
	"os"
	"path/filepath"

	"github.com/santhosh-tekuri/jsonschema/v6"

	prettylog "stellaris-mod-sorter-go/internal/utils"
)

// LoadJsonOrder reads a JSON file and returns its data and the file path.
func LoadJsonOrder(settingPath, file, bakExt string) (map[string]interface{}, string) {
	filePath := filepath.Join(settingPath, file)
	jsonData := make(map[string]interface{})
	if _, err := os.Stat(filePath); err == nil {
		// Remove old backup
		if _, err := os.Stat(filePath + bakExt); err == nil {
			prettylog.PrintPretty("LoadJsonOrder", "Removing old backup: "+filePath+bakExt, prettylog.LogInfo)
			os.Remove(filePath + bakExt)
		}
		content, err := os.ReadFile(filePath)
		if err != nil {
			prettylog.PrintPretty("LoadJsonOrder", "Loading failed: "+filePath, prettylog.LogError)
		} else {
			json.Unmarshal(content, &jsonData)
			if len(jsonData) < 1 {
				prettylog.PrintPretty("LoadJsonOrder", "Loading failed: "+filePath, prettylog.LogError)
			}
		}
	} else {
		prettylog.PrintPretty("LoadJsonOrder", "Please enable at least one mod: "+filePath, prettylog.LogError)
	}
	return jsonData, filePath
}

// WriteJsonOrder writes a JSON file, backing up the old one.
func WriteJsonOrder(data map[string]interface{}, file, bakExt string) {
	// Backup
	err := os.Rename(file, file+bakExt)
	if err != nil && !os.IsNotExist(err) {
		prettylog.PrintPretty("WriteJsonOrder", "Could not backup file: "+err.Error(), prettylog.LogWarning)
	}
	content, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		prettylog.PrintPretty("WriteJsonOrder", "Failed to marshal JSON: "+err.Error(), prettylog.LogError)
		return
	}
	err = os.WriteFile(file, content, 0644)
	if err != nil {
		prettylog.PrintPretty("WriteJsonOrder", "Failed to write file: "+err.Error(), prettylog.LogError)
	}
}

// DecodeJSON decodes JSON from an io.Reader into v.
func DecodeJSON(r io.Reader, v interface{}) error {
	err := json.NewDecoder(r).Decode(v)
	if err != nil {
		prettylog.PrintPretty("DecodeJSON", "JSON decode error: "+err.Error(), prettylog.LogError)
	}
	return err
}

// ValidateJSONSchema validates a JSON file against a JSON Schema file using santhosh-tekuri/jsonschema.
func ValidateJSONSchema(jsonPath, schemaPath string) error {
	compiler := jsonschema.NewCompiler()
	absSchemaPath, err := filepath.Abs(schemaPath)
	if err != nil {
		prettylog.PrintPretty("ValidateJSONSchema", "Schema path error: "+err.Error(), prettylog.LogError)
		return err
	}
	schema, err := compiler.Compile("file://" + absSchemaPath)
	if err != nil {
		prettylog.PrintPretty("ValidateJSONSchema", "Schema compile error: "+err.Error(), prettylog.LogError)
		return err
	}
	f, err := os.Open(jsonPath)
	if err != nil {
		prettylog.PrintPretty("ValidateJSONSchema", "Open JSON error: "+err.Error(), prettylog.LogError)
		return err
	}
	defer f.Close()
	var doc interface{}
	if err := json.NewDecoder(f).Decode(&doc); err != nil {
		prettylog.PrintPretty("ValidateJSONSchema", "JSON decode error: "+err.Error(), prettylog.LogError)
		return err
	}
	if err := schema.Validate(doc); err != nil {
		prettylog.PrintPretty("ValidateJSONSchema", "JSON validation error: "+err.Error(), prettylog.LogError)
		return err
	}
	return nil
}

// Standalone function to validate example_registry.json against mods_registry.schema.json for manual check
func ValidateExampleRegistry() error {
	return ValidateJSONSchema("example_registry.json", "mods_registry.schema.json")
}

// BackupFile copies the source file to the destination as a backup.
func BackupFile(src, dst string) error {
	input, err := os.ReadFile(src)
	if err != nil {
		prettylog.PrintPretty("BackupFile", "Backup read error: "+err.Error(), prettylog.LogError)
		return err
	}
	err = os.WriteFile(dst, input, 0644)
	if err != nil {
		prettylog.PrintPretty("BackupFile", "Backup write error: "+err.Error(), prettylog.LogError)
	}
	return err
}
