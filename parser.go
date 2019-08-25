package main

import (
	"siteparser/core"
	"siteparser/parser"
)

func main() {
	config := core.ParseConfigFromFlags()
	storage := core.NewSiteStorage(config.SitesDir)
	parser.Parse(config, storage)
}
