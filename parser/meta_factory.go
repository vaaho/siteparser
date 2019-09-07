package parser

var metaCollectionInstance *MetaCollection

// Синглтон-коллекция метаданных
func GetMetaCollection() *MetaCollection {
	if metaCollectionInstance == nil {
		instance := newMetaCollection()
		initMeta(instance)
		metaCollectionInstance = instance
	}
	return metaCollectionInstance
}

// Инициализация метаданных
func initMeta(c *MetaCollection) {

	// Примеры CRM:
	// WordPress   https://orekh.design/ http://www.zavodstil.ru/
	// 1С-Битрикс  https://mango-office.ru/
	// OpenCart    https://dolinaroz52.ru/
	// Tilda       http://shoqan.com/
	// InSales     https://www.bakerhouse.ru/
	// NetCat      https://startshina.ru/

	c.add("crm", "WordPress", `(wordpress|WordPress)`)
	c.add("crm", "1С-Битрикс", `\/bitrix\/`)
	c.add("crm", "OpenCart", `addToCart`)
	c.add("crm", "Tilda", `static\.tildacdn\.com`)
	c.add("crm", "InSales", `insales\.ru`)
	c.add("crm", "NetCat", `netcat_files`)
}
