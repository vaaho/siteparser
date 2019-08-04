package core

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

type SiteStatus int

const (
	StatusNone    SiteStatus = iota // не скачен
	StatusSuccess                   // успешно скачен
	StatusFail                      // была ошибка при скачивании
)

type Site struct {
	Domain  string     `json:"domain"`
	Status  SiteStatus `json:"status"`
	Url     string     `json:"url"`
	Content string     `json:"content"`
	Error   SiteError  `json:"error"`
}

type SiteError struct {
	Code    int    `json:"code"` // http статус или иное
	Message string `json:"message"`
}

type SiteStorage struct {
	dir string // папка для файлов с сайтами
}

func NewSite(domain string) Site {
	return Site{domain, StatusNone, "", "", SiteError{}}
}

func NewSuccessSite(domain, url, content string) Site {
	return Site{domain, StatusSuccess, url, content, SiteError{}}
}

func NewFailSite(domain, url string, error SiteError) Site {
	return Site{domain, StatusFail, url, "", error}
}

func NewSiteStorage(dir string) *SiteStorage {
	s := &SiteStorage{dir}
	s.ensureDir()
	return s
}

func NewSiteError(code int, message string) SiteError {
	return SiteError{code, message}
}

func (e SiteError) Error() string {
	return e.Message
}

func (s *SiteStorage) siteFilePath(domain string) string {
	if len(domain) == 0 {
		log.Fatal("Empty domain name")
	}
	return s.dir + domain + ".json"
}

func (s *SiteStorage) ensureDir() {
	_, err := os.Stat(s.dir)
	if os.IsNotExist(err) {
		err := os.Mkdir(s.dir, os.ModePerm)
		FailOnError(err)
	}
}

func (s *SiteStorage) Load(domain string) Site {
	path := s.siteFilePath(domain)

	// проверяем есть ли сайт в хранилище, если нет, то возращаем сайт статусе "Не скачен"
	fileInfo, err := os.Stat(path)
	if os.IsNotExist(err) || fileInfo.IsDir() {
		return NewSite(domain)
	}

	file, err := os.Open(path)
	FailOnError(err)
	defer file.Close()

	bytes, err := ioutil.ReadAll(file)
	FailOnError(err)

	var site Site
	err = json.Unmarshal(bytes, &site)
	FailOnError(err)

	return site
}

func (s *SiteStorage) Save(site Site) {
	path := s.siteFilePath(site.Domain)

	bytes, err := json.MarshalIndent(site, "", "  ")
	FailOnError(err)

	err = ioutil.WriteFile(path, bytes, os.ModePerm)
	FailOnError(err)
}
