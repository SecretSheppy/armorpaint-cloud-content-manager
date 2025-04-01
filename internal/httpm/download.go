package httpm

import (
	"errors"
	"io"
	"net/http"
	"os"
)

var (
	ErrInvalidURL         = errors.New("invalid URL")
	ErrFileDownloadFailed = errors.New("file download failed")
	ErrCreateFileFailed   = errors.New("file create failed")
	ErrSaveFileFailed     = errors.New("file save failed")
)

func DownloadToCache(URL string) ([]byte, error) {
	resp, err := http.Get(URL)
	if err != nil {
		return nil, ErrInvalidURL
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, ErrFileDownloadFailed
	}

	URLContent, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return URLContent, nil
}

func DownloadToFile(URL, path string) error {
	out, err := os.Create(path)
	if err != nil {
		return ErrCreateFileFailed
	}
	defer out.Close()

	resp, err := http.Get(URL)
	if err != nil {
		return ErrFileDownloadFailed
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return ErrFileDownloadFailed
	}

	if _, err := io.Copy(out, resp.Body); err != nil {
		return ErrSaveFileFailed
	}

	return nil
}
