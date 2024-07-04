package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/cloudykit/jet"
)

type application struct {
	appName string
	server  server
	debug   bool
	errLog  *log.Logger
	infoLog *log.Logger
	view    *jet.Set
	session *scs.SessionManager
}

type server struct {
	host string
	port string
	url  string
}

func main() {

	server := server{
		host: "localhost",
		port: "8009",
		url:  "http://localhost:8009",
	}

	// init application
	app := &application{
		server:  server,
		appName: "HNews",
		debug:   true,
		infoLog: log.New(os.Stdout, "INFO\t", log.Ltime|log.Ldate|log.Lshortfile),
		errLog:  log.New(os.Stderr, "ERROR\t", log.Ltime|log.Ldate|log.Llongfile),
	}

	//init jet template
	// if app.debug {
	// 	app.view = jet.NewSet(jet.NewOSFileSystemLoader("./views"), jet.IndevelopmentMode())
	// } else {
	// 	app.view = jet.NewSet(jet.NewOSFileSystemLoader("./views"))
	// }

	if err := app.listenAndServer(); err != nil {
		log.Fatal(err)
	}

	// init session
	app.session = scs.New()
	app.session.Lifetime = 24 * time.Hour
	app.session.Cookie.Persist = true
	app.session.Cookie.Name = app.appName
	app.session.Cookie.Domain = app.server.host
	app.session.Cookie.SameSite = http.SameSiteStrictMode
}
