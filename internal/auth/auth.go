package auth

import (
	"errors"
	"net/http"
	"strings"
)

// "Authorization" : "Bearer {API_KEY}"
func GetApiKey(header http.Header) (string, error) {
	val := header.Get("Authorization")

	if val == "" {
		return "", errors.New("no authentication info found")
	}

	vals := strings.Split(val, " ")

	if len(vals) != 2 {
		return "", errors.New("wrong authorization header format")
	}

	if vals[0] != "Bearer" {
		return "", errors.New("wrong authorization header format")
	}

	return vals[1], nil
}
