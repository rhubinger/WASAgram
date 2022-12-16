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
	if !schemes.ValidUserId(uid) {
		rt.baseLogger.Error("UserId (uid) invalid")
		w.WriteHeader(http.StatusBadRequest)
		return
	} else if userExists, err := rt.db.UserExists(uid); err != nil {
		rt.baseLogger.Error("GetBanned: Error while checking for user in db")
		w.WriteHeader(http.StatusInternalServerError)
		return
	} else if userExists {
		rt.baseLogger.Error("GetBanned: User doesn't exist in db")
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// Get list of banned users by uid
	banned, err := rt.db.GetBanned(uid)
	if err != nil {
		rt.baseLogger.WithError(err).Error("GetBanned: failed to get banned from db")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	// Shorten if too many banned
	if len(banned) > 1000 {
		banned = banned[0:1000]
	}

	// Send the response
	var response = schemes.UserList{Length: len(banned), Users: banned}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("content-type", "application/json")
	_ = json.NewEncoder(w).Encode(response)
}

func (rt *_router) GetBannedCount(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// Parse parameters
	uid := ps.ByName("uid")
	if !schemes.ValidUserId(uid) {
		rt.baseLogger.Error("UserId (uid) invalid")
		w.WriteHeader(http.StatusBadRequest)
		return
	} else if userExists, err := rt.db.UserExists(uid); err != nil {
		rt.baseLogger.Error("GetBannedCount: Error while checking for user in db")
		w.WriteHeader(http.StatusInternalServerError)
		return
	} else if userExists {
		rt.baseLogger.Error("GetBannedCount: User doesn't exist in db")
		w.WriteHeader(http.StatusNotFound)
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
	w.WriteHeader(http.StatusOK)
	w.Header().Set("content-type", "application/json")
	_ = json.NewEncoder(w).Encode(response)
}

func (rt *_router) Ban(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// Parse parameters
	uid := ps.ByName("uid")
	if !schemes.ValidUserId(uid) {
		rt.baseLogger.Error("UserId (uid) invalid")
		w.WriteHeader(http.StatusBadRequest)
		return
	} else if userExists, err := rt.db.UserExists(uid); err != nil {
		rt.baseLogger.Error("Ban: Error while checking for user in db")
		w.WriteHeader(http.StatusInternalServerError)
		return
	} else if userExists {
		rt.baseLogger.Error("Ban: User doesn't exist in db")
		w.WriteHeader(http.StatusNotFound)
		return
	}
	bid := ps.ByName("bid")
	if !schemes.ValidUserId(bid) {
		rt.baseLogger.Error("BannedId (bid) invalid")
		w.WriteHeader(http.StatusBadRequest)
		return
	} else if userExists, err := rt.db.UserExists(bid); err != nil {
		rt.baseLogger.Error("Ban: Error while checking for user in db")
		w.WriteHeader(http.StatusInternalServerError)
		return
	} else if userExists {
		rt.baseLogger.Error("Ban: User doesn't exist in db")
		w.WriteHeader(http.StatusNotFound)
		return
	}

	if uid == bid {
		rt.baseLogger.Error("Ban: Users cant't ban themselves")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Check whether ban allready exists and if so return
	if banExists, err := rt.db.BanExists(uid, bid); err != nil {
		rt.baseLogger.Error("Ban: Error while checking for ban in db")
		w.WriteHeader(http.StatusInternalServerError)
		return
	} else if banExists {
		w.WriteHeader(http.StatusNoContent)
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
	if !schemes.ValidUserId(uid) {
		rt.baseLogger.Error("UserId (uid) invalid")
		w.WriteHeader(http.StatusBadRequest)
		return
	} else if userExists, err := rt.db.UserExists(uid); err != nil {
		rt.baseLogger.Error("Unban: Error while checking for user in db")
		w.WriteHeader(http.StatusInternalServerError)
		return
	} else if userExists {
		rt.baseLogger.Error("Unban: User doesn't exist in db")
		w.WriteHeader(http.StatusNotFound)
		return
	}
	bid := ps.ByName("bid")
	if !schemes.ValidUserId(bid) {
		rt.baseLogger.Error("BannedId (bid) invalid")
		w.WriteHeader(http.StatusBadRequest)
		return
	} else if userExists, err := rt.db.UserExists(uid); err != nil {
		rt.baseLogger.Error("Unban: Error while checking for user in db")
		w.WriteHeader(http.StatusInternalServerError)
		return
	} else if userExists {
		rt.baseLogger.Error("Unban: User doesn't exist in db")
		w.WriteHeader(http.StatusNotFound)
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
