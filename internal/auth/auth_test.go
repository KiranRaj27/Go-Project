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
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			key, err := GetApiKey(tc.headers)

			if err != nil && err.Error() != tc.expectedError {
				t.Errorf("unexpected error; got %s, want %s", err.Error(), tc.expectedError)
			}

			if key != tc.expectedKey {
				t.Errorf("unexpected key; got %s, want %s", key, tc.expectedKey)
			}
		})
	}
}
