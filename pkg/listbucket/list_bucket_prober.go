package listbucket

import "github.com/SecretSheppy/armorpaint-cloud-content-manager/internal/httpm"

func ProbeOnce(URL string) (*ListBucketResult, error) {
	raw, err := httpm.DownloadToCache(URL)
	if err != nil {
		return &ListBucketResult{}, err
	}

	return NewListBucketResult(raw)
}

func ProbeAll(URL string) (*ListBucketResult, error) {
	result, err := ProbeOnce(URL)
	if err != nil {
		return &ListBucketResult{}, err
	}

	if !result.IsTruncated {
		return result, nil
	}

	markerURL, err := GetMarkerURL(URL, result.NextMarker)
	if err != nil {
		return &ListBucketResult{}, err
	}

	allResults, err := ProbeAll(markerURL)
	if err != nil {
		return &ListBucketResult{}, err
	}

	allResults.Contents = append(result.Contents, allResults.Contents...)
	return allResults, nil
}
