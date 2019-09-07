package main

import (
	"siteparser/core"
	"siteparser/parser"
)

// Утилита для парсинга с сайтов полезной информации.
//
// Пробегается по папке `./sites/` со скаченными сайтами и для каждого сайта парсин данные.
// Результат парсинга сохраняется в том же JSON файле.
func main() {
	config := core.ParseConfigFromFlags()
	storage := core.NewSiteStorage(config.SitesDir)
	parser.Parse(config, storage)
}
