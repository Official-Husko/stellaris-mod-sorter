# 🪐 Stellaris `.mod` and `descriptor.mod` File Format

<img src="https://tuttu.github.io/StellarisAssets/img/Stellaris-Icon-Small.png" height="19"/> ![Version](https://img.shields.io/badge/version-v1.0.0-blue?style=flat-square)

A `.mod` or `descriptor.mod` file defines the metadata and configuration for a Stellaris mod. It is used by the game launcher to display mod information, manage dependencies, and determine load order. This document provides a comprehensive reference for the format, based on official documentation and real-world examples.

---

## Example

```plaintext
name="SomeMod"
path="mod/SomeMod"
dependencies={
    "othermod"
    "another mod"
}
tags={
    "Graphics"
    "Economy"
    "Overhaul"
}
picture="thumbnail.png"
remote_file_id="1234567890"
version="1.0"
supported_version="v3.12.*"
```

---

## Field Reference

| Field               | Required     | Type         | Description                                                                 |
|---------------------|-------------|--------------|-----------------------------------------------------------------------------|
| 🏷️ `name`              | **Yes**     | String       | The display name of the mod as shown in the launcher.                       |
| 📁 `path`              | **Yes\***   | String       | Path to the mod’s folder. Required in `.mod`, ignored in `descriptor.mod`.  |
| 🔗 `dependencies`      | No          | List[String] | List of mod names/IDs this mod should load after (for submods/patches).     |
| 🏷️ `tags`              | No          | List[String] | List of tags for categorization (max 10, quoted if spaces).                 |
| 🖼️ `picture`           | No          | String       | Path to thumbnail image (usually `thumbnail.png`).                          |
| 🌐 `remote_file_id`    | No          | String/Int   | Steam Workshop file ID (added by launcher, can be ignored).                 |
| 🏷️ `version`           | No          | String       | Mod version (displayed in launcher, not game version).                      |
| 🎮 `supported_version` | Recommended | String       | Game version(s) this mod supports (e.g., `v3.12.*`).                        |

> \* `path` is required in `.mod` files, but ignored in `descriptor.mod`.

---

## Field Details

### `name`

The mod's display name (**required**).

- Example: `name="My Stellaris Mod"`

### `path`

Path to the mod’s folder (**required in `.mod`**, ignored in `descriptor.mod`).

- Can be absolute or relative to the Stellaris mods directory.
- Use forward slashes `/` (not backslashes).
- Example: `path="mod/MyStellarisMod"`

### `dependencies`

List of mods that must be loaded before this mod.

- Useful for sub-mods or compatibility patches.
- Each dependency is a string (mod name or ID).

```plaintext
dependencies={
    "My Other Stellaris Mod"
    "Not My Stellaris Mod"
}
```

### `tags`

List of tags for Steam Workshop and launcher categorization.

- Max 10 tags (launcher may enforce this).
- Use quotes for tags with spaces.

```plaintext
tags={
    "Graphics"
    "Economy"
    "Overhaul"
}
```

### `picture`

Path to a thumbnail image (usually `thumbnail.png`).

- Example: `picture="thumbnail.png"`

### `remote_file_id`

Steam Workshop file ID (added by launcher, can be ignored).

- Example: `remote_file_id="1234567890"`

### `version`

Mod version (displayed in launcher, not the game version).

- Any string is accepted.
- Example: `version="1.0"`

### `supported_version`

Specifies which Stellaris versions the mod supports.

- The last number can be replaced with `*` (wildcard).
- Example: `supported_version="v3.13.*"`

---

## Notes

- Fields can appear in any order.
- The file uses a simple key-value and block structure (not strict JSON or XML).
- Strings are usually quoted with double quotes (`"`).
- Blocks (like `dependencies` and `tags`) use curly braces `{ ... }` and list one item per line.
- Comments are not officially supported, but some launchers may ignore lines starting with `#` or `//`.
- The Paradox Launcher will automatically correct a relative path to the corresponding absolute path.
- For Steam Workshop: Using tags besides the predefined ones may prevent uploading on Paradox Mods. Max 10 tags allowed (since launcher version 2022.10).

---

## References

- [Stellaris Wiki: Modding](https://stellaris.paradoxwikis.com/Modding)
- [Stellaris Wiki: Mod structure](https://stellaris.paradoxwikis.com/Mod_structure)
- [Paradox Forum: Modding Documentation](https://forum.paradoxplaza.com/forum/forums/stellaris-modding-den.943/)

---

*This document describes the structure and fields of a Stellaris `.mod`/`descriptor.mod` file, based on real examples and official documentation.*
