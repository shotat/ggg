package main

import "github.com/shotat/ggg/interfaces"
import "net/http"

func main() {
	http.ListenAndServe(":8080", &interfaces.Handler{})
}
