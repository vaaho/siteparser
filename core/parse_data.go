package core

// Результат парсинга сайта
type ParseData struct {
	Numbers    []string            `json:"numbers"`    // список телфонных номеров
	Categories map[string][]string `json:"categories"` // спаршенные категории и системы в них, пример: {"crm": ["WordPress", "Tilda"]}
}

func NewParseData() ParseData {
	return ParseData{}
}
