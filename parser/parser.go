package parser

import (
	"log"
	"siteparser/core"
)

// Фильтрует канал с доменами на предмет уже скаченных файлов
func ParseSites(sites <-chan core.Site) <-chan core.Site {
	out := make(chan core.Site)

	go func() {
		for site := range sites {
			if site.Status == core.StatusSuccess {
				result := ParseSite(site)
				site.ParseData = result
				out <- site
			}
		}
		close(out)
	}()

	return out
}

func ParseSite(site core.Site) *core.ParseData {
	result := core.NewParseData()

	result.Numbers = parseNumbers(site.Content)
	result.Categories = parseCategories(site.Content)

	return &result
}

func CollectStatus(sites <-chan core.Site, status *ParseStatus) <-chan core.Site {
	out := make(chan core.Site)

	go func() {
		for site := range sites {
			if site.ParseData != nil {
				status.Parsed++
				if len(site.ParseData.Numbers) > 0 {
					status.NumbersCount++
					//log.Printf("[NUMBERS] [%s] %s", site.Domain, strings.Join(site.ParseData.Numbers, "; "))
				}
				if len(site.ParseData.Categories) > 0 {
					status.CategoriesCount++
					log.Printf("[CATEGORIES] [%s] %+v", site.Domain, site.ParseData.Categories)
				}
				out <- site
			}
		}
		close(out)
	}()

	return out
}

func WaitAndLog(sites <-chan core.Site, status *ParseStatus) {
	// главный цикл ожидания, заканчивается только когда все сайты будут обработаны
	for _ = range sites {
	}

	log.Printf("[PARSED] %+v", *status)
}

func Parse(config *core.Config, storage *core.SiteStorage) ParseStatus {
	status := NewParseStatus()

	sites, total := core.LoadSites(storage)
	status.Total = total
	sites = ParseSites(sites)
	sites = core.SaveSites(sites, storage)
	sites = CollectStatus(sites, status)

	WaitAndLog(sites, status)

	return *status
}
