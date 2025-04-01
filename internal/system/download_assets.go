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
	"sync"
)

const numWorkers = 14

var log = logger.Get()

func DownloadAllAssets(path string) {
	log.Info(fmt.Sprintf("attempting to install all ArmorPaint cloud resources into %s",
		filepath.Join(path, "apccm")))

	cache := apcloud.NewLocalCache(path)
	log.Info("acquired local cache")

	makePath(cache.Root)
	makePath(cache.Materials)

	assets, err := apcloud.GetAssets()
	if err != nil {
		log.Panic("failed to get assets list")
		panic(err)
	}
	log.Info(fmt.Sprintf("acquired %d assets from %s", len(assets.Assets), apcloud.BaseURL))

	jobs := make(chan DownloadJob)

	var wg sync.WaitGroup

	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go downloadWorker(i, jobs, &wg)
	}

	go func() {
		for _, asset := range assets.Assets {
			if files.GetPathState(asset.URL) == files.File {
				jobs <- NewDownloadJob(asset, *cache)
			}
		}

		close(jobs)
	}()

	wg.Wait()

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

type DownloadJob struct {
	Asset apcloud.Asset
	Cache apcloud.LocalCache
}

func NewDownloadJob(asset apcloud.Asset, cache apcloud.LocalCache) DownloadJob {
	return DownloadJob{
		Asset: asset,
		Cache: cache,
	}
}

func downloadWorker(ID int, jobs <-chan DownloadJob, wg *sync.WaitGroup) {
	defer wg.Done()

	for job := range jobs {
		URL, err := url.JoinPath(apcloud.BaseURL, job.Asset.URL)
		if err != nil {
			log.Panic(fmt.Sprintf("worker %d :: failed to join url", ID))
			panic(err)
		}

		filename := filepath.Base(job.Asset.URL)
		path := filepath.Join(job.Cache.Materials, filename)

		err = apcloud.DownloadAsset(URL, path)
		if err != nil {
			log.Panic(fmt.Sprintf("worker %d :: failed to download asset %s", ID, URL))
			panic(err)
		}

		log.Info(fmt.Sprintf("worker %d :: downloaded > %s", ID, URL))
	}
}

func RemoveAllAssets(path string) {
	// TODO: use the .assets_list.json to only remove the assets and then remove
	// 	the .assets_list.json
}

func UpdateAssetCache(path string) {
	// cache := apcloud.NewLocalCache(path)

	// assets, err := apcloud.LoadAssetList(cache.AssetList)
}
