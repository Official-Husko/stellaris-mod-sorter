package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"stellaris-mod-sorter-go/internal/config"
	"stellaris-mod-sorter-go/internal/mods"
	prettylog "stellaris-mod-sorter-go/internal/utils"
)
func main() {
	var rootCmd = &cobra.Command{
		Use:   "stellaris-mod-sorter",
		Short: "Stellaris Mod Sorter and Manager",
		Long:  `A CLI tool for sorting, validating, and managing Stellaris mods and registries.`,
		Run: func(cmd *cobra.Command, args []string) {
			// Default mode: run the original mod sorting logic
			const modsRegistry = "mods_registry.json"
			const bakExt = ".bak"

			settingsPath, err := config.FindStellarisPath(modsRegistry)
			if err != nil {
				prettylog.PrintError("main", err, fmt.Sprintf("Unable to locate %s", modsRegistry), true)
			}
			prettylog.PrintPretty("main", fmt.Sprintf("Found Stellaris settings at %s", settingsPath), prettylog.LogInfo)

			// Load enabled mods and display order
			enabledModsRaw, dlcLoadPath := mods.LoadJsonOrder(settingsPath, "dlc_load.json", bakExt)
			displayOrderRaw, gameDataPath := mods.LoadJsonOrder(settingsPath, "game_data.json", bakExt)

			// Load mods registry
			modsRegistryPath := settingsPath + string(os.PathSeparator) + modsRegistry
			file, err := os.Open(modsRegistryPath)
			if err != nil {
				prettylog.PrintError("main", err, fmt.Sprintf("Could not open %s", modsRegistryPath), true)
			}
			defer file.Close()

			var data map[string]map[string]interface{}
			if err := mods.DecodeJSON(file, &data); err != nil {
				prettylog.PrintError("main", err, fmt.Sprintf("Could not decode %s", modsRegistryPath), true)
			}

			modList := mods.GetModList(data)
			modList = mods.TweakModOrder(modList)

			var idList []string
			if raw, ok := enabledModsRaw["enabled_mods"]; ok {
				if arr, ok := raw.([]interface{}); ok {
					for _, v := range arr {
						if s, ok := v.(string); ok {
							idList = append(idList, s)
						}
					}
				}
			}
			if len(idList) == 0 {
				prettylog.PrintPretty("main", "No enabled_mods found in dlc_load.json", prettylog.LogWarning)
				os.Exit(1)
			}

			allTags := make(map[string][]string)
			mods.GetModDescription(modList, data, allTags, settingsPath)
			modList = mods.SortAfterTags(allTags, modList)
			modList = mods.SpecialOrder(modList)
			modList = mods.SortDependencies(modList, idList, data)

			// Update and write output files
			displayOrderRaw["modsOrder"] = mods.GetModHashKeys(modList)
			enabledModsRaw["enabled_mods"] = mods.GetModIdsReversed(modList, idList)
			mods.WriteJsonOrder(enabledModsRaw, dlcLoadPath, bakExt)
			mods.WriteJsonOrder(displayOrderRaw, gameDataPath, bakExt)

			for i, mod := range modList {
				prettylog.PrintPretty("main", fmt.Sprintf("%d: %s", i, mod.SortedKey), prettylog.LogMessage)
			}
			prettylog.PrintPretty("main", "done", prettylog.LogInfo)
		},
	}

	rootCmd.AddCommand(
		&cobra.Command{
			Use:   "validate-json <json> <schema>",
			Short: "Validate a JSON file against a JSON Schema",
			Args:  cobra.ExactArgs(2),
			RunE: func(cmd *cobra.Command, args []string) error {
				return mods.ValidateJSONSchema(args[0], args[1])
			},
		},
		&cobra.Command{
			Use:   "backup-registry <src> <dst>",
			Short: "Backup the official mods_registry.json",
			Args:  cobra.ExactArgs(2),
			RunE: func(cmd *cobra.Command, args []string) error {
				return mods.BackupFile(args[0], args[1])
			},
		},
		&cobra.Command{
			Use:   "dry-run",
			Short: "Perform a dry run of the mod sorting process (no changes written)",
			Run: func(cmd *cobra.Command, args []string) {
				fmt.Println("[DRY RUN] Simulating mod sorting. No changes will be written.")
				// TODO: Implement dry-run logic
			},
		},
		&cobra.Command{
			Use:   "custom-stellaris-path <path>",
			Short: "Set a custom Stellaris user data path",
			Args:  cobra.ExactArgs(1),
			Run: func(cmd *cobra.Command, args []string) {
				fmt.Printf("Custom Stellaris path set to: %s\n", args[0])
				// TODO: Implement logic to use custom path
			},
		},
		&cobra.Command{
			Use:   "validate",
			Short: "Validate the official mods_registry.json against the schema",
			RunE: func(cmd *cobra.Command, args []string) error {
				// TODO: Detect official registry path dynamically
				return mods.ValidateJSONSchema("/path/to/official/mods_registry.json", "mods_registry.schema.json")
			},
		},
		&cobra.Command{
			Use:   "backup <dst>",
			Short: "Backup the official mods_registry.json",
			Args:  cobra.ExactArgs(1),
			RunE: func(cmd *cobra.Command, args []string) error {
				// TODO: Detect official registry path dynamically
				return mods.BackupFile("/path/to/official/mods_registry.json", args[0])
			},
		},
	)

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
