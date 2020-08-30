package page

import (
	"net/http"
)

import (
	"rohandvivedi.com/src/templateManager"
)

var PageHandler = http.HandlerFunc(pageHandler)
 
func pageHandler(w http.ResponseWriter, r *http.Request) {
	templateManager.RenderHtmlWithParameters(w, "index.html", nil);
}