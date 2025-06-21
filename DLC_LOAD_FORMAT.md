# 🗂️ Stellaris `dlc_load.json` File Format

<img src="https://tuttu.github.io/StellarisAssets/img/Stellaris-Icon-Small.png" height="19"/> ![Version](https://img.shields.io/badge/version-v1.0.0-blue?style=flat-square)

A `dlc_load.json` file stores the enabled and disabled DLCs and mods for a Stellaris user profile. It is used by the game launcher to determine which mods and DLCs are active for the next game session.

---

## Example

```json
{
    "disabled_dlcs": [],
    "enabled_mods": [
        "mod/ugc_2407436476.mod",
        "mod/ugc_683230077.mod",
        "mod/ugc_3155789435.mod",
        // ... more mod descriptor paths ...
        "mod/Yellow Rain.mod"
    ]
}
```

---

## Field Reference

| Field           | Required | Type           | Description                                                      |
|-----------------|----------|----------------|------------------------------------------------------------------|
| `disabled_dlcs` | Yes      | List[String]   | List of DLC IDs that are disabled for this profile.              |
| `enabled_mods`  | Yes      | List[String]   | List of mod descriptor file paths (relative to mods folder).     |

---

## Field Details

### `disabled_dlcs`

- An array of DLC IDs (strings) that are currently disabled.
- Usually empty if all DLCs are enabled.
- Example: `["dlc001", "dlc002"]`

### `enabled_mods`

- An array of mod descriptor file paths (strings).
- Each entry is a path like `mod/ugc_1234567890.mod` or `mod/SomeMod.mod`.
- The order of this list determines the mod load order in Stellaris.
- Example: `["mod/ugc_2407436476.mod", "mod/Yellow Rain.mod"]`

---

## Notes

- The file is a simple JSON object with two fields: `disabled_dlcs` and `enabled_mods`.
- Paths in `enabled_mods` are relative to the Stellaris mods directory.
- The order of `enabled_mods` is important for mod load order and compatibility.
- This file is automatically updated by the Stellaris launcher when you enable/disable mods or DLCs.

---

## References

- [Stellaris Wiki: Modding](https://stellaris.paradoxwikis.com/Modding)
- [Stellaris Wiki: Mod structure](https://stellaris.paradoxwikis.com/Mod_structure)
- [Paradox Forum: Modding Documentation](https://forum.paradoxplaza.com/forum/forums/stellaris-modding-den.943/)

---

*This document describes the structure and fields of a Stellaris `dlc_load.json` file, based on real examples and official documentation.*
