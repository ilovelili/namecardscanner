package app

// InitializeRoutes Init routes
func (app *App) InitializeRoutes() {
	// text
	app.Router.HandleFunc("/text", app.detectTextByImageStream).Methods("POST")
	app.Router.HandleFunc("/text/{content}", app.detectTextByBase64).Methods("GET")

	// alive check
	app.Router.HandleFunc("/", app.aliveCheck).Methods("GET")
}
