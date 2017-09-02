package app

// InitializeRoutes Init routes
func (app *App) InitializeRoutes() {
	// text
	app.Router.HandleFunc("/text/filestream", app.detectTextByImageStream).Methods("POST")
	app.Router.HandleFunc("/text/base64", app.detectTextByBase64).Methods("POST")

	// alive check
	app.Router.HandleFunc("/", app.aliveCheck).Methods("GET")
}
