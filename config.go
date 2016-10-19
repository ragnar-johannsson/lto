package main

import (
	"time"

	"github.com/kelseyhightower/envconfig"
)

func init() {
	envconfig.MustProcess("", &Config)
}

type EnvConfig struct {
	RedisAddr   string `default:"redis:6379" envconfig:"redis_addr"`
	RedisPasswd string `default:"" envconfig:"redis_passwd"`
	RedisDB     int    `default:"0" envconfig:"redis_db"`

	S3AccessKey string `envconfig:"s3_access_key"`
	S3SecretKey string `envconfig:"s3_secret_key"`
	S3Region    string `envconfig:"s3_region"`
	S3Bucket    string `envconfig:"s3_bucket"`

	FilesPath string `envconfig:"files_path"`

	UrlTTL    time.Duration `default:"3600s" envconfig:"url_ttl"`
	UrlSecret string        `envconfig:"url_secret"`
	BaseUrl   string        `envconfig:"base_url"`
	TokenSize int           `default:"12"`
	Listen    string        `default:":3000"`
}

func (e *EnvConfig) LocalFileServer() bool {
	return e.FilesPath != ""
}

func (e *EnvConfig) S3FileServer() bool {
	return (e.S3AccessKey != "" &&
		e.S3SecretKey != "" &&
		e.S3Region != "" &&
		e.S3Bucket != "")
}

var Config EnvConfig
