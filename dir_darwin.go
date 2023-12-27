package main

import (
	"path/filepath"
)

var sp = ""

func jetbrainsDir() string {
	return filepath.Join(userHomeDir, "Library/Application Support", "JetBrains")
}
