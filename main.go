package main

import (
	"net/http"

	"github.com/shotat/ggg/interfaces"
)

func main() {
	http.ListenAndServe(":8080", interfaces.NewAppHandler())
}
