package system

import (
	"fmt"
	"github.com/SecretSheppy/armorpaint-cloud-content-manager/internal/apcloud"
	"path/filepath"
)

func UpdateAssetCache(path string) {
	log.Info(fmt.Sprintf("attempting to update all ArmorPaint cloud resources in %s",
		filepath.Join(path, "apccm")))

	cache := getValidLocalCache(path)
	installed := getLocalAssetList(cache)
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
