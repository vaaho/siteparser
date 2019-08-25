package core

// Статус скачивания сайта
type SiteStatus int

const (
	StatusNone    SiteStatus = iota // не скачен
	StatusSuccess                   // успешно скачен
	StatusFail                      // была ошибка при скачивании
)

// Структура сайта из хранилища сайтов. Сереализуема в json
type Site struct {
	Domain    string     `json:"domain"`
	Status    SiteStatus `json:"status"`
	Url       string     `json:"url"`
	Content   string     `json:"content"`
	Error     SiteError  `json:"error"`
	ParseData *ParseData `json:"parseData,omitempty"`
}

type SiteError struct {
	Code    int    `json:"code"` // http статус или иное
	Message string `json:"message"`
}

func NewSite(domain string) Site {
	return Site{domain, StatusNone, "", "", SiteError{}, nil}
}

func NewSuccessSite(domain, url, content string) Site {
	return Site{domain, StatusSuccess, url, content, SiteError{}, nil}
}

func NewFailSite(domain, url string, error SiteError) Site {
	return Site{domain, StatusFail, url, "", error, nil}
}
