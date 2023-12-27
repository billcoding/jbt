package main

import (
	"path/filepath"
)

var sp = "64.exe"

func jetbrainsDir() string {
	return filepath.Join(userHomeDir, "AppData", "Roaming", "JetBrains")
}
