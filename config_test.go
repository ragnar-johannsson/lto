package main

import (
	"testing"
)

func TestDefaultConfigValues(t *testing.T) {
	if Config.UrlTTL.Seconds() != 3600.0 {
		t.Error("Incorrect default UrlTTL: ", Config.UrlTTL.Seconds())
	}
}
