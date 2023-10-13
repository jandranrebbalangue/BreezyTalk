package main

import (
	"fmt"
	"net/http"
	"strings"
)

func PlayerServer(w http.ResponseWriter, r *http.Request) {
	player := strings.TrimPrefix(r.URL.Path, "/players/")
	fmt.Fprint(w, GetPlayerScore(player))
}
func GetPlayerScore(name string) string {

	if name == "patat" {
		return "20"
	}
	if name == "floyd" {
		return "10"
	}
	return ""
}
