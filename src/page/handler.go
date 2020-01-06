package page

import (
	"net/http"
)

import (
	"go-react-template/src/templateManager"
)
 
func Handler(w http.ResponseWriter, r *http.Request) {
	templateManager.RenderHtmlWithParameters(w, "index.html", nil);
}