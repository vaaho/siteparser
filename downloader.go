package main

import (
	"siteparser/core"
	"siteparser/downloader"
)

func main() {
	config := core.ParseConfigFromFlags()
	storage := core.NewSiteStorage(config.SitesDir)
	downloader.Download(config, storage)
}
