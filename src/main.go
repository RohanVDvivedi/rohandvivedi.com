package main

// go utilities
import (
	"time"
	"os"
	"fmt"
	"log"
	"net/http"
	"golang.org/x/net/websocket"
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
    "rohandvivedi.com/src/searchindex"
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
)

// web socket handler for chatting 
import (
	"rohandvivedi.com/src/chatter"
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
	mux.Handle("/", CountApiHitsInSessionValues(GzipCompressor(Send404OnFolderRequest(fs))))

	// attach all the handlers of all the pages here
	mux.Handle("/pages/", CountApiHitsInSessionValues(GzipCompressor(page.PageHandler)));

	// attach all the handlers for websockets here
	mux.Handle("/chat", AuthorizeIfHasSession(chatter.AuthorizeChat(CountApiHitsInSessionValues(websocket.Handler(chatter.ChatHandler)))));

	// attach all the handlers of all the apis here
	mux.Handle("/api/person", 				CountApiHitsInSessionValues(api.GetPerson));
	mux.Handle("/api/project", 				CountApiHitsInSessionValues(api.FindProject));
	mux.Handle("/api/all_categories", 		CountApiHitsInSessionValues(api.GetAllCategories));
	mux.Handle("/api/owner", 				CountApiHitsInSessionValues(api.GetOwner));
	mux.Handle("/api/sessions", 			AuthorizeIfOwner(api.PrintAllUserSessions));
	mux.Handle("/api/search", 				(api.ProjectsSearch));
	mux.Handle("/api/anon_mails", 			AuthorizeIfHasSession(CountApiHitsInSessionValues(mails.SendAnonymousMail)));
	mux.Handle("/api/project_github_syncup",AuthorizeIfOwner(api.SyncProjectFromGithubRepository));

	// setup database connection
	data.Db, _ = sql.Open("sqlite3", config.GetGlobalConfig().GetSqlite3DatabaseFile())
	defer data.Db.Close()
	data.InitializeSchema()

	// setup and initialize search index, and insert all projects
	searchindex.InitProjectSearchIndex(config.GetGlobalConfig().GetBleveIndexFile());
	if(config.GetGlobalConfig().Recreate_search_index) {
		go searchindex.InsertAllProjectsInSearchIndex();
	}

	// set up session store
	ownerSessionId := ""
	if(config.GetGlobalConfig().Create_user_sessions) {
		fmt.Println("Initializing SessionStore");
		session.InitGlobalSessionStore("r_sess_id", 31 * 24 * time.Hour)
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
