package parser

// Парсит системы из метаданных по категориям
// Пример результата: {"crm": ["WordPress", "Tilda"]}
func parseCategories(content string) map[string][]string {
	result := make(map[string][]string)

	for _, meta := range GetMetaCollection().GetAll() {
		if meta.Pattern.MatchString(content) {
			result[meta.Category] = append(result[meta.Category], meta.Name)
		}
	}

	return result
}
