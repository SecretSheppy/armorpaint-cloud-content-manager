package system

import (
	"fmt"
	"github.com/SecretSheppy/armorpaint-cloud-content-manager/internal/apcloud"
	"os"
	"path/filepath"
)

func UpdateAssetCache(path string) {
	log.Info(fmt.Sprintf("attempting to update all ArmorPaint cloud resources in %s",
		filepath.Join(path, "apccm")))

	cache := apcloud.NewLocalCache(path)
	for _, p := range []string{cache.Root, cache.Materials} {
		if exists, err := directoryExists(p); err != nil || !exists {
			log.Panic(fmt.Sprintf("directory %s does not exist", p))
			panic(err)
		}

		log.Info(fmt.Sprintf("aquired directory %s", p))
	}
	log.Info("acquired local cache")

	installed, err := apcloud.LoadAssetList(cache.AssetList)
	if err != nil {
		log.Panic(fmt.Sprintf("failed to load local asset list %s", cache.AssetList))
		panic(err)
	}
	log.Info(fmt.Sprintf("acquired local asset list, %d assets currently installed*", len(installed.Assets)))
	log.Warn(fmt.Sprintf("    *local asset list includes skipped items, count inaccurate"))

	assets := getOnlineAssetsList()

	installedM := apcloud.AssetListToMap(installed)
	var toGet apcloud.AssetList
	for i, a := range assets.Assets {
		if installedM[a.URL].LastModified != a.LastModified {
			toGet.Assets = append(toGet.Assets, a)
		}
		log.Info(fmt.Sprintf("progress (%d/%d) :: checked > %s", i+1, len(assets.Assets), a.URL))
	}

	if len(toGet.Assets) == 0 {
		log.Info("no assets to update")
		return
	}

	log.Info(fmt.Sprintf("updating %d assets", len(toGet.Assets)))
	workerPoolDownload(assets, cache)
	saveLocalAssetList(assets, cache)
}

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
