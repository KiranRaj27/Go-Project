package auth

import (
	"errors"
	"net/http"
	"strings"
)

func GetApiKey(headers http.Header) (string, error) {
	val := headers.Get("Authorization")
	if val == "" {
		return "", errors.New("no authontication info found")
	}

	vals := strings.Split(val, " ")
	if len(vals) != 2 {
		return "", errors.New("malformed header length is not two")
	}

	if vals[0] != "ApiKey" {
		return "", errors.New("malformed header")
	}
	return vals[1], nil
}
