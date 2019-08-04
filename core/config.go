package core

import "flag"

type Config struct {
	InputFile  string
	SitesDir   string
	MaxThreads int
	LogFails   bool
}

func ParseConfigFromFlags() *Config {
	config := &Config{}

	flag.StringVar(&config.InputFile, "input", "./input.csv", "Файл со списком доменов для парсина")
	flag.StringVar(&config.SitesDir, "sites", "./sites/", "Папка для сохранения результатов парсинга сайтов")
	flag.IntVar(&config.MaxThreads, "threads", 1000, "Максимально число паралельно скачиваемых сайтов")
	flag.BoolVar(&config.LogFails, "logFails", false, "Признак логирования ошибок скачивания сайтов")

	flag.Parse()

	return config
}
