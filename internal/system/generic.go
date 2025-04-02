package system

import (
	"fmt"
	"github.com/SecretSheppy/armorpaint-cloud-content-manager/internal/apcloud"
	"github.com/SecretSheppy/armorpaint-cloud-content-manager/internal/logger"
)

var log = logger.Get()

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
