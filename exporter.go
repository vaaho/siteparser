package main

import (
	"siteparser/core"
	"siteparser/io"
)

// Утилита дляэкспорта спарщенных данных в CSV
//
// На вход принимает CSV файл `./input.csv` со списком доменов для скачивания.
// На выходе формирует CSV файл `./output.csv`, в котором колонки с результатами эпарсинга.
// В работе используются результаты парсинга из папки `./sites/`.
//
// Обязательным параметром програмы является опция `-export`, которая задаёт набор колонок для экспорта.
func main() {
	config := core.ParseConfigFromFlags()
	storage := core.NewSiteStorage(config.SitesDir)
	io.Export(config, storage)
}
