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

// utilities
import (
	"rohandvivedi.com/src/templateManager"
	//"rohandvivedi.com/src/data/mysql"
	//"rohandvivedi.com/src/data/memcached"
	"rohandvivedi.com/src/mailManager"
)

// handlers for the pages and apis
import (
	"rohandvivedi.com/src/page"
	"rohandvivedi.com/src/api"
	"rohandvivedi.com/src/socket"
	"rohandvivedi.com/src/api/project"
)

func main() {
	// this will ask the template manager to initialize the templates variable
	templateManager.InitializeTemplateEngine()

	// we use a FileServer to host the static contents of the website (js, css, img)
	fs := http.FileServer(http.Dir("public/static"))
	http.Handle("/", handlerForFolder404(fs))

	// attach all the handlers of all the pages here
	// we have only one page handler, because this is a react app, but will have many apis
	http.HandleFunc("/index", page.Handler);

	// attach all the handlers for websockets here
	// we have only one page handler, because this is a react app, but will have many apis
	http.Handle("/soc", websocket.Handler(socket.Handler));

	// attach all the handlers of all the apis here
	// we have only one page handler, because this is a react app, but will have many apis
	http.HandleFunc("/api", api.Handler);
	http.HandleFunc("/api/project", project.Handler);
	
	// before we start listenning and we start to serve, start the database connections,
	// both to the cache and the sql database
	// memcached.Initialize();
	// mysql.Initialize();
	// defer mysql.Close();

	mailManager.InitMailClient(os.Getenv("EMAIL_PASS"))
	fmt.Println(os.Getenv("EMAIL_PASS"))
	dest := []string{"rohan.dvivedi@oyorooms.com", "rohan.dvivedi@belvilla.com"};
	mail := mailManager.WritePlainEmail(dest, "hello", "Hello World!!");
	err := mailManager.SendMail(dest, "hello", mail);
	if(err != nil){
		fmt.Println(err);
	}
	
	fmt.Println("Application starting");
	log.Fatal(http.ListenAndServe(":80", nil));
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