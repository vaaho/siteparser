package core

// Результат парсинга сайта
type ParseData struct {
	Numbers []string `json:"numbers"` // список телфонных номеров
}

func NewParseData() ParseData {
	return ParseData{}
}
