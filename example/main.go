package main

import (
	_ "embed"
	"fmt"
	"github.com/mar-coding/swaggerui"
	"log"
	"net/http"
)

//go:embed spec/petstore.yml
var spec []byte
var port = 3000

func main() {
	log.SetFlags(0)
	mux := http.NewServeMux()
	mux.Handle("/swagger/", http.StripPrefix("/swagger", swaggerui.Handler(spec)))
	log.Printf("Listening on port %d...", port)

	err := http.ListenAndServe(fmt.Sprintf(":%d", port), mux)
	if err != nil {
		log.Fatal(err)
	}
}
