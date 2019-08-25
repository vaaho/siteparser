package io

import (
	"bufio"
	"log"
	"os"
	"siteparser/core"
	"strings"
)

func ReadInputByLine(sourceFile string) <-chan string {
	out := make(chan string)

	go func() {
		file, err := os.Open(sourceFile)
		core.FailOnError(err)
		defer file.Close()

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			line := scanner.Text()
			out <- line
		}

		close(out)
	}()

	return out
}

func WriteOutputByLine(destFile string, lines <-chan string) <-chan string {
	out := make(chan string)

	go func() {
		file, err := os.Create(destFile)
		core.FailOnError(err)
		defer file.Close()

		for line := range lines {
			_, err := file.WriteString(line + "\n")
			core.FailOnError(err)
			out <- line
		}

		close(out)
	}()

	return out
}

func AppendExportColumns(lines <-chan string, columns ExportColumns, storage *core.SiteStorage) <-chan string {
	out := make(chan string)

	go func() {
		for line := range lines {
			outLine := line + GetExportColumnsData(line, columns, storage)
			//log.Printf("[LINE] " + outLine)
			out <- outLine
		}
		close(out)
	}()

	return out
}

func GetExportColumnsData(line string, columns ExportColumns, storage *core.SiteStorage) string {
	domain := line // TODO: В будущем сделать более сложную функцию извлечения домена из CSV строчки
	if domain == "" {
		return strings.Repeat(core.CsvSeparator, columns.Total)
	}

	site := storage.Load(domain)
	if site.Status != core.StatusSuccess || site.ParseData == nil {
		return strings.Repeat(core.CsvSeparator, columns.Total)
	}

	// берём спаршенные значения для колонок
	columnsData := columns.GetData(site.ParseData)

	result := core.CsvSeparator + strings.Join(columnsData, core.CsvSeparator)
	return result
}

func Export(config *core.Config, storage *core.SiteStorage) {
	columns := NewExportColumns(config.ExportColumns)
	if columns.Total == 0 {
		log.Printf("[EXPORT] Нечего экспортировать")
		return
	}

	lines := ReadInputByLine(config.InputFile)
	lines = AppendExportColumns(lines, columns, storage)
	lines = WriteOutputByLine(config.OutputFile, lines)

	for _ = range lines {
	}
}
