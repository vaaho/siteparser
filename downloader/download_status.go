package downloader

type DownloadStatus struct {
	Total      int // всего сайтов для скачивания
	Missed     int // пропущенных сайтов, которые уже скачены в хранилище
	Processed  int // число обработанных сайтов
	Success    int // число успешно скаченных сайтов
	Fail       int // число сайтов с ошибками скачивания
	InProgress int // число сайтов на текущий момент в процессе скачивания
}

func NewDownloadStatus() *DownloadStatus {
	return &DownloadStatus{}
}
