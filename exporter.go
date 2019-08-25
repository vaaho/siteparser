package main

import (
	"siteparser/core"
	"siteparser/io"
)

func main() {
	config := core.ParseConfigFromFlags()
	storage := core.NewSiteStorage(config.SitesDir)
	io.Export(config, storage)
}
