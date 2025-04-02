package system

import (
	"errors"
	"fmt"
	"github.com/SecretSheppy/armorpaint-cloud-content-manager/internal/apcloud"
	"github.com/SecretSheppy/armorpaint-cloud-content-manager/internal/logger"
	"os"
)

var log = logger.Get()

func directoryExists(path string) (bool, error) {
	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false, nil
	}

	if err != nil {
		return false, err
	}

	return info.IsDir(), nil
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

func getOnlineAssetsList() *apcloud.AssetList {
	assets, err := apcloud.GetAssets()
	if err != nil {
		log.Panic("failed to get assets list")
		panic(err)
	}

	log.Info(fmt.Sprintf("acquired %d assets from %s", len(assets.Assets), apcloud.BaseURL))
	return assets
}

func saveLocalAssetList(assets *apcloud.AssetList, cache *apcloud.LocalCache) {
	err := apcloud.SaveAssetList(assets, cache.AssetList)
	if err != nil {
		log.Warn("failed to save asset list, updating will not work")
	} else {
		log.Info("saved asset list (.asset_list.json)")
	}
}

func getLocalAssetList(cache *apcloud.LocalCache) *apcloud.AssetList {
	installed, err := apcloud.LoadAssetList(cache.AssetList)
	if err != nil {
		log.Panic(fmt.Sprintf("failed to load local asset list %s", cache.AssetList))
		panic(err)
	}
	log.Info(fmt.Sprintf("acquired local asset list, %d assets currently installed*", len(installed.Assets)))
	log.Warn(fmt.Sprintf("    *local asset list includes skipped items, count inaccurate"))
	return installed
}

func getValidLocalCache(path string) *apcloud.LocalCache {
	cache := apcloud.NewLocalCache(path)
	for _, p := range []string{cache.Root, cache.Materials} {
		if exists, err := directoryExists(p); err != nil || !exists {
			log.Panic(fmt.Sprintf("directory %s does not exist", p))
			panic(err)
		}

		log.Info(fmt.Sprintf("aquired directory %s", p))
	}
	log.Info("acquired local cache")
	return cache
}
