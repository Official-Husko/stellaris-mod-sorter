package mods

import (
	"fmt"
)

// CLICommand represents a command that can be run from the CLI.
type CLICommand struct {
	Name        string
	Description string
	Handler     func(args []string) error
}

// AvailableCommands lists all supported CLI commands for the app.
var AvailableCommands = []CLICommand{
	{
		Name:        "validate-json",
		Description: "Validate a JSON file against a JSON Schema (usage: validate-json <json> <schema>)",
		Handler:     ValidateJSONCommand,
	},
	{
		Name:        "backup-registry",
		Description: "Backup the official mods_registry.json (usage: backup-registry <src> <dst>)",
		Handler:     BackupRegistryCommand,
	},
	{
		Name:        "dry-run",
		Description: "Perform a dry run of the mod sorting process (no changes written)",
		Handler:     DryRunCommand,
	},
	{
		Name:        "custom-stellaris-path",
		Description: "Set a custom Stellaris user data path (usage: custom-stellaris-path <path>)",
		Handler:     CustomStellarisPathCommand,
	},
	{
		Name:        "validate",
		Description: "Validate the official mods_registry.json against the schema",
		Handler:     ValidateOfficialRegistryCommand,
	},
	{
		Name:        "backup",
		Description: "Backup the official mods_registry.json (usage: backup <dst>)",
		Handler:     BackupOfficialRegistryCommand,
	},
	// Add more commands here as needed
}

// ValidateJSONCommand validates a JSON file against a schema.
func ValidateJSONCommand(args []string) error {
	if len(args) < 2 {
		return fmt.Errorf("usage: validate-json <json> <schema>")
	}
	return ValidateJSONSchema(args[0], args[1])
}

// BackupRegistryCommand backs up the mods registry file.
func BackupRegistryCommand(args []string) error {
	if len(args) < 2 {
		return fmt.Errorf("usage: backup-registry <src> <dst>")
	}
	return BackupFile(args[0], args[1])
}

// DryRunCommand simulates the mod sorting process without writing changes.
func DryRunCommand(args []string) error {
	fmt.Println("[DRY RUN] Simulating mod sorting. No changes will be written.")
	// TODO: Implement dry-run logic
	return nil
}

// CustomStellarisPathCommand sets a custom Stellaris user data path.
func CustomStellarisPathCommand(args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("usage: custom-stellaris-path <path>")
	}
	fmt.Printf("Custom Stellaris path set to: %s\n", args[0])
	// TODO: Implement logic to use custom path
	return nil
}

// ValidateOfficialRegistryCommand validates the official mods_registry.json against the schema.
func ValidateOfficialRegistryCommand(args []string) error {
	// TODO: Detect official registry path dynamically
	return ValidateJSONSchema("/path/to/official/mods_registry.json", "mods_registry.schema.json")
}

// BackupOfficialRegistryCommand backs up the official mods_registry.json.
func BackupOfficialRegistryCommand(args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("usage: backup <dst>")
	}
	// TODO: Detect official registry path dynamically
	return BackupFile("/path/to/official/mods_registry.json", args[0])
}
