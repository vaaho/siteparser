package io

import (
	"siteparser/core"
	"strings"
)

type ExportColumns struct {
	Total   int
	Numbers bool
}

func NewExportColumns(columns string) ExportColumns {
	result := ExportColumns{}

	// инициализируем колокни
	for _, column := range strings.Split(columns, ",") {
		switch column {
		case "numbers":
			result.Numbers = true
			result.Total++
		}
	}

	return result
}

func (c *ExportColumns) GetData(data *core.ParseData) []string {
	result := make([]string, 0)

	if c.Numbers {
		result = append(result, c.getNumbersData(data))
	}

	return result
}

func (c *ExportColumns) getNumbersData(data *core.ParseData) string {
	return strings.Join(data.Numbers, ",")
}
