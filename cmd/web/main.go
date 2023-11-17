package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/BuzzBumble/alwaysallin/config"
	"github.com/BuzzBumble/alwaysallin/pkg/handlers"
	"github.com/BuzzBumble/alwaysallin/pkg/render"

	"github.com/alexedwards/scs/v2"
)

const PORT = ":8080"
const USECACHE = false

var app config.AppConfig
var sm *scs.SessionManager

func main() {
	app.UseCache = USECACHE

	sm = scs.New()
	sm.Lifetime = 24 * time.Hour
	sm.Cookie.Persist = true
	sm.Cookie.SameSite = http.SameSiteLaxMode
	sm.Cookie.Secure = app.IsSecure

	app.SessionManager = sm

	render.SetConfig(&app)
	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatalln("Cannot create template cache:", err)
	}

	app.TemplateCache = tc

	//	http.HandleFunc("/", handlers.Repo.Home)
	//	http.HandleFunc("/about", handlers.Repo.About)

	fmt.Println(fmt.Sprintf("Listening on localhost:%s", PORT))

	srv := &http.Server{
		Addr:    PORT,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	log.Fatalln(err)
}
