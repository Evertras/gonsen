package gonsen

import "net/http"

func addHtmlHeaders(w http.ResponseWriter) {
	w.Header().Add("Content-Type", "text/html")
}
