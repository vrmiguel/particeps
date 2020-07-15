package particeps

import (
	"os"
	"path/filepath"
	"runtime"
)

// GetPrefFolder returns the preference (not system-wide) for the three major OSes
func GetPrefFolder() string {
	switch runtime.GOOS {
	case "linux": // Preference folder for Linux
		if os.Getenv("XDG_CONFIG_HOME") != "" {
			return os.Getenv("XDG_CONFIG_HOME")
		}
		return filepath.Join(os.Getenv("HOME"), ".config")
	case "windows": // Preference folder for Windows
		return os.Getenv("APPDATA")
	case "darwin": // Preference folder for Darwin & macOS
		return os.Getenv("HOME") + "/Library/Application Support"
	default: // I don't know what're the config. folders for the other OSes. Please PR if you do.
		return "./"
	}
}
