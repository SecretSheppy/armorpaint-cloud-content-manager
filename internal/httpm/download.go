package httpm

import (
	"errors"
	"io"
	"net/http"
)

var (
	ErrInvalidURL         = errors.New("invalid URL")
	ErrFileDownloadFailed = errors.New("file download failed")
)

func Download(URL string) ([]byte, error) {
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
