package main

import (
	"log"
	"net/http"
	"strings"
)

func main() {
	http.HandleFunc("/badge/", ServerBadge)
	http.ListenAndServe(":8080", nil)
}

func ServerBadge(w http.ResponseWriter, r *http.Request) {
	branchID := strings.TrimPrefix(r.URL.Path, "/badge/")
	w.Header().Set("Content-type", "image/png")
	http.ServeFile(w, r, StateIcon(branchID))
}

func StateIcon(branchID string) string {
	switch branchID {
	case "show-fail":
		return "fail.png"
	case "show-unstable":
		return "unstable.png"
	default:
		return FetchState(branchID)
	}
}

/* everything after the /badge/ - use it to work out what state is.
   Typically includes repo and branch name. If just repo, default to master.  */
func FetchState(branchID string) string {
	pipeline, branch := SplitBranch(branchID)
	log.Printf("pipe = %s  branch = %s ", pipeline, branch)
	//TODO implement me to look things up from Jenkins X here!
	return "success.png"

}

func SplitBranch(branchID string) (string, string) {
	elems := strings.Split(branchID, "/")
	return elems[0], elems[1]

}
