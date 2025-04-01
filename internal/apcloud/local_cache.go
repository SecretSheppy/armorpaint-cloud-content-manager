package apcloud

import "path/filepath"

type LocalCache struct {
	Root      string
	Materials string
	AssetList string
}

func NewLocalCache(root string) *LocalCache {
	root = filepath.Join(root, "apccm/")
	return &LocalCache{
		Root:      root,
		Materials: filepath.Join(root, "materials/"),
		AssetList: filepath.Join(root, "asset_list.json"),
	}
}
