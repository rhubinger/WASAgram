package api

import (
	"database/sql"
	"errors"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/rhubinger/WASAgram/service/schemes"
)

func (rt *_router) GetPicture(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// Parse parameters
	pid := ps.ByName("pid")
	if !schemes.ValidId(pid) {
		rt.baseLogger.Error("PictureId (pid) invalid")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Get picture from db
	fileBytes, err := rt.db.GetPicture(pid)
	if errors.Is(err, sql.ErrNoRows) {
		rt.baseLogger.WithError(err).Error("GetPicture: Picture doesn't exist in db")
		w.WriteHeader(http.StatusNotFound)
		return
	} else if err != nil {
		rt.baseLogger.WithError(err).Error("GetPicture: Failed to get picture from db")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Send the response
	w.Header().Set("Content-Type", "image/png")
	_, err = w.Write(fileBytes)
	if err != nil {
		rt.baseLogger.WithError(err).Error("GetPicture: Failed to send picture to user")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
