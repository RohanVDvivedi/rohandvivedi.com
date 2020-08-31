package page

import (
	"net/http"
)

import (
	"rohandvivedi.com/src/templateManager"
)

var PageHandler = http.HandlerFunc(pageHandler)
 
func pageHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	templateManager.RenderHtmlWithParameters(w, "index.html", nil);
}