package api

import (
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

	// Get picture from db
	fileBytes, err := rt.db.GetPicture(pid)
	if err != nil {
		rt.baseLogger.WithError(err).Error("GetPicture: Failed to get picture from db")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Send the response
	w.Header().Set("Content-Type", "image/png")
	w.Write(fileBytes)
	return
}
