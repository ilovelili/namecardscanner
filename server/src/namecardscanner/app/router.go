package app

// InitializeRoutes Init routes
func (app *App) InitializeRoutes() {
	// text
	app.Router.HandleFunc("/text", app.detectText).Methods("POST")

	// alive check
	app.Router.HandleFunc("/", app.aliveCheck).Methods("GET")
}
