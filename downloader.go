package main

import (
	"siteparser/core"
	"siteparser/downloader"
)

// Утилита для скачивания сайтов.
//
// На вход принимает CSV файл `./input.csv` со списком доменов для скачивания.
// На выходе формирует папку `./sites/`, где для каждого домена создаётся JSON файл
// с результами скачки сайта. В нёмже сохраняется код ошибки, если скачивание не удалось.
//
// Скачивание происходит  в несколько потоков и может занять продолжительное время.
func main() {
	config := core.ParseConfigFromFlags()
	storage := core.NewSiteStorage(config.SitesDir)
	downloader.Download(config, storage)
}
