package downloader

import (
	"bufio"
	"log"
	"os"
	"siteparser/core"
	"time"
)

// Загружает список доменов из файла
func LoadDomains(sourceFile string, hasColumnsRow bool) <-chan string {
	file, err := os.Open(sourceFile)
	core.FailOnError(err)
	defer file.Close()

	// используем map как set, чтобы сохранять только уникальные домены
	domains := make(map[string]bool)

	scanner := bufio.NewScanner(file)
	if hasColumnsRow {
		scanner.Scan()
	}

	for scanner.Scan() {
		if domain := scanner.Text(); domain != "" {
			domains[domain] = true
		}
	}

	err = scanner.Err()
	core.FailOnError(err)

	// превращаем set в канал
	out := make(chan string, len(domains))
	for domain := range domains {
		out <- domain
	}

	close(out)
	return out
}

// Фильтрует канал с доменами на предмет уже скаченных файлов
func FilterDownloadedDomains(domains <-chan string, storage *core.SiteStorage) <-chan string {
	out := make(chan string, len(domains))

	for domain := range domains {
		site := storage.Load(domain)
		// статус none означает, что сайт в хранилище не найден, т.е. ещё раз ниразу не скачивался
		if site.Status == core.StatusNone {
			out <- domain
		}
	}

	close(out)
	return out
}

// Скачивает сайты в параллельном режиме
// Возвращает канал со скаченными сайтами и ещё один канал с количеством сайтов в обработке
func ParallelDownloadSites(domains <-chan string, maxTreads int) (<-chan core.Site, <-chan bool) {
	inProgress := make(chan bool, maxTreads) // пул потоков для паралельного скачивания
	out := make(chan core.Site)

	go func() {
		for domain := range domains {
			inProgress <- true // ждёт, если запущено более maxTreads потоков

			go func(domain string) {
				site := downloadSite(domain)

				<-inProgress // освобождаем поток
				out <- site

				// если каналы доменов и какнал прогресса закрылся, то выходим
				if len(inProgress) == 0 && len(domains) == 0 {
					close(out)
				}
			}(domain)
		}

		// проверка на выход, если канал доменов был изначально пустым
		if len(inProgress) == 0 && len(domains) == 0 {
			close(out)
		}
	}()

	return out, inProgress
}

func CollectStatus(sites <-chan core.Site, status *DownloadStatus, logFails bool) <-chan core.Site {
	out := make(chan core.Site)

	go func() {
		for site := range sites {
			status.Processed++
			if site.Status == core.StatusSuccess {
				status.Success++
			}
			if site.Status == core.StatusFail {
				status.Fail++
				if logFails {
					log.Printf("[FAIL] [%s] %s", site.Domain, site.Error.Message)
				}
			}
			out <- site
		}
		close(out)
	}()

	return out
}

func WaitAndLog(sites <-chan core.Site, status *DownloadStatus, inProgress <-chan bool) {
	// поток который будет раз в 10 сек выводить статус
	go func() {
		tick := time.Tick(5 * time.Second)
		for {
			<-tick // ждём следующего тика

			// обновляем число сайта в обработке
			status.InProgress = len(inProgress)

			log.Printf("[STATUS] %+v", *status)
		}
	}()

	// главный цикл ожидания, заканчивается только когда все сайты будут обработаны
	for _ = range sites {
	}

	status.InProgress = len(inProgress)

	log.Printf("[FINISHED] %+v", *status)
}

func Download(config *core.Config, storage *core.SiteStorage) DownloadStatus {
	status := NewDownloadStatus()

	domains := LoadDomains(config.InputFile, !config.NoColumnsRow)
	status.Total = len(domains)
	domains = FilterDownloadedDomains(domains, storage)
	status.Missed = status.Total - len(domains)
	sites, inProgress := ParallelDownloadSites(domains, config.MaxThreads)
	sites = core.SaveSites(sites, storage)
	sites = CollectStatus(sites, status, config.LogFails)

	// ждём окончания процесса скачки сайтов и выводим промежуточный статус
	WaitAndLog(sites, status, inProgress)

	return *status
}
