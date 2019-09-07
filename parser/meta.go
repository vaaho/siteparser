package parser

import "regexp"

// Метаданные для парсинга, сгрупированные по категориям
type Meta struct {
	Category string
	Name     string
	Pattern  *regexp.Regexp
}

// Колекция с метаданными
type MetaCollection struct {
	items map[string]map[string]*Meta
}

func newMetaCollection() *MetaCollection {
	return &MetaCollection{
		items: make(map[string]map[string]*Meta),
	}
}

// Добавляет метаданные в коллекцию
func (c *MetaCollection) add(category string, name string, pattern string) *MetaCollection {
	group, ok := c.items[category]
	if !ok {
		group = make(map[string]*Meta)
		c.items[category] = group
	}
	group[name] = &Meta{
		Category: category,
		Name:     name,
		Pattern:  regexp.MustCompile(pattern),
	}
	return c
}

// Метаданные по катеогии и названию
func (c *MetaCollection) Get(category string, name string) *Meta {
	if group, ok := c.items[category]; ok {
		if meta, ok := group[name]; ok {
			return meta
		}
	}
	return nil
}

// Список метаданных по категории
func (c *MetaCollection) GetCategory(category string) []*Meta {
	group, ok := c.items[category]
	if !ok {
		return make([]*Meta, 0, 0)
	}
	result := make([]*Meta, 0, len(group))
	for _, meta := range group {
		result = append(result, meta)
	}
	return result
}

func (c *MetaCollection) HasCategory(category string) bool {
	_, ok := c.items[category]
	return ok
}

func (c *MetaCollection) GetCategoryLength(category string) int {
	if group, ok := c.items[category]; ok {
		return len(group)
	}
	return 0
}

// Список всех категорий
func (c *MetaCollection) GetCategories() []string {
	result := make([]string, 0, len(c.items))

	for category := range c.items {
		result = append(result, category)
	}

	return result
}

// Все метаданные
func (c *MetaCollection) GetAll() []*Meta {
	result := make([]*Meta, 0, len(c.items)*10)

	for _, group := range c.items {
		for _, meta := range group {
			result = append(result, meta)
		}
	}

	return result
}
