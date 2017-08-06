package app

import (
	"bytes"
	"io"
	"namecardscanner/core"
	"namecardscanner/middleware"
	"net/http"
)

func (app *App) aliveCheck(w http.ResponseWriter, r *http.Request) {
	middleware.RespondWithJSON(w, http.StatusOK, struct{ Healthy bool }{true})
}

func (app *App) detectText(w http.ResponseWriter, r *http.Request) {
	var buf bytes.Buffer
	defer buf.Reset()

	file, _, err := r.FormFile("file")
	if err != nil {
		middleware.RespondWithJSON(w, http.StatusBadRequest, err.Error())
	}
	defer file.Close()
	io.Copy(&buf, file)

	response := core.DetectText(buf)
	if response.Success {
		middleware.RespondWithJSON(w, http.StatusOK, response)
	} else {
		middleware.RespondWithJSON(w, http.StatusBadRequest, response)
	}
}
