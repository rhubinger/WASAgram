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
	} else if pictureExists, err := rt.db.PictureExists(pid); err != nil {
		rt.baseLogger.Error("GetPicture: Error while checking for picture in db")
		w.WriteHeader(http.StatusInternalServerError)
		return
	} else if !pictureExists {
		rt.baseLogger.Error("GetPicture: Picture doesn't exist in db")
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// Authentification as not banned by the user with userId post.UserId
	identifier, err := ParseIdentifier(r)
	if err != nil {
		rt.baseLogger.WithError(err).Error("GetPicture: Failed to parse identifier from reques")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	post, err := rt.db.GetPostByPictureId(pid)
	if err != nil {
		rt.baseLogger.WithError(err).Error("GetPicture: Failed to get post from db")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	authorized, err := rt.db.AuthorizeAsNotBanned(identifier, post.UserId)
	if err != nil {
		rt.baseLogger.WithError(err).Error("GetPicture: Error occured during authorization")
		w.WriteHeader(http.StatusInternalServerError)
		return
	} else if !authorized {
		rt.baseLogger.Error("GetPicture: User unauthorized to access resource")
		w.WriteHeader(http.StatusUnauthorized)
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
}
