# 🗂️ Stellaris `mods_registry.json` File Format

<img src="https://tuttu.github.io/StellarisAssets/img/Stellaris-Icon-Small.png" height="19"/> ![Version](https://img.shields.io/badge/version-v1.0.0-blue?style=flat-square)

A `mods_registry.json` file stores metadata for all discovered Stellaris mods, typically as a map of unique mod IDs to their properties. This registry is used by mod managers and tools to track mod locations, display names, Steam IDs, tags, and other relevant information for sorting, launching, or managing mods.

---

## Example

```json
{
    "21153e40-4eea-4b4e-bae1-59ec0ccc8016": {
        "dirPath": "/path/to/mod",
        "displayName": "My Mod Name",
        "gameRegistryId": "mod/ugc_1234567890.mod",
        "id": "21153e40-4eea-4b4e-bae1-59ec0ccc8016",
        "requiredVersion": "v4.0.*",
        "source": "steam",
        "status": "ready_to_play",
        "steamId": "1234567890",
        "tags": ["Graphics", "Gameplay"]
    }
}
```

---

## Field Reference

| Field                 | Required    | Type           | Description                                                        |
|-----------------------|-------------|----------------|--------------------------------------------------------------------|
| 🆔 `id`               | **Yes**     | String         | Unique mod identifier (UUID or hash).                              |
| 🏷️ `displayName`      | **Yes**     | String         | The display name of the mod.                                       |
| 📁 `dirPath`          | **Yes**     | String         | Absolute path to the mod's directory.                              |
| 🗂️ `gameRegistryId`   | **Yes**     | String         | Path to the mod's `.mod` descriptor file.                          |
| 🏷️ `requiredVersion`  | No          | String         | Required Stellaris version for this mod (e.g., `v4.0.*`).          |
| 🌐 `source`           | **Yes**     | String         | Source of the mod (`steam`, `local`, etc.).                        |
| 📦 `status`           | No          | String         | Status of the mod (e.g., `ready_to_play`, `missing`, etc.).        |
| 🌐 `steamId`          | No          | String/Int     | Steam Workshop file ID (if applicable).                            |
| 🏷️ `tags`             | No          | List[String]   | List of tags for categorization.                                   |

---

## Field Details

### 🆔 `id`

A unique identifier for the mod (usually a UUID or hash).

### 🏷️ `displayName`

The mod's display name as shown in the launcher or manager.

### 📁 `dirPath`

Absolute path to the mod's directory on disk.

### 🗂️ `gameRegistryId`

Path to the mod's `.mod` descriptor file (relative or absolute).

### 🏷️ `requiredVersion`

The Stellaris version required by this mod (e.g., `v4.0.*`).

### 🌐 `source`

Where the mod was sourced from (`steam`, `local`, etc.).

### 📦 `status`

Current status of the mod (e.g., `ready_to_play`, `missing`).

### 🌐 `steamId`

Steam Workshop file ID (if the mod is from Steam Workshop).

### 🏷️ `tags`

List of tags for categorization and filtering.

---

## Notes

- The file is a JSON object mapping unique mod IDs to their metadata objects.
- All paths should be absolute for reliability.
- Fields may vary depending on the mod source and manager implementation.
- Additional fields may be present for advanced features or compatibility.

---

## References

- [Stellaris Wiki: Modding](https://stellaris.paradoxwikis.com/Modding)
- [Stellaris Wiki: Mod structure](https://stellaris.paradoxwikis.com/Mod_structure)
- [Paradox Forum: Modding Documentation](https://forum.paradoxplaza.com/forum/forums/stellaris-modding-den.943/)

---

*This document describes the structure and fields of a Stellaris `mods_registry.json` file, based on real examples and best practices.*
