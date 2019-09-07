package io

import (
	"siteparser/core"
	"siteparser/parser"
	"strings"
)

type ExportColumns struct {
	Total      int
	Numbers    bool
	Categories map[string]bool
}

func NewExportColumns(columns string) ExportColumns {
	result := ExportColumns{Categories: make(map[string]bool)}
	meta := parser.GetMetaCollection()

	// инициализируем колокни
	for _, column := range strings.Split(columns, ",") {
		switch column {
		case "numbers": // парсинг номеров
			result.Numbers = true
			result.Total++
		default: // парсинг систем, представленных в категориях мета данных
			if meta.HasCategory(column) {
				result.Categories[column] = true
				result.Total += meta.GetCategoryLength(column)
			}
		}
	}

	return result
}

// Доп.колонки CSV с названиями колонок
func (c *ExportColumns) GetColumnsData() []string {
	result := make([]string, 0, c.Total)

	// добавляем название колонки с номерами
	if c.Numbers {
		result = append(result, "Номера")
	}

	// добавляем название колонок по категориям
	for category, _ := range c.Categories {
		for _, meta := range parser.GetMetaCollection().GetCategory(category) {
			result = append(result, meta.Name)
		}
	}

	return result
}

// Доп.колонки CSV с данными по колонкам
func (c *ExportColumns) GetData(data *core.ParseData) []string {
	result := make([]string, 0, c.Total)

	// добавляем колонку с номерами
	if c.Numbers {
		result = append(result, c.getNumbersData(data))
	}

	// добавляем колонки с системами по категориям
	for category, _ := range c.Categories {
		for _, meta := range parser.GetMetaCollection().GetCategory(category) {
			result = append(result, c.getMetaData(data, meta))
		}
	}

	return result
}

// Возвращает список номеров разделённый запятой, иначе пустая строка
func (c *ExportColumns) getNumbersData(data *core.ParseData) string {
	return strings.Join(data.Numbers, ",")
}

// Возвращает "1", если ситема найдена, иначе пустая строка
func (c *ExportColumns) getMetaData(data *core.ParseData, meta *parser.Meta) string {
	for _, name := range data.Categories[meta.Category] {
		if name == meta.Name {
			return "1"
		}
	}
	return ""
}
