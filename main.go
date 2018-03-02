package main

import (
	"net/http"

	"github.com/shotat/ggg/interfaces"
)

func main() {
	api := interfaces.NewApi()
	http.ListenAndServe(":8080", api)
}
