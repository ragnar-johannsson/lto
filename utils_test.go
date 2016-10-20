package main

import (
	"strings"
	"testing"
)

func TestGenerateToken(t *testing.T) {
	token := generateToken()
	defaultLength := 24

	if len(token) != defaultLength {
		t.Errorf("Generated random token not of correct default length: %s", token)
	}
}

func TestGenerateJSON(t *testing.T) {
	Config.BaseUrl = "http://a.host/"
	key := "path"
	token := "abc123"
	out := generateJSON(key, token)
	expected := `{"url":"http://a.host/path?token=abc123"}`

	if strings.Compare(out, expected) != 0 {
		t.Errorf("Generated JSON not as expected: %s", out)
	}
}
