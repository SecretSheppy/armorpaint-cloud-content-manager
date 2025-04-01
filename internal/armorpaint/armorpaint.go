package armorpaint

import (
	"github.com/SecretSheppy/armorpaint-cloud-content-manager/pkg/jsonutils"
	"runtime"
)

type ConfigFiles struct {
	Linux  string
	Win32  string
	Darwin string
}

func getConfigs() ConfigFiles {
	return ConfigFiles{
		Linux:  "~/.local/share/ArmorPaint/config.json",
		Win32:  "", // TODO: unknown
		Darwin: "", // TODO: unknown
	}
}

var Configs = getConfigs()

func CreateBrowserShortcut(bookmarkPath string) error {
	switch runtime.GOOS {
	case "darwin":
		err := createBrowserShortcut(bookmarkPath, Configs.Darwin)
		if err != nil {
			return err
		}
	case "windows":
		err := createBrowserShortcut(bookmarkPath, Configs.Win32)
		if err != nil {
			return err
		}
	default:
		err := createBrowserShortcut(bookmarkPath, Configs.Linux)
		if err != nil {
			return err
		}
	}

	return nil
}

func createBrowserShortcut(bookmarkPath, path string) error {
	var config Config

	err := jsonutils.Load(config, path)
	if err != nil {
		return err
	}

	config.Bookmarks = append(config.Bookmarks, bookmarkPath)

	err = jsonutils.Save(config, path)
	if err != nil {
		return err
	}

	return nil
}
