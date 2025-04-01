package apcloud

import (
	"testing"
)

const validXMLFormat = `
<ListBucketResult>
	<Name>armorpaint</Name>
	<Prefix/>
	<MaxKeys>1000</MaxKeys>
	<IsTruncated>true</IsTruncated>
	<Contents>
		<Key>cloud/</Key>
		<LastModified>2023-08-14T17:59:17.472Z</LastModified>
		<ETag>"1441a7909c087dbbe7ce59881b9df8b9"</ETag>
		<Size>15</Size>
		<StorageClass>STANDARD</StorageClass>
		<Owner>
			<ID>633352</ID>
			<DisplayName>633352</DisplayName>
		</Owner>
		<Type>Normal</Type>
	</Contents>
	<NextMarker>
		cloud/materials/ambientcg/photogrammetry/Ground052_4K-JPG_icon.jpg
	</NextMarker>
</ListBucketResult>`

func TestNewListBucketResult(t *testing.T) {
	list, err := NewListBucketResult([]byte(validXMLFormat))
	if err != nil {
		t.Fatal(err)
	}

	if list.Name != "armorpaint" {
		t.Errorf("NewListBucketResult: got %s, want %s", list.Name, "armorpaint")
	}

	if list.Contents[0].Key != "cloud/" {
		t.Errorf("NewListBucketResult: got %s, want %s", list.Contents[0].Key, "cloud/")
	}

	if list.Contents[0].Owner.ID != 633352 {
		t.Errorf("NewListBucketResult: got %d, want %s", list.Contents[0].Owner.ID, "633352")
	}
}
