package main

import (
	"log"
	"net/http"

	"github.com/justinas/alice"
)

func main() {
	mainHandler := http.FileServer(http.Dir(Config.FilesPath))
	chain := alice.New(
		logger,
		registerToken,
		verifyToken,
		s3Redirect,
	).Then(mainHandler)

	log.Printf("Starting up server on %s\n", Config.Listen)
	http.ListenAndServe(Config.Listen, chain)
}
