package listbucket

import "net/url"

func GetMarkerURL(URL, NextMarker string) (string, error) {
	u, err := url.Parse(URL)
	if err != nil {
		return "", err
	}

	params := url.Values{}
	params.Set("marker", NextMarker)

	u.RawQuery = params.Encode()

	return u.String(), nil
}
