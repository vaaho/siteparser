package downloader

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"siteparser/core"
)

func downloadSite(domain string) core.Site {
	url := "http://" + domain
	content, siteError := downloadSiteContent(url)

	// если протокол http не поддерживается (была сетевая ошибка), то делаем вторую попытку с другим протоколом
	if siteError.Code < 0 {
		url = "https://" + domain
		content, siteError = downloadSiteContent(url)
	}

	if content == "" {
		return core.NewFailSite(domain, url, siteError)
	}

	return core.NewSuccessSite(domain, url, content)
}

func downloadSiteContent(url string) (string, core.SiteError) {
	resp, err := http.Get(url)
	if err != nil {
		return "", core.NewSiteError(-1, fmt.Sprintf("GET error: %v", err))
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", core.NewSiteError(resp.StatusCode, fmt.Sprintf("Status error: %v", resp.StatusCode))
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", core.NewSiteError(-1, fmt.Sprintf("Read body: %v", err))
	}

	return string(data), core.NewSiteError(0, "")
}
