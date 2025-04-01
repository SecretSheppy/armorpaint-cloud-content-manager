package apcloud

import (
	"github.com/SecretSheppy/armorpaint-cloud-content-manager/internal/data"
	"github.com/SecretSheppy/armorpaint-cloud-content-manager/pkg/jsonutils"
	"github.com/SecretSheppy/armorpaint-cloud-content-manager/pkg/listbucket"
)

type Asset struct {
	URL          string `json:"url"`
	LastModified string `json:"last_modified"`
}

func NewAsset(url, lastModified string) *Asset {
	return &Asset{
		URL:          url,
		LastModified: lastModified,
	}
}

type AssetList struct {
	ApccmVersion string  `json:"apccm_version"`
	Assets       []Asset `json:"assets"`
}

func NewAssetList(assets []Asset) *AssetList {
	return &AssetList{
		ApccmVersion: data.AccpmVersion,
		Assets:       assets,
	}
}

func GetAssets() (*AssetList, error) {
	result, err := listbucket.ProbeAll(BaseURL)
	if err != nil {
		return nil, err
	}

	assets := make([]Asset, len(result.Contents))
	for i, r := range result.Contents {
		assets[i] = *NewAsset(r.Key, r.LastModified)
	}

	return NewAssetList(assets), nil
}

func SaveAssetList(assets *AssetList, path string) error {
	err := jsonutils.Save(assets, path)
	if err != nil {
		return err
	}

	return nil
}

func LoadAssetList(path string) (*AssetList, error) {
	var assets AssetList

	err := jsonutils.Load(&assets, path)
	if err != nil {
		return nil, err
	}

	return &assets, nil
}
