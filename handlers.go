package main

import (
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/ragnar-johannsson/rww"
)

const (
	REGISTER_KEY = "register"
	TOKEN_KEY    = "token"
)

func logger(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ww := rww.New(w)

		h.ServeHTTP(ww, r)

		log.Printf(
			"- %s - \"%s %s %s\" %d %d",
			r.RemoteAddr,
			r.Method,
			r.URL.RequestURI(),
			r.Proto,
			ww.Status,
			ww.Size,
		)
	})
}

func registerToken(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Query().Get(REGISTER_KEY) == Config.UrlSecret {
			token := generateToken()
			path := r.URL.EscapedPath()
			url := ""

			if Config.S3FileServer() {
				url = generateSignedURL(path)
			}

			Redis.Set(token, path, url, Config.UrlTTL)

			w.Header().Set("Content-Type", "application/json")
			io.WriteString(w, generateJSON(path, token))

			return
		}

		h.ServeHTTP(w, r)
	})
}

func verifyToken(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.URL.Query().Get(TOKEN_KEY)
		if token == "" {
			http.Error(w, "No token specified", http.StatusUnauthorized)
			return
		}

		path := r.URL.EscapedPath()
		rpath, _, err := Redis.Get(token)
		if err != nil || strings.Compare(path, rpath) != 0 {
			http.Error(w, "Token not found or expired", http.StatusUnauthorized)
			return
		}

		h.ServeHTTP(w, r)
	})
}

func s3Redirect(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.URL.Query().Get(TOKEN_KEY)
		_, url, err := Redis.Get(token)
		ww := rww.New(w)

		if Config.S3FileServer() && err == nil {
			ww.AddIntercept(
				http.StatusNotFound,
				http.StatusTemporaryRedirect,
				func(d []byte) (int, error) {
					return len(d), nil
				},
				map[string]string{
					"Location": url,
				},
			)
		}

		h.ServeHTTP(ww, r)
	})
}
