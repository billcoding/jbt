package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

var (
	userHomeDir, _ = os.UserHomeDir()

	appMap = map[string]string{
		"goland":    "GoLand",
		"idea":      "IntelliJIdea",
		"pycharm":   "PyCharm",
		"webstorm":  "WebStorm",
		"phpstorm":  "PhpStorm",
		"rider":     "Rider",
		"clion":     "CLion",
		"dataspell": "DataSpell",
		"datagrip":  "DataGrip",
	}
)

func copyJetBrainsFiles() {
	dir := jetbrainsDir()
	if _, err := os.Stat(dir); err != nil {
		fmt.Println("Not found JetBrains any products")
		os.Exit(1)
	}
	fn := func(name string, buf []byte) {
		if _, err := os.Stat(name); err == nil {
			return
		}
		_ = os.WriteFile(name, buf, os.ModePerm)
	}
	configDir := filepath.Join(dir, "config")
	pluginsDir := filepath.Join(dir, "plugins")
	_ = os.MkdirAll(configDir, 0700)
	_ = os.MkdirAll(pluginsDir, 0700)
	ddd := func(name string) string { return filepath.Join(dir, name) }
	ccc := func(name string) string { return filepath.Join(configDir, name) }
	ppp := func(name string) string { return filepath.Join(pluginsDir, name) }

	fn(ccc("dns.conf"), configDnsConfFs)
	fn(ccc("power.conf"), configPowerConfFs)
	fn(ccc("url.conf"), configUrlConfFs)

	fn(ppp("dns.jar"), pluginsDnsJarFs)
	fn(ppp("hideme.jar"), pluginsHidemeJarFs)
	fn(ppp("power.jar"), pluginsPowerJarFs)
	fn(ppp("url.jar"), pluginsUrlJarFs)

	fn(ddd("active-agt.jar"), activeAgtJarFs)
}

func copyAppKey(name string) {
	buf := appKeyBytes(name)
	keyDirs := appKeyDirs(name)
	for _, d := range keyDirs {
		_ = os.Remove(d)
		_ = os.WriteFile(d, buf, os.ModePerm)
	}
}

func copyAppVmOptions(name string) {
	for _, d := range appVmOptionsDirs(name) {
		_ = os.WriteFile(d, []byte(getVmOptionsContent()), os.ModePerm)
	}
}

func appDataDirs(name string) []string {
	dir := jetbrainsDir()
	readDirs, _ := os.ReadDir(dir)
	if readDirs == nil || len(readDirs) <= 0 {
		return nil
	}
	var ds []string
	for _, d := range readDirs {
		if d.IsDir() && strings.HasPrefix(d.Name(), name) {
			ds = append(ds, filepath.Join(dir, d.Name()))
		}
	}
	if len(ds) <= 0 {
		fmt.Println("Not found JetBrains " + name + " installation")
		os.Exit(1)
	}
	return ds
}

func appKeyBytes(name string) []byte {
	buf, _ := keysFs.ReadFile("keys/" + name + ".key")
	return buf
}

func appVmOptionsDirs(name string) []string {
	var ds []string
	for _, d := range appDataDirs(getAppName(name)) {
		ds = append(ds, filepath.Join(d, name+sp+".vmoptions"))
	}
	return ds
}

func appKeyDirs(name string) []string {
	var ds []string
	for _, d := range appDataDirs(getAppName(name)) {
		ds = append(ds, filepath.Join(d, name+".key"))
	}
	return ds
}

func getVmOptionsContent() string {
	jar := filepath.Join(jetbrainsDir(), "active-agt.jar")
	return fmt.Sprintf(`-javaagent:` + jar + `
--add-opens=java.base/jdk.internal.org.objectweb.asm=ALL-UNNAMED
--add-opens=java.base/jdk.internal.org.objectweb.asm.tree=ALL-UNNAMED
`)
}

func getAppName(name string) string { return appMap[name] }
