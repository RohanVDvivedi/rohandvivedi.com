package main

// go utilities
import (
	"time"
	"os"
	"fmt"
	"log"
	"net/http"
	"golang.org/x/net/websocket"
	"strings"
)

// maintains global configuration for the application
import (
	"rohandvivedi.com/src/config"
)

// maintains session, (in memory)
import (
	"rohandvivedi.com/src/session"
)

// data / databse
import (
    "database/sql"
    _ "github.com/mattn/go-sqlite3"
    "rohandvivedi.com/src/data"
)

// template manager to write templated html files as strings
import (
	"rohandvivedi.com/src/templateManager"
)

// mail client support for the application
import (
	"rohandvivedi.com/src/mailManager"
	"rohandvivedi.com/src/mails"
)

// handlers for the pages and apis
import (
	"rohandvivedi.com/src/page"
	"rohandvivedi.com/src/api"
	"rohandvivedi.com/src/socket"
)

// the fitst command line argument has to be "prod" for production
// no arguments or "dev", results in starting the go backend in development mode
func main() {
	environment := "dev"
	if(len(os.Args) >= 2){
		environment = os.Args[1]
	}

	// initialize the global configuration form the appropriate config file, and initialize it
	config.InitGlobalConfig(environment)

	// this will ask the template manager to initialize the templates variable
	templateManager.InitializeTemplateEngine()

	// create a server mux
	mux := http.NewServeMux();

	// we use a FileServer to host the static contents of the website (js, css, img)
	fs := http.FileServer(http.Dir("public/static"))
	mux.Handle("/", Send404OnFolderRequest(fs))

	// attach all the handlers of all the pages here
	// we have only one page handler, because this is a react app, but will have many apis
	mux.HandleFunc("/pages/", page.Handler);

	// attach all the handlers for websockets here
	// we have only one page handler, because this is a react app, but will have many apis
	mux.Handle("/soc", websocket.Handler(socket.Handler));

	// attach all the handlers of all the apis here
	// we have only one page handler, because this is a react app, but will have many apis
	mux.Handle("/api/person", 			CountApiHitsInSessionValues(api.GetPerson));
	mux.Handle("/api/project", 			CountApiHitsInSessionValues(api.FindProject));
	mux.Handle("/api/all_categories", 	CountApiHitsInSessionValues(api.GetAllCategories));
	mux.Handle("/api/owner", 			CountApiHitsInSessionValues(api.GetOwner));
	mux.Handle("/api/sessions", 		AuthorizeIfOwner(api.PrintAllUserSessions));

	// setup database connection
	data.Db, _ = sql.Open("sqlite3", "./db/data.db")
	defer data.Db.Close()
	data.InitializeSchema()

	// set up session store
	ownerSessionId := ""
	if(config.GetGlobalConfig().Create_user_sessions) {
		fmt.Println("Initializing SessionStore");
		session.InitGlobalSessionStore("r_sess_id", 96 * time.Hour)
		ownerSessionId = session.GlobalSessionStore.InitializeOwnerSession().SessionId
	} else {
		fmt.Println("Configuration declines setting up of SessionStore");
		ownerSessionId = "****No SessionStore, hence no OwnerSessionId****"
	}

	// initialize mail smtp client, and authenticate
	if(config.GetGlobalConfig().Auth_mail_client) {
		fmt.Println("Initializing and Authenticating SMTP mail client");
		mailManager.InitMailClient(config.GetGlobalConfig().From_mailid, config.GetGlobalConfig().From_password)
		mails.SendDeploymentMail(ownerSessionId)
	} else {
		fmt.Println("Configuration declines setting up of SMTP mail client");
	}
	
	if(!config.GetGlobalConfig().SSL_enabled){
		fmt.Println("Application starting with ssl disabled on port 80");
		log.Fatal(http.ListenAndServe(":80", mux));
	} else {
		fmt.Println("Application starting with SSL enabled on port 443");
		log.Fatal(http.ListenAndServeTLS(":443",
			"/etc/letsencrypt/live/rohandvivedi.com/fullchain.pem",
			"/etc/letsencrypt/live/rohandvivedi.com/privkey.pem", mux))
	}
	fmt.Println("Application shutdown");
}

// Below are the middlewares for the http api handlers

// this function is a middleware to send 404 response, if the requested path is a folder
// i.e. request path ending in "/"
func Send404OnFolderRequest(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        if strings.HasSuffix(r.URL.Path, "/") && len(r.URL.Path) != 1 {
            http.NotFound(w, r)
            return
        }
        next.ServeHTTP(w, r)
    })
}

// this middleware lets you maintain data regarding api access that each sessioned user has made
func CountApiHitsInSessionValues(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        
        if(config.GetGlobalConfig().Create_user_sessions) {
	        // you must need a session to allow me to maintain the count
	        s := session.GlobalSessionStore.GetOrCreateSession(w, r);
			_ = s.ExecuteOnValues(func (SessionVals map[string]interface{}, add_par interface{}) interface{} {
				reqPathCountKey := "<" + r.URL.Path + ">_count"		// this is the key we will use to store count of hits in session values
				count, exists := SessionVals[reqPathCountKey]
				if(exists) {
					intCount, isInt := count.(int)
					if isInt {
						SessionVals[reqPathCountKey] = intCount + 1
						return nil
					}
				}
				SessionVals[reqPathCountKey] = 1
				return nil
			}, nil);
		}

        next.ServeHTTP(w, r)
    })
}

func AuthorizeIfOwner(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        
        if(config.GetGlobalConfig().Create_user_sessions) {
	        // you must need a session to allow me to maintain the count
	        s := session.GlobalSessionStore.GetOrCreateSession(w, r);
			isOwner, ownerKeyExists := s.GetValue("owner");
			if(ownerKeyExists) {
				valIsOwner, ok := isOwner.(bool)
				if(ok && valIsOwner){
					next.ServeHTTP(w, r)
					return
				}
			}
		}

		// if any thing fails, just unautorize
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("You are not owner of rohandvivedi.com"))

    })
}
