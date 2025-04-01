package apcloud

import (
	"encoding/xml"
	"github.com/SecretSheppy/armorpaint-cloud-content-manager/internal/httpm"
)

const cloudSourceURL = "https://armorpaint.fra1.digitaloceanspaces.com/"

type Owner struct {
	XMLName     xml.Name `xml:"Owner"`
	ID          int      `xml:"ID"`
	DisplayName string   `xml:"DisplayName"`
}

type Contents struct {
	XMLName      xml.Name `xml:"Contents"`
	Key          string   `xml:"Key"`
	LastModified string   `xml:"LastModified"`
	ETag         string   `xml:"ETag"`
	Size         int      `xml:"Size"`
	StorageClass string   `xml:"StorageClass"`
	Owner        Owner    `xml:"Owner"`
	Type         string   `xml:"Type"`
}

type ListBucketResult struct {
	XMLName     xml.Name   `xml:"ListBucketResult"`
	Name        string     `xml:"Name"`
	MaxKeys     int        `xml:"MaxKeys"`
	IsTruncated bool       `xml:"IsTruncated"`
	Contents    []Contents `xml:"Contents"`
	Marker      string     `xml:"Marker"`
	NextMarker  string     `xml:"NextMarker"`
}

func NewListBucketResult(raw []byte) (*ListBucketResult, error) {
	var list ListBucketResult

	err := xml.Unmarshal(raw, &list)
	if err != nil {
		return &ListBucketResult{}, err
	}

	return &list, nil
}

func ProbeOnce(URL string) (*ListBucketResult, error) {
	raw, err := httpm.Download(URL)
	if err != nil {
		return &ListBucketResult{}, err
	}

	return NewListBucketResult(raw)
}
