package auth

import (
	"net/http"
	"strings"
	"testing"
)

func TestGetApiKey(t *testing.T) {
	tests := []struct {
		name        string
		headerKey   string
		headerVal   string
		expectedKey string
		expectedErr string
	}{
		{
			name:        "No Auth Header",
			expectedErr: ErrNoAuthHeaderIncluded.Error(),
		},
		{
			name:        "Empty Authorization Header",
			headerKey:   "Authorization",
			expectedErr: ErrNoAuthHeaderIncluded.Error(),
		},
		{
			name:        "Malformed Header",
			headerKey:   "Authorization",
			headerVal:   "-",
			expectedErr: ErrMalformedAuthHeader.Error(),
		},
		{
			name:        "Wrong Header Prefix",
			headerKey:   "Authorization",
			headerVal:   "Bearer xxxxxx",
			expectedErr: ErrMalformedAuthHeader.Error(),
		},
		{
			name:        "Correctly Formatted",
			headerKey:   "Authorization",
			headerVal:   "ApiKey xxxxxx",
			expectedKey: "xxxxxx",
			expectedErr: "",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			header := http.Header{}
			header.Add(tc.headerKey, tc.headerVal)

			apiKey, err := GetAPIKey(header)
			if err != nil {
				if strings.Contains(err.Error(), tc.expectedErr) {
					return
				}
				t.Errorf("Unexpected: TestGetApiKey:%v\n", err)
				return
			}

			if apiKey != tc.expectedKey {
				t.Errorf("Unexpected: TestGetApiKey:%s\n", apiKey)
				return
			}
		})
	}

}
