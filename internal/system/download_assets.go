package system

import (
	"errors"
	"fmt"
	"github.com/SecretSheppy/armorpaint-cloud-content-manager/internal/apcloud"
	"github.com/SecretSheppy/armorpaint-cloud-content-manager/internal/armorpaint"
	"github.com/SecretSheppy/armorpaint-cloud-content-manager/internal/logger"
	"github.com/SecretSheppy/armorpaint-cloud-content-manager/pkg/files"
	"net/url"
	"os"
	"path/filepath"
)

var log = logger.Get()

func DownloadAllAssets(path string) {
	log.Info(fmt.Sprintf("attempting to install all ArmorPaint cloud resources into %s",
		filepath.Join(path, "apccm")))

	cache := apcloud.NewLocalCache(path)
	makePath(cache.Root)
	makePath(cache.Materials)

	log.Info("acquired local cache")

	assets, err := apcloud.GetAssets()
	if err != nil {
		log.Panic("failed to get assets list")
		panic(err)
	} else {
		log.Info(fmt.Sprintf("acquired %d assets from %s", len(assets.Assets), apcloud.BaseURL))
	}

	for i, asset := range assets.Assets {
		if files.GetPathState(asset.URL) != files.File {
			log.Info(fmt.Sprintf("download progress (%d/%d) > skipping %s", i, len(assets.Assets), asset.URL))
			continue
		}

		completeURL, err := url.JoinPath(apcloud.BaseURL, asset.URL)
		if err != nil {
			log.Panic("failed to join url")
			panic(err)
		}

		filename := filepath.Base(asset.URL)
		savePath := filepath.Join(cache.Materials, filename)

		err = apcloud.DownloadAsset(completeURL, savePath)
		if err != nil {
			log.Panic(fmt.Sprintf("failed to download asset %s: %s", completeURL, err.Error()))
			panic(err)
		}

		log.Info(fmt.Sprintf("download progress (%d/%d) > downloaded %s", i, len(assets.Assets), asset.URL))
	}

	err = apcloud.SaveAssetList(assets, cache.AssetList)
	if err != nil {
		log.Warn("failed to save asset list, updating will not work")
	} else {
		log.Info("saved asset list (.asset_list.json)")
	}

	err = armorpaint.CreateBrowserShortcut(cache.Root)
	if err != nil {
		log.Warn("failed to create browser shortcut")
	} else {
		log.Info("created browser shortcut")
	}

	log.Info("cloud content installed successfully, restart ArmorPainter for quick access")
}

func makePath(path string) {
	err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
		if errors.Is(err, os.ErrExist) {
			log.Warn(fmt.Sprintf("path %s already exists, continuing...", path))
		} else {
			log.Panic(fmt.Sprintf("failed to create path %s", path))
		}
	}
	log.Info(fmt.Sprintf("acquired path %s", path))
}

func RemoveAllAssets(path string) {
	// TODO: use the .assets_list.json to only remove the assets and then remove
	// 	the .assets_list.json
}

func UpdateAssetCache(path string) {
	// cache := apcloud.NewLocalCache(path)

	// assets, err := apcloud.LoadAssetList(cache.AssetList)
}
