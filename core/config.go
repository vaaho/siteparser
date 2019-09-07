package core

import "flag"

type Config struct {
	InputFile     string
	OutputFile    string
	SitesDir      string
	MaxThreads    int
	LogFails      bool
	NoColumnsRow  bool
	ExportColumns string
}

func ParseConfigFromFlags() *Config {
	config := &Config{}

	flag.StringVar(&config.InputFile, "input", "./input.csv", "Файл со списком доменов для парсина")
	flag.StringVar(&config.OutputFile, "output", "./output.csv", "Файл с результатами парсина")
	flag.StringVar(&config.SitesDir, "sites", "./sites/", "Папка для сохранения результатов парсинга сайтов")
	flag.IntVar(&config.MaxThreads, "threads", 1000, "Максимально число паралельно скачиваемых сайтов")
	flag.BoolVar(&config.LogFails, "logFails", false, "Признак логирования ошибок скачивания сайтов")
	flag.BoolVar(&config.NoColumnsRow, "noColumnsRow", false, "Срока с названием колонок отсутсвует")
	flag.StringVar(&config.ExportColumns, "export", "", "Набор колонок через запятую для выгрузки")

	flag.Parse()

	return config
}
