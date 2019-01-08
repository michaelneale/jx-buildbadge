package main

import (
	"net/http"
	"strings"
)

func main() {
	http.HandleFunc("/badge/", serverBadge)
	http.ListenAndServe(":8080", nil)
}

func serverBadge(w http.ResponseWriter, r *http.Request) {
	branchID := strings.TrimPrefix(r.URL.Path, "/badge/")
	w.Header().Set("Content-type", "image/png")
	http.ServeFile(w, r, stateIcon(branchID))
}

func stateIcon(branchID string) string {
	switch branchID {
	case "show-fail":
		return "fail.png"
	case "show-unstable":
		return "unstable.png"
	default:
		return fetchState(branchID)
	}
}

/* everything after the /badge/ - use it to work out what state is.
   Typically includes repo and branch name. If just repo, default to master.  */
func fetchState(branchID string) string {
	//TODO implement me!
	return "success.png"
}
