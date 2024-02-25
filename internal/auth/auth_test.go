package auth

import (
	"net/http"
	"testing"
)

func TestGetApiKey(t *testing.T) {
	// Define test cases
	tests := []struct {
		name          string
		headers       http.Header
		expectedKey   string
		expectedError string
	}{
		{
			name: "Valid Authorization Header",
			headers: http.Header{
				"Authorization": []string{"ApiKey my-api-key"},
			},
			expectedKey:   "my-api-key",
			expectedError: "",
		},
		{
			name: "No Authorization Header",
			headers: http.Header{
				"Another-Header": []string{"Some value"},
			},
			expectedKey:   "",
			expectedError: "no authentication info found",
		},
		{
			name: "Malformed Authorization Header (Length)",
			headers: http.Header{
				"Authorization": []string{"ApiKey"},
			},
			expectedKey:   "",
			expectedError: "malformed header length is not two",
		},
		{
			name: "Malformed Authorization Header (Type)",
			headers: http.Header{
				"Authorization": []string{"InvalidType my-api-key"},
			},
			expectedKey:   "",
			expectedError: "malformed header",
		},
	}

	// Iterate over test cases
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Call the function being tested
			key, err := GetApiKey(tc.headers)

			// Check if the error matches the expected error
			if err != nil && err.Error() != tc.expectedError {
				t.Errorf("unexpected error; got %s, want %s", err.Error(), tc.expectedError)
			}

			// Check if the key matches the expected key
			if key != tc.expectedKey {
				t.Errorf("unexpected key; got %s, want %s", key, tc.expectedKey)
			}
		})
	}
}
