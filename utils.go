package main

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"net/url"
)

func generateToken() string {
	token := make([]byte, Config.TokenSize)
	rand.Read(token[:])

	return fmt.Sprintf("%x", token)
}

func generateJSON(key, token string) string {
	u, _ := url.Parse(Config.BaseUrl)
	v := url.Values{}
	v.Add(TOKEN_KEY, token)
	u.Path = key
	u.RawQuery = v.Encode()

	m := map[string]string{"url": u.String()}
	out, err := json.Marshal(m)
	if err != nil {
		return ""
	}

	return string(out)
}
