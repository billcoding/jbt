package main

import (
	"embed"
)

var (
	//go:embed config/dns.conf
	configDnsConfFs []byte
	//go:embed config/power.conf
	configPowerConfFs []byte
	//go:embed config/url.conf
	configUrlConfFs []byte

	//go:embed plugins/dns.jar
	pluginsDnsJarFs []byte
	//go:embed plugins/hideme.jar
	pluginsHidemeJarFs []byte
	//go:embed plugins/power.jar
	pluginsPowerJarFs []byte
	//go:embed plugins/url.jar
	pluginsUrlJarFs []byte

	//go:embed active-agt.jar
	activeAgtJarFs []byte

	//go:embed keys
	keysFs embed.FS
)
