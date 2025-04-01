package httpm

import (
	"errors"
	"testing"
)

const (
	testURL        = "https://raw.githubusercontent.com/SecretSheppy/armorpaint-cloud-content-manager/refs/heads/main/download_test.txt"
	testURLContent = "download me :)"
)

func TestDownloadToCache(t *testing.T) {
	resp, err := DownloadToCache(testURL)
	if err != nil {
		t.Fatal(err)
	}

	if string(resp) != testURLContent {
		t.Errorf("got %q, want %q", resp, testURLContent)
	}
}

func TestDownloadToCacheInvalidURL(t *testing.T) {
	_, err := DownloadToCache("I am not a url")
	if err != nil {
		if errors.Is(err, ErrInvalidURL) {
			return
		}
	}

	t.Errorf("expected error %s, got none", ErrInvalidURL)
}
