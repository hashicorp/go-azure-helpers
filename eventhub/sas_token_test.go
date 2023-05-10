// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package eventhub

import (
	"strings"
	"testing"
)

func TestParseEventHubConnectionString(t *testing.T) {
	testCases := []struct {
		input                       string
		expectedSharedAccessKeyName string
		expectedSharedAccessKey     string
		expectedError               bool
	}{
		{
			"Endpoint=sb://acctesteventhubnamespace-test01.servicebus.windows.net/;SharedAccessKeyName=RootManageSharedAccessKey;SharedAccessKey=IUSvXLiPZ3uAQcso/cL7vTiL4zsc/EMtcUzNCC2dhaM=",
			"RootManageSharedAccessKey",
			"IUSvXLiPZ3uAQcso/cL7vTiL4zsc/EMtcUzNCC2dhaM=",
			false,
		},
		{
			"Endpoint=sb://acctesteventhubnamespace-test01.servicebus.windows.net/;SharedAccessKeyName=acctest-test01;SharedAccessKey=R9v9VaHiU/ktFIka8Q4aUbQnZeiSKaevncrsxOTTILw=;EntityPath=acctesteventhub-test01",
			"acctest-test01",
			"R9v9VaHiU/ktFIka8Q4aUbQnZeiSKaevncrsxOTTILw=",
			false,
		},
		{
			"Endpoint=sb://acctesteventhubnamespace-test01.servicebus.windows.net/;SharedAccessKeyName=acctest-test01;SharedAccessKey=R9v9VaHiU/ktFIka8Q4aUbQnZeiSKaevncrsxOTTILw=;EntityPath",
			"",
			"",
			true,
		},
	}

	for _, test := range testCases {
		result, err := ParseEventHubSASConnectionString(test.input)

		if test.expectedError {
			if err == nil {
				t.Fatalf("Expected error for %s: %q", test.input, err)
			}
			return
		}

		if !test.expectedError && err != nil {
			t.Fatalf("Failed to parse resource type string: %s, %q", test.input, result)
		}

		if val, pres := result[connStringSharedAccessKeyKey]; !pres || val != test.expectedSharedAccessKey {
			t.Fatalf("Failed to parse Shared Access Key: Expected: %s, Found: %s", test.expectedSharedAccessKey, val)
		}
		if val, pres := result[connStringSharedAccessKeyNameKey]; !pres || val != test.expectedSharedAccessKeyName {
			t.Fatalf("Failed to parse Shared Access Name: Expected: %s, Found: %s", test.expectedSharedAccessKeyName, val)
		}
	}
}

func TestComputeEventHubSASToken(t *testing.T) {
	testCases := []struct {
		sharedAccessKeyName string
		sharedAccessKey     string
		eventHubUri         string
		expiry              string
		knownSasToken       string
	}{
		{
			"RootManageSharedAccessKey",
			"IUSvXLiPZ3uAQcso/cL7vTiL4zsc/EMtcUzNCC2dhaM=",
			"sb://acctesteventhubnamespace-test01.servicebus.windows.net",
			"2022-01-11T08:24:49Z",
			"sr=sb%3A%2F%2Facctesteventhubnamespace-test01.servicebus.windows.net&sig=8dgxKVwLsOWxX7f4mNtyiez47lxYCJ8h%2FeViD%2BMWY2E%3D&se=1641889489&skn=RootManageSharedAccessKey",
		},
		{
			"acctest-test01",
			"R9v9VaHiU/ktFIka8Q4aUbQnZeiSKaevncrsxOTTILw=",
			"sb://acctesteventhubnamespace-test01.servicebus.windows.net/acctesteventhub-test01",
			"2022-01-11T08:24:49Z",
			"sr=sb%3A%2F%2Facctesteventhubnamespace-test01.servicebus.windows.net%2Facctesteventhub-test01&sig=%2FaPTDsDZhwpdysw1klgV1fm5a%2Bo3vw2Lb7HsDHyZr4M%3D&se=1641889489&skn=acctest-test01",
		},
	}

	for _, test := range testCases {
		computedToken, err := ComputeEventHubSASToken(test.sharedAccessKeyName,
			test.sharedAccessKey,
			test.eventHubUri,
			test.expiry)

		if err != nil {
			t.Fatalf("Test Failed: Error computing Event Hub Sas: %q", err)
		}

		if computedToken != test.knownSasToken {
			t.Fatalf("Test failed: Expected Azure SAS %s but was %s", test.knownSasToken, computedToken)
		}
	}
}

func TestComputeEventHubSASConnectionString(t *testing.T) {
	testCases := []struct {
		sasToken            string
		sasConnectionString string
	}{
		{
			"sr=sb%3A%2F%2Facctest-ehn-test01.servicebus.windows.net%2Facctest-eh-test01&sig=ozpLwoOHPAWD1s4GE2Khhu508JbcVA4%2FWutXZIV7VfI%3D&se=1672531200&skn=acctest-ehar-test01",
			"SharedAccessSignature sr=sb%3A%2F%2Facctest-ehn-test01.servicebus.windows.net%2Facctest-eh-test01&sig=ozpLwoOHPAWD1s4GE2Khhu508JbcVA4%2FWutXZIV7VfI%3D&se=1672531200&skn=acctest-ehar-test01",
		},
	}

	for _, test := range testCases {
		computedConnectionString := ComputeEventHubSASConnectionString(test.sasToken)

		if computedConnectionString != test.sasConnectionString {
			t.Fatalf("Test failed: Expected SAS connection string is %s but was %s", computedConnectionString, test.sasConnectionString)
		}
	}
}

func TestComputeEventHubSASConnectionUrl(t *testing.T) {
	testCases := []struct {
		endpoint              string
		entityPath            string
		eventHubConnectionUrl string
	}{
		{
			"sb://acctesteventhubnamespace-test01.servicebus.windows.net/",
			"acctesteventhub-test01",
			"sb://acctesteventhubnamespace-test01.servicebus.windows.net/acctesteventhub-test01",
		},
		{
			"sb://acctesteventhubnamespace-test01.servicebus.windows.net/",
			"",
			"sb://acctesteventhubnamespace-test01.servicebus.windows.net",
		},
	}

	for _, test := range testCases {
		computedEventHubConnectionUrl, err := ComputeEventHubSASConnectionUrl(test.endpoint, test.entityPath)
		if err != nil {
			t.Fatalf("Test failed: This call should not have thrown an error")
		} else if strings.Compare(*computedEventHubConnectionUrl, test.eventHubConnectionUrl) != 0 {
			t.Fatalf("Test failed: Expected connection url is %s but was %s", *computedEventHubConnectionUrl, test.eventHubConnectionUrl)
		}
	}
}
