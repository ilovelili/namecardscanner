package app

import (
	"bytes"
	"encoding/json"
	"io"
	"namecardscanner/core"
	"namecardscanner/middleware"
	"namecardscanner/model"
	"net/http"
)

func (app *App) aliveCheck(w http.ResponseWriter, r *http.Request) {
	middleware.RespondWithJSON(w, http.StatusOK, struct{ Healthy bool }{true})
}

// detectTextByImageStream Detect Text By Image Stream (POST)
func (app *App) detectTextByImageStream(w http.ResponseWriter, r *http.Request) {
	var buf bytes.Buffer
	defer buf.Reset()

	file, _, err := r.FormFile("file")
	if err != nil {
		middleware.RespondWithJSON(w, http.StatusBadRequest, err.Error())
	}
	defer file.Close()
	io.Copy(&buf, file)

	response := core.DetectTextByImageStream(buf)
	if response.Success {
		middleware.RespondWithJSON(w, http.StatusOK, response)
	} else {
		middleware.RespondWithJSON(w, http.StatusBadRequest, response)
	}
}

// detectTextByBase64 Detect Text By Base64 (POST)
func (app *App) detectTextByBase64(w http.ResponseWriter, r *http.Request) {
	var request model.Request

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&request); err != nil {
		middleware.RespondWithJSON(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	response := core.DetectTextByBase64(request.Content)
	if response.Success {
		middleware.RespondWithJSON(w, http.StatusOK, response)
	} else {
		middleware.RespondWithJSON(w, http.StatusBadRequest, response)
	}
}
