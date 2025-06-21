package mods

// Mod represents a Stellaris mod and its metadata.
type Mod struct {
	HashKey     string
	Name        string
	ModId       string
	SortedKey   string
	Dependencies []string
}
