package main

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"io"
	"io/ioutil"
)

func UnpackBody(body io.ReadCloser, to interface{}) error {
	rawBody, err := ioutil.ReadAll(body)
	defer body.Close()
	if err != nil {
		return err
	}
	err = json.Unmarshal(rawBody, to)
	return err
}

func HashPassword(s string) string {
	hash := sha256.New()
	hash.Write([]byte(s))
	hashBytes := hash.Sum(nil)
	encoded := base64.StdEncoding.EncodeToString(hashBytes)
	return encoded
}
