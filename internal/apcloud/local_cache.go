package apcloud

import "path/filepath"

type LocalCache struct {
	Root      string
	Materials string
	HDRI      string
	AssetList string
}

func NewLocalCache(root string) *LocalCache {
	root = filepath.Join(root, "apccm/")
	return &LocalCache{
		Root:      root,
		Materials: filepath.Join(root, "materials/"),
		HDRI:      filepath.Join(root, "hdri/"),
		AssetList: filepath.Join(root, "asset_list.json"),
	}
}
