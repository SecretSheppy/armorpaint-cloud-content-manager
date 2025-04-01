package httpm

import (
	"errors"
	"testing"
)

const (
	testURL        = "https://raw.githubusercontent.com/SecretSheppy/armorpaint-cloud-content-manager/refs/heads/main/download_test.txt"
	testURLContent = "download me :)"
)

func TestDownload(t *testing.T) {
	resp, err := Download(testURL)
	if err != nil {
		t.Fatal(err)
	}

	if string(resp) != testURLContent {
		t.Errorf("got %q, want %q", resp, testURLContent)
	}
}

func TestDownloadInvalidURL(t *testing.T) {
	_, err := Download("I am not a url")
	if err != nil {
		if errors.Is(err, ErrInvalidURL) {
			return
		}
	}

	t.Errorf("expected error %s, got none", ErrInvalidURL)
}
