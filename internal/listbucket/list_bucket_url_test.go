package listbucket

import "testing"

type ListBucketURLTestCase struct {
	testName string
	url      string
	marker   string
	expected string
}

var listBucketURLTestCase = []ListBucketURLTestCase{
	{
		testName: "Basic",
		url:      "https://example.com/?marker=xyz",
		marker:   "hello world",
		expected: "https://example.com/?marker=hello+world",
	},
	{
		testName: "Url with no default marker variable",
		url:      "https://example.com/",
		marker:   "hello world",
		expected: "https://example.com/?marker=hello+world",
	},
}

func assertListBucketURL(t *testing.T, testCase *ListBucketURLTestCase) {
	u, err := GetMarkerURL(testCase.url, testCase.marker)
	if err != nil {
		t.Errorf("Test Failed: error while getting marker URL: %v", err)
	}

	if u != testCase.expected {
		t.Errorf("Test Failed: got: %v expected: %v", u, testCase.expected)
	}
}

func TestListBucketURL(t *testing.T) {
	for _, c := range listBucketURLTestCase {
		t.Run(c.testName, func(t *testing.T) {
			assertListBucketURL(t, &c)
		})
	}
}
