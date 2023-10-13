package main

import (
	"fmt"
	"net/http"
)

func PlayerServer(w http.ResponseWriter, request *http.Request) {
	fmt.Fprint(w, "20")
}
