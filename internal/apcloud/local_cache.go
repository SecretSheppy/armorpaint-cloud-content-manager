package apcloud

type LocalCache struct {
	Root      string
	Materials string
	HDRI      string
	AssetList string
}

func NewLocalCache(root string) *LocalCache {
	return &LocalCache{
		Root:      root,
		Materials: "materials",
		HDRI:      "hdri",
		AssetList: "asset_list.json",
	}
}
