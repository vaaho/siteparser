package core

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// Хранилище сайтов. Каждый сайт сохраняется в отдельный json файл в общей дириктории
type SiteStorage struct {
	dir string // папка для файлов с сайтами
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

func (s *SiteStorage) domainFromFileName(filename string) string {
	ext := filepath.Ext(filename)
	if ext != ".json" {
		return ""
	}
	return strings.TrimSuffix(filename, ext)
}

func (s *SiteStorage) ensureDir() {
	_, err := os.Stat(s.dir)
	if os.IsNotExist(err) {
		err := os.Mkdir(s.dir, os.ModePerm)
		FailOnError(err)
	}
}

// Загружает сайт из хранилища по его домену. Если сайта в хранилище нет, то возвращает пустой сайт
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

// Сохраняет сайт в хранилище
func (s *SiteStorage) Save(site Site) {
	path := s.siteFilePath(site.Domain)

	bytes, err := json.MarshalIndent(site, "", "  ")
	FailOnError(err)

	err = ioutil.WriteFile(path, bytes, os.ModePerm)
	FailOnError(err)
}

// Возрващает список доменов всех сайтов в хранилище
func (s *SiteStorage) GetDomains() []string {
	files, err := ioutil.ReadDir(s.dir)
	FailOnError(err)

	domains := make([]string, 0, len(files))
	for _, f := range files {
		if domain := s.domainFromFileName(f.Name()); domain != "" {
			domains = append(domains, domain)
		}
	}

	return domains
}
