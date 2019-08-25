package main

import (
	"siteparser/core"
	"siteparser/downloader"
)

// Утилита для скачивания сайтов.
//
// Принимает на вход файл input.csv с доменами сайтов,
// а на выходе формирует папку sites с файлами,
// в каждом из которых результат скачивания главной страницы сайта.
//
// Скачивание происходит  в несколько потоков и может занять продолжительное время.
func main() {
	config := core.ParseConfigFromFlags()
	storage := core.NewSiteStorage(config.SitesDir)
	downloader.Download(config, storage)
}
