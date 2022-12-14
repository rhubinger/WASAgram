package api

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/rhubinger/WASAgram/service/schemes"
)

// Prevent self ban
func (rt *_router) GetBanned(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// Parse parameters
	uid := ps.ByName("uid")
	if !ValidUid(uid) {
		rt.baseLogger.Error("UserId (uid) invalid")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Get list of banned users by uid
	banned, err := rt.db.GetBanned(uid)
	if err != nil {
		rt.baseLogger.WithError(err).Error("GetBanned: failed to get banned from db")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Send the response
	var response = schemes.UserList{Length: len(banned), Users: banned}
	w.Header().Set("content-type", "application/json")
	_ = json.NewEncoder(w).Encode(response)
}

func (rt *_router) GetBannedCount(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// Parse parameters
	uid := ps.ByName("uid")
	if !ValidUid(uid) {
		rt.baseLogger.Error("UserId (uid) invalid")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Get count of banned account
	count, err := rt.db.GetBannedCount(uid)
	if err != nil {
		rt.baseLogger.WithError(err).Error("GetBannedCount: failed to get banned count from db")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Send the response
	var response = GetCountResult{Count: count}
	w.Header().Set("content-type", "application/json")
	_ = json.NewEncoder(w).Encode(response)
}

func (rt *_router) Ban(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// Parse parameters
	uid := ps.ByName("uid")
	if !ValidUid(uid) {
		rt.baseLogger.Error("UserId (uid) invalid")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	bid := ps.ByName("bid")
	if !ValidUid(bid) {
		rt.baseLogger.Error("BannedId (bid) invalid")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Insert ban in db
	err := rt.db.Ban(uid, bid)
	if err != nil {
		rt.baseLogger.WithError(err).Error("Ban: failed to insert ban into db")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Send the response
	w.WriteHeader(http.StatusNoContent)
}

func (rt *_router) Unban(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// Parse parameters
	uid := ps.ByName("uid")
	if !ValidUid(uid) {
		rt.baseLogger.Error("UserId (uid) invalid")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	bid := ps.ByName("bid")
	if !ValidUid(bid) {
		rt.baseLogger.Error("BannedId (bid) invalid")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Delete ban from db
	err := rt.db.Unban(uid, bid)
	if err != nil {
		rt.baseLogger.WithError(err).Error("Ban: failed to delete ban from db")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Send the response
	w.WriteHeader(http.StatusNoContent)
}
