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
	// CRM
	c.add("crm", "WordPress", `(wordpress|WordPress)`)
	c.add("crm", "1С-Битрикс", `\/bitrix\/`)
	c.add("crm", "OpenCart", `addToCart`)
	c.add("crm", "Tilda", `static\.tildacdn\.com`)
	c.add("crm", "InSales", `insales\.ru`)
	c.add("crm", "NetCat", `netcat_files`)

	// Коллтрекинги
	c.add("ct", "CallTracking.ru", `calltracking\.ru`)
	c.add("ct", "Callibri", `cdn\.callibri\.ru`)
	c.add("ct", "CoMagic", `app\.comagic\.ru`)
	c.add("ct", "Calltouch", `calltouch\.ru`)
	c.add("ct", "Mango-Office", `widgets\.mango-office\.ru`)
	c.add("ct", "Roistat", `roistat\.com`)
}
