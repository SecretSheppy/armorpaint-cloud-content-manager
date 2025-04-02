package system

import (
	"fmt"
	"os"
	"path/filepath"
)

func RemoveAllAssets(path string) {
	log.Info(fmt.Sprintf("attempting to remove all ArmorPaint cloud resources from %s",
		filepath.Join(path, "apccm")))

	cache := getValidLocalCache(path)
	installed := getLocalAssetList(cache)

	failedCount := 0
	for i, a := range installed.Assets {
		filename := filepath.Base(a.URL)
		file := filepath.Join(cache.Materials, filename)

		err := os.Remove(file)
		if err != nil {
			log.Error(fmt.Sprintf("failed to remove asset %s: %s", file, err.Error()))
			failedCount++
			continue
		}

		log.Info(fmt.Sprintf("progress (%d/%d) :: removed > %s", i+1, len(installed.Assets), file))
	}

	log.Info(fmt.Sprintf("completed :: failed to remove %d assets", failedCount))

	err := os.Remove(cache.AssetList)
	if err != nil {
		log.Error(fmt.Sprintf("failed to remove asset list"))
	} else {
		log.Info(fmt.Sprintf("completed :: removed > %s", cache.AssetList))
	}

	log.Info("ArmorPaint cloud content removal process completed")
}
