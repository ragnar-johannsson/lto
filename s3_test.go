package main

import (
	"net/url"
	"strings"
	"testing"
	"time"
)

func TestGenerateSignedURL(t *testing.T) {
	Config.S3AccessKey = "abcd1234abcd1234"
	Config.S3SecretKey = "1234abcd1234abcd"
	Config.S3Region = "us-east-1"
	Config.S3Bucket = "bucket"
	Config.UrlTTL, _ = time.ParseDuration("2h")

	signed := generateSignedURL("key")
	u, err := url.Parse(signed)
	if err != nil {
		t.Error("Error parsing signed url", err)
	}

	signature := u.Query().Get("X-Amz-Signature")
	credential := u.Query().Get("X-Amz-Credential")
	expires := u.Query().Get("X-Amz-Expires")

	if len(signature) == 0 {
		t.Error("Signature not generated")
	}

	if !strings.Contains(credential, Config.S3AccessKey) ||
		!strings.Contains(credential, Config.S3Region) {

		t.Error("Invalid credentials")
	}

	if expires != "7200" {
		t.Error("Expiry not correct")
	}
}
