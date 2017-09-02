package app

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

// App They call me God object. So I think I am very cool
type App struct {
	Router         *mux.Router
	Request        *http.Request
	ResponseWriter http.ResponseWriter
}

// Initialize init the app
func (app *App) Initialize() {
	// set up new router
	app.Router = mux.NewRouter()
	// init routes
	app.InitializeRoutes()
}

// Run ListenAndServe
func (app *App) Run(addr string) {
	handler := cors.New(cors.Options{
		// AllowedOrigins:   []string{"http://localhost:4200", "http://localhost:4400", "http://46.101.141.88"},
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
	}).Handler(app.Router)

	//handler := cors.Default().Handler(app.Router)
	log.Fatal(http.ListenAndServe(addr, handler))
}
