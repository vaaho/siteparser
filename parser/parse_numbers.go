package parser

import (
	"regexp"
)

var (
	notDigitRegExp *regexp.Regexp // выражение для очистки строки от не цифр
	numbersRegExp  *regexp.Regexp // выражение для парсинка телефонных номеров
)

func init() {
	notDigitRegExp = regexp.MustCompile(`[^0-9]+`)
	numbersRegExp = regexp.MustCompile(`[^0-9](\+?[78][\- ]?\(?\d{3}\)?[\- ]?[\d\- ]{7,10})[^0-9]`)
}

// Парсит номера телефонов с сайта
func parseNumbers(content string) []string {
	items := numbersRegExp.FindAllStringSubmatch(content, -1)

	// сперва сохраняем номера в хеш, чтобы не было повторов
	numbersMap := make(map[string]bool)
	for _, item := range items {
		// сохрянем номера в формате did, т.е. 10 цифр с кодом 7
		if number := ensureDid(item[1]); number != "" {
			numbersMap[number] = true
		}
	}

	// формируем плоский список номеров без повторов
	numbers := make([]string, 0, len(numbersMap))
	for number, _ := range numbersMap {
		numbers = append(numbers, number)
	}

	return numbers
}

// Извлекает из строки did номер в формате: код 7 и 10 цифр, если не получается - то пустую строку
func ensureDid(number string) string {
	did := notDigitRegExp.ReplaceAllString(number, "") // номер без лишних знаков
	if len(did) != 11 {                                // не является 10-значным номер с кодом 7 или 8
		return ""
	}
	return "7" + did[1:]
}
