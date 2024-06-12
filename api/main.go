package main

import (
	"log"
	"net/http"

	"github.com/Lucasmartinsn/DevBook/api/src/router"
)

func main() {
	r := router.Generator()
	log.Fatal(http.ListenAndServe(":5000", r))
}
