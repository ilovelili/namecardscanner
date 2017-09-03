package app

import "net/http"

// InitializeRoutes Init routes
func (app *App) InitializeRoutes() {
	// text
	app.Router.HandleFunc("/text/filestream", app.detectTextByImageStream).Methods("POST")
	app.Router.HandleFunc("/text/base64", app.detectTextByBase64).Methods("POST")

	// alive check
	app.Router.HandleFunc("/", app.aliveCheck).Methods("GET")

	// apk download
	app.Router.PathPrefix("/apk/").Handler(http.StripPrefix("/apk/", http.FileServer(http.Dir("./download/"))))
}
