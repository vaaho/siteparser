package core

// Создат поток сайтов из хранилища. Дополнительно возвращет общее количество сайтов в хранилище
func LoadSites(storage *SiteStorage) (<-chan Site, int) {
	out := make(chan Site)

	domains := storage.GetDomains()

	go func() {
		for _, domain := range domains {
			site := storage.Load(domain)
			out <- site
		}
		close(out)
	}()

	return out, len(domains)
}

// Сохраняет поток сайтов в хранилище
func SaveSites(sites <-chan Site, storage *SiteStorage) <-chan Site {
	out := make(chan Site)

	go func() {
		for site := range sites {
			storage.Save(site)
			out <- site
		}
		close(out)
	}()

	return out
}
