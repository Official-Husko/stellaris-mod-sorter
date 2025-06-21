# Stellaris Mod Sorter

A modern, fast, and reliable tool to manage and sort your Stellaris mods. Written in Go, this application reads your mod configuration, sorts mods based on dependencies, tags, and custom rules, and writes the correct load order for the game.

---

## ✨ Features

- **Automatic mod sorting** based on dependencies, tags, and special rules
- **Handles special cases** for known mods (e.g., UI Overhaul)
- **Backs up** your existing configuration before making changes
- **Cross-platform** (works anywhere Go runs)
- **Fast and robust** thanks to Go's concurrency and type safety

## 📁 Project Structure

```plaintext
stellaris-mod-sorter-go/
├── cmd/
│   └── main.go            # Application entry point
├── internal/
│   ├── mods/
│   │   ├── mod.go         # Mod struct and related types
│   │   ├── sorter.go      # Sorting, tag, and dependency logic
│   │   └── utils.go       # Utility functions (file, string, zip helpers)
│   └── config/
│       └── config.go      # Configuration and settings path detection
├── go.mod                 # Go module definition
├── README.md              # This documentation
└── old.py                 # Original Python script (for reference)
```

## 🚀 Installation

1. **Clone the repository:**

   ```sh
   git clone <repository-url>
   cd stellaris-mod-sorter-go
   ```

2. **Install Go:**
   - Download and install from [golang.org](https://golang.org/dl/)
3. **Download dependencies:**

   ```sh
   go mod tidy
   ```

## 🛠️ Usage

1. **Ensure your Stellaris mod configuration files** (`mods_registry.json`, `dlc_load.json`, `game_data.json`) are in the correct Paradox Interactive Stellaris directory (the tool will auto-detect this on Linux and Windows).
2. **Run the application:**

   ```sh
   go run ./cmd/main.go
   ```

   - The tool will print the detected settings path, process your mods, and output the new sorted order.
   - Backups of your original config files will be created with a `.bak` extension.

## 🤝 Contributing

Contributions, bug reports, and feature requests are welcome! Please open an issue or submit a pull request.

## 📄 License

MIT License. See the [LICENSE](LICENSE) file for details.

---

**Original Python version:** [haifengkao/StellairsLoadOrderFixer24](https://github.com/haifengkao/StellairsLoadOrderFixer24)

**Maintained by:** [Your Name or Organization]
