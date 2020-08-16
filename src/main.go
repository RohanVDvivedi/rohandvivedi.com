package main

// go utilities
import (
	"os"
	"fmt"
	"log"
	"net/http"
	"golang.org/x/net/websocket"
	"strings"
)

// data
import (
    "database/sql"
    _ "github.com/mattn/go-sqlite3"
)

// utilities
import (
	"rohandvivedi.com/src/templateManager"
	"rohandvivedi.com/src/mailManager"
)

// handlers for the pages and apis
import (
	"rohandvivedi.com/src/page"
	"rohandvivedi.com/src/api"
	"rohandvivedi.com/src/socket"
)

// data
import (
	"rohandvivedi.com/src/data"
)

var ssl_enabled bool = true;

func main() {
	// this will ask the template manager to initialize the templates variable
	templateManager.InitializeTemplateEngine()

	// we use a FileServer to host the static contents of the website (js, css, img)
	fs := http.FileServer(http.Dir("public/static"))
	http.Handle("/", handlerForFolder404(fs))

	// attach all the handlers of all the pages here
	// we have only one page handler, because this is a react app, but will have many apis
	http.HandleFunc("/pages/", page.Handler);

	// attach all the handlers for websockets here
	// we have only one page handler, because this is a react app, but will have many apis
	http.Handle("/soc", websocket.Handler(socket.Handler));

	// attach all the handlers of all the apis here
	// we have only one page handler, because this is a react app, but will have many apis
	http.HandleFunc("/api/person",api.GetPerson);
	http.HandleFunc("/api/project", api.FindProject);
	http.HandleFunc("/api/all_categories", api.GetAllCategories);
	http.HandleFunc("/api/owner", api.GetOwner);

	// initialize mail client
	mailManager.InitMailClient(os.Getenv("EMAIL_PASS"))

	// setup database connection
	data.Db, _ = sql.Open("sqlite3", "./db/data.db")
	defer data.Db.Close()
	data.InitializeSchema()

	
	fmt.Println("Application starting (config: ssl enabled ", ssl_enabled, ")");
	if(!ssl_enabled){
		log.Fatal(http.ListenAndServe(":80", nil));
	} else {
		log.Fatal(http.ListenAndServeTLS(":443",
			"/etc/letsencrypt/live/rohandvivedi.com/fullchain.pem",
			"/etc/letsencrypt/live/rohandvivedi.com/privkey.pem", nil))
	}
	fmt.Println("Application shutdown");
}

// this function is a handler to send 404 response, if the requested path is a folder
// i.e. ending in /
func handlerForFolder404(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        if strings.HasSuffix(r.URL.Path, "/") && len(r.URL.Path) != 1 {
            http.NotFound(w, r)
            return
        }
        next.ServeHTTP(w, r)
    })
}