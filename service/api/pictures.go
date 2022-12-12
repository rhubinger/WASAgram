package api

import (
	"io/ioutil"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) GetPicture(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// Parse parameters
	pid := ps.ByName("pid")
	if !ValidId(pid) {
		rt.baseLogger.Error("PostId (pid) invalid")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	fileBytes, err := ioutil.ReadFile("pictures/portrait_20.png")
	if err != nil {
		rt.baseLogger.WithError(err).Error("GetPicture: Failed to parse request body")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Get picture from db

	// Send the response
	w.Header().Set("Content-Type", "image/png")
	w.Write(fileBytes)
	return
}
