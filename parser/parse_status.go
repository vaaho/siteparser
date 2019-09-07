package parser

type ParseStatus struct {
	Total           int // Общее число скаченных сайтов
	Parsed          int // Число успешно скаченных и распарщенных сайтов
	NumbersCount    int // Число сайтов на которых были найдены телефонные номера
	CategoriesCount int // Число сайтов на которых были найдены систем из категорий
}

func NewParseStatus() *ParseStatus {
	return &ParseStatus{}
}
