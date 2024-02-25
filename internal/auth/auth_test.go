package auth

import (
	"net/http"
	"testing"
)

func TestGetApiKey(t *testing.T) {
	// Define test cases
	testCases := []struct {
		name          string
		headers       http.Header
		expectedKey   string
		expectedError string
	}{
		{
			name: "ValidApiKey",
			headers: http.Header{
				"Authorization": []string{"ApiKey myApiKey123"},
			},
			expectedKey:   "myApiKey123",
			expectedError: "",
		},
		{
			name: "NoAuthorizationInfo",
			headers: http.Header{
				"Content-Type": []string{"application/json"},
			},
			expectedKey:   "",
			expectedError: "no authentication info found", // Corrected typo
		},
		{
			name: "MalformedHeader",
			headers: http.Header{
				"Authorization": []string{"Bearer myToken"},
			},
			expectedKey:   "",
			expectedError: "malformed header",
		},
		{
			name: "MalformedHeaderLengthNotTwo",
			headers: http.Header{
				"Authorization": []string{"ApiKey"},
			},
			expectedKey:   "",
			expectedError: "malformed header length is not two",
		},
	}

	// Run test cases
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			key, err := GetApiKey(tc.headers)

			// Check if the error matches the expected error
			if err != nil && err.Error() != tc.expectedError {
				t.Errorf("Expected error: %v, got: %v", tc.expectedError, err.Error())
			}

			// Check if the key matches the expected key
			if key != tc.expectedKey {
				t.Errorf("Expected key: %v, got: %v", tc.expectedKey, key)
			}
		})
	}
}
