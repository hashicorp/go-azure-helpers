package storage

import (
	"net/url"
	"strings"
	"testing"

	"github.com/Azure/go-autorest/autorest/azure"
)

func TestParseStorageAccountConnectionString(t *testing.T) {
	testCases := []struct {
		input               string
		expectedAccountName string
		expectedAccountKey  string
		expectedError       bool
	}{
		{
			"DefaultEndpointsProtocol=https;AccountName=azurermtestsa0;AccountKey=2vJrjEyL4re2nxCEg590wJUUC7PiqqrDHjAN5RU304FNUQieiEwS2bfp83O0v28iSfWjvYhkGmjYQAdd9x+6nw==;EndpointSuffix=core.windows.net",
			"azurermtestsa0",
			"2vJrjEyL4re2nxCEg590wJUUC7PiqqrDHjAN5RU304FNUQieiEwS2bfp83O0v28iSfWjvYhkGmjYQAdd9x+6nw==",
			false,
		},
		{
			"DefaultEndpointsProtocol=https;AccountName=azurermtestsa0;AccountKey=2vJrjEyL4re2nxCEg590wJUUC7PiqqrDHjAN5RU304FNUQieiEwS2bfp83O0v28iSfWjvYhkGmjYQAdd9x+6nw==;EndpointSuffix",
			"",
			"",
			true,
		},
	}

	for _, test := range testCases {
		result, err := ParseAccountSASConnectionString(test.input)

		if test.expectedError {
			if err == nil {
				t.Fatalf("Expected error for %s: %q", test.input, err)
			}
			return
		}

		if !test.expectedError && err != nil {
			t.Fatalf("Failed to parse resource type string: %s, %q", test.input, result)
		}

		if val, pres := result[connStringAccountKeyKey]; !pres || val != test.expectedAccountKey {
			t.Fatalf("Failed to parse Account Key: Expected: %s, Found: %s", test.expectedAccountKey, val)
		}
		if val, pres := result[connStringAccountNameKey]; !pres || val != test.expectedAccountName {
			t.Fatalf("Failed to parse Account Name: Expected: %s, Found: %s", test.expectedAccountName, val)
		}
	}
}

// This connection string was for a real storage account which has been deleted
// so its safe to include here for reference to understand the format.
// DefaultEndpointsProtocol=https;AccountName=azurermtestsa0;AccountKey=T0ZQouXBDpWud/PlTRHIJH2+VUK8D+fnedEynb9Mx638IYnsMUe4mv1fFjC7t0NayTfFAQJzPZuV1WHFKOzGdg==;EndpointSuffix=core.windows.net
func TestComputeAccountSASToken(t *testing.T) {
	testCases := []struct {
		accountName    string
		accountKey     string
		permissions    string
		services       string
		resourceTypes  string
		start          string
		expiry         string
		signedProtocol string
		signedIp       string
		signedVersion  string
		knownSasToken  string
	}{
		{
			"azurermtestsa0",
			"T0ZQouXBDpWud/PlTRHIJH2+VUK8D+fnedEynb9Mx638IYnsMUe4mv1fFjC7t0NayTfFAQJzPZuV1WHFKOzGdg==",
			"rwac",
			"b",
			"c",
			"2018-03-20T04:00:00Z",
			"2020-03-20T04:00:00Z",
			"https",
			"",
			"2017-07-29",
			"?sv=2017-07-29&ss=b&srt=c&sp=rwac&se=2020-03-20T04:00:00Z&st=2018-03-20T04:00:00Z&spr=https&sig=SQigK%2FnFA4pv0F0oMLqr6DxUWV4vtFqWi6q3Mf7o9nY%3D",
		},
		{
			"azurermtestsa0",
			"2vJrjEyL4re2nxCEg590wJUUC7PiqqrDHjAN5RU304FNUQieiEwS2bfp83O0v28iSfWjvYhkGmjYQAdd9x+6nw==",
			"rwdlac",
			"b",
			"sco",
			"2018-03-20T04:00:00Z",
			"2018-03-28T05:04:25Z",
			"https,http",
			"",
			"2017-07-29",
			"?sv=2017-07-29&ss=b&srt=sco&sp=rwdlac&se=2018-03-28T05:04:25Z&st=2018-03-20T04:00:00Z&spr=https,http&sig=OLNwL%2B7gxeDQQaUyNdXcDPK2aCbCMgEkJNjha9te448%3D",
		},
	}

	for _, test := range testCases {
		computedToken, err := ComputeAccountSASToken(test.accountName,
			test.accountKey,
			test.permissions,
			test.services,
			test.resourceTypes,
			test.start,
			test.expiry,
			test.signedProtocol,
			test.signedIp,
			test.signedVersion)

		if err != nil {
			t.Fatalf("Test Failed: Error computing storage account Sas: %q", err)
		}

		if computedToken != test.knownSasToken {
			t.Fatalf("Test failed: Expected Azure SAS %s but was %s", test.knownSasToken, computedToken)
		}
	}
}

func TestComputeContainerSASToken(t *testing.T) {
	testCases := []struct {
		signedPermissions  string
		signedStart        string
		signedExpiry       string
		accountName        string
		accountKey         string
		containerName      string
		signedIdentifier   string
		signedIp           string
		signedProtocol     string
		signedSnapshotTime string
		cacheControl       string
		contentDisposition string
		contentEncoding    string
		contentLanguage    string
		contentType        string
		knownSasToken      string
	}{
		{
			"rwl",
			"2019-03-27",
			"2019-09-21T09:21Z",
			"azurermblobcontainertest",
			"y3PNtHAAyjMSRHZ26n/ISyXt1IpXLIqiwAUQ602Un8AJX2JL3MMEbxK7ue45nr9BB0BibegTkQ5rdrgMR5CZkA==",
			"test-container",
			"",
			"",
			"https",
			"",
			"",
			"",
			"",
			"",
			"",
			"?st=2019-03-27&se=2019-09-21T09%3A21Z&sp=rwl&spr=https&sv=2018-11-09&sr=c&sig=DnHeyj11jpfGEmdskSIuASIZorLghcjLbKN90n%2B6UO4%3D",
		},
		{
			"rwl",
			"2019-03-27",
			"2019-09-21T09:21Z",
			"azurermblobcontainertest",
			"y3PNtHAAyjMSRHZ26n/ISyXt1IpXLIqiwAUQ602Un8AJX2JL3MMEbxK7ue45nr9BB0BibegTkQ5rdrgMR5CZkA==",
			"test-container",
			"",
			"93.23.223.54",
			"https",
			"",
			"no-cache",
			"attachment",
			"gzip",
			"en-US",
			"text/html; charset=utf-8",
			"?st=2019-03-27&se=2019-09-21T09%3A21Z&sp=rwl&sip=93.23.223.54&spr=https&sv=2018-11-09&sr=c&rscc=no-cache&rscd=attachment&rsce=gzip&rscl=en-US&rsct=text/html%3B%20charset%3Dutf-8&sig=M2TaUVEGlRVJjNt/c7Eqt2zH6%2BA8dpiLmTXR0ZevEX8%3D",
		},
	}

	for _, test := range testCases {
		computedToken, err := ComputeContainerSASToken(test.signedPermissions,
			test.signedStart,
			test.signedExpiry,
			test.accountName,
			test.accountKey,
			test.containerName,
			test.signedIdentifier,
			test.signedIp,
			test.signedProtocol,
			test.signedSnapshotTime,
			test.cacheControl,
			test.contentDisposition,
			test.contentEncoding,
			test.contentLanguage,
			test.contentType)

		if err != nil {
			t.Fatalf("Test Failed: Error computing blob container Sas: %q", err)
		}

		if !compareSASTokens(computedToken, test.knownSasToken) {
			t.Fatalf("Test failed: Expected Azure SAS %s but was %s", test.knownSasToken, computedToken)
		}
	}
}

func TestComputeAccountSASConnectionString(t *testing.T) {
	testCases := []struct {
		env                 azure.Environment
		accountName         string
		sasToken            string
		sasConnectionString string
	}{
		{
			azure.PublicCloud,
			"testaccount",
			"?st=2019-03-27&se=2019-09-21T09%3A21Z&sp=rwl&sip=93.23.223.54&spr=https&sv=2018-11-09&sr=c&rscc=no-cache&rscd=attachment&rsce=gzip&rscl=en-US&rsct=text/html%3B%20charset%3Dutf-8&sig=M2TaUVEGlRVJjNt/c7Eqt2zH6%2BA8dpiLmTXR0ZevEX8%3D",
			"BlobEndpoint=https://testaccount.blob.core.windows.net/;FileEndpoint=https://testaccount.file.core.windows.net/;QueueEndpoint=https://testaccount.queue.core.windows.net/;TableEndpoint=https://testaccount.table.core.windows.net/;SharedAccessSignature=st=2019-03-27&se=2019-09-21T09%3A21Z&sp=rwl&sip=93.23.223.54&spr=https&sv=2018-11-09&sr=c&rscc=no-cache&rscd=attachment&rsce=gzip&rscl=en-US&rsct=text/html%3B%20charset%3Dutf-8&sig=M2TaUVEGlRVJjNt/c7Eqt2zH6%2BA8dpiLmTXR0ZevEX8%3D",
		},
		{
			azure.ChinaCloud,
			"testaccount",
			"?st=2019-03-27&se=2019-09-21T09%3A21Z&sp=rwl&sip=93.23.223.54&spr=https&sv=2018-11-09&sr=c&rscc=no-cache&rscd=attachment&rsce=gzip&rscl=en-US&rsct=text/html%3B%20charset%3Dutf-8&sig=M2TaUVEGlRVJjNt/c7Eqt2zH6%2BA8dpiLmTXR0ZevEX8%3D",
			"BlobEndpoint=https://testaccount.blob.core.chinacloudapi.cn/;FileEndpoint=https://testaccount.file.core.chinacloudapi.cn/;QueueEndpoint=https://testaccount.queue.core.chinacloudapi.cn/;TableEndpoint=https://testaccount.table.core.chinacloudapi.cn/;SharedAccessSignature=st=2019-03-27&se=2019-09-21T09%3A21Z&sp=rwl&sip=93.23.223.54&spr=https&sv=2018-11-09&sr=c&rscc=no-cache&rscd=attachment&rsce=gzip&rscl=en-US&rsct=text/html%3B%20charset%3Dutf-8&sig=M2TaUVEGlRVJjNt/c7Eqt2zH6%2BA8dpiLmTXR0ZevEX8%3D",
		},
		{
			azure.USGovernmentCloud,
			"testaccount",
			"?st=2019-03-27&se=2019-09-21T09%3A21Z&sp=rwl&sip=93.23.223.54&spr=https&sv=2018-11-09&sr=c&rscc=no-cache&rscd=attachment&rsce=gzip&rscl=en-US&rsct=text/html%3B%20charset%3Dutf-8&sig=M2TaUVEGlRVJjNt/c7Eqt2zH6%2BA8dpiLmTXR0ZevEX8%3D",
			"BlobEndpoint=https://testaccount.blob.core.usgovcloudapi.net/;FileEndpoint=https://testaccount.file.core.usgovcloudapi.net/;QueueEndpoint=https://testaccount.queue.core.usgovcloudapi.net/;TableEndpoint=https://testaccount.table.core.usgovcloudapi.net/;SharedAccessSignature=st=2019-03-27&se=2019-09-21T09%3A21Z&sp=rwl&sip=93.23.223.54&spr=https&sv=2018-11-09&sr=c&rscc=no-cache&rscd=attachment&rsce=gzip&rscl=en-US&rsct=text/html%3B%20charset%3Dutf-8&sig=M2TaUVEGlRVJjNt/c7Eqt2zH6%2BA8dpiLmTXR0ZevEX8%3D",
		},
	}

	for _, test := range testCases {
		computedConnectionString := ComputeAccountSASConnectionString(&test.env, test.accountName, test.sasToken)

		if computedConnectionString != test.sasConnectionString {
			t.Fatalf("Test failed: Expected SAS connection string is %s but was %s", computedConnectionString, test.sasConnectionString)
		}
	}
}

func TestComputeAccountSASConnectionUrlForType(t *testing.T) {
	testCases := []struct {
		env                  azure.Environment
		accountName          string
		sasToken             string
		storageType          string
		storageConnectionUrl string
	}{
		{
			azure.PublicCloud,
			"testaccount",
			"?st=2019-03-27&se=2019-09-21T09%3A21Z&sp=rwl&sip=93.23.223.54&spr=https&sv=2018-11-09&sr=c&rscc=no-cache&rscd=attachment&rsce=gzip&rscl=en-US&rsct=text/html%3B%20charset%3Dutf-8&sig=M2TaUVEGlRVJjNt/c7Eqt2zH6%2BA8dpiLmTXR0ZevEX8%3D",
			"blob",
			"https://testaccount.blob.core.windows.net?st=2019-03-27&se=2019-09-21T09%3A21Z&sp=rwl&sip=93.23.223.54&spr=https&sv=2018-11-09&sr=c&rscc=no-cache&rscd=attachment&rsce=gzip&rscl=en-US&rsct=text/html%3B%20charset%3Dutf-8&sig=M2TaUVEGlRVJjNt/c7Eqt2zH6%2BA8dpiLmTXR0ZevEX8%3D",
		},
		{
			azure.ChinaCloud,
			"testaccount",
			"?st=2019-03-27&se=2019-09-21T09%3A21Z&sp=rwl&sip=93.23.223.54&spr=https&sv=2018-11-09&sr=c&rscc=no-cache&rscd=attachment&rsce=gzip&rscl=en-US&rsct=text/html%3B%20charset%3Dutf-8&sig=M2TaUVEGlRVJjNt/c7Eqt2zH6%2BA8dpiLmTXR0ZevEX8%3D",
			"file",
			"https://testaccount.file.core.chinacloudapi.cn?st=2019-03-27&se=2019-09-21T09%3A21Z&sp=rwl&sip=93.23.223.54&spr=https&sv=2018-11-09&sr=c&rscc=no-cache&rscd=attachment&rsce=gzip&rscl=en-US&rsct=text/html%3B%20charset%3Dutf-8&sig=M2TaUVEGlRVJjNt/c7Eqt2zH6%2BA8dpiLmTXR0ZevEX8%3D",
		},
		{
			azure.USGovernmentCloud,
			"testaccount",
			"?st=2019-03-27&se=2019-09-21T09%3A21Z&sp=rwl&sip=93.23.223.54&spr=https&sv=2018-11-09&sr=c&rscc=no-cache&rscd=attachment&rsce=gzip&rscl=en-US&rsct=text/html%3B%20charset%3Dutf-8&sig=M2TaUVEGlRVJjNt/c7Eqt2zH6%2BA8dpiLmTXR0ZevEX8%3D",
			"table",
			"https://testaccount.table.core.usgovcloudapi.net?st=2019-03-27&se=2019-09-21T09%3A21Z&sp=rwl&sip=93.23.223.54&spr=https&sv=2018-11-09&sr=c&rscc=no-cache&rscd=attachment&rsce=gzip&rscl=en-US&rsct=text/html%3B%20charset%3Dutf-8&sig=M2TaUVEGlRVJjNt/c7Eqt2zH6%2BA8dpiLmTXR0ZevEX8%3D",
		},
		{
			azure.PublicCloud,
			"testaccount",
			"?st=2019-03-27&se=2019-09-21T09%3A21Z&sp=rwl&sip=93.23.223.54&spr=https&sv=2018-11-09&sr=c&rscc=no-cache&rscd=attachment&rsce=gzip&rscl=en-US&rsct=text/html%3B%20charset%3Dutf-8&sig=M2TaUVEGlRVJjNt/c7Eqt2zH6%2BA8dpiLmTXR0ZevEX8%3D",
			"unexpected",
			"",
		},
	}

	for _, test := range testCases {
		computedStorageConnectionUrl, err := ComputeAccountSASConnectionUrlForType(&test.env, test.accountName, test.sasToken, test.storageType)
		if strings.Compare("unexpected", test.storageType) == 0 {
			if err == nil {
				t.Fatalf("Test failed: This call should have thrown an error because an unexpected storage type was specified.")
			}
		} else if err != nil {
			t.Fatalf("Test failed: This call should not have thrown an error")
		} else if strings.Compare(*computedStorageConnectionUrl, test.storageConnectionUrl) != 0 {
			t.Fatalf("Test failed: Expected connection url is %s but was %s", *computedStorageConnectionUrl, test.storageConnectionUrl)
		}
	}
}

func compareSASTokens(token1 string, token2 string) bool {
	queryParams1 := parseSASToken(token1)
	queryParams2 := parseSASToken(token2)

	if len(queryParams1) != len(queryParams2) {
		return false
	}

	for k, v1 := range queryParams1 {
		if v2, ok := queryParams2[k]; ok {
			// values need to be unescaped because apperently azure cli escape does not seem to work correctly (e.g. for contentType value)
			// example: text/html; charset=utf-8
			// -> escaped in go: text%2Fhtml%3B+charset%3Dutf-8
			// -> escaped in python / azure cli: text/html%3B%20charset%3Dutf-8
			unescapedValue1, _ := url.QueryUnescape(v1)
			unescapedValue2, _ := url.QueryUnescape(v2)
			if unescapedValue1 != unescapedValue2 {
				return false
			}
		} else {
			return false
		}
	}

	return true
}

func parseSASToken(token string) map[string]string {
	queryParts := strings.Split(token[1:], "&")

	kvp := make(map[string]string)
	for _, queryPart := range queryParts {
		kv := strings.SplitN(queryPart, "=", 2)
		key := kv[0]
		value := kv[1]
		kvp[key] = value
	}

	return kvp
}
