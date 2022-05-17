package main

import (
	"crypto/sha1"
)

func sha1Sum(src []byte) ([]byte, error) {
	hasher := sha1.New()

	_, err := hasher.Write(src)
	if err != nil {
		return nil, err
	}

	return hasher.Sum(nil), nil
}
