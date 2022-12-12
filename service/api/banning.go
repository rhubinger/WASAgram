package api

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/rhubinger/WASAgram/service/schemes"
)

func (rt *_router) GetBanned(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// Parse parameters
	uid := ps.ByName("uid")
	if !ValidUid(uid) {
		rt.baseLogger.Error("UserId (uid) invalid")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Get list of banned users by uid

	// Send the response
	var length = 3
	var users = make([]schemes.User, 0, length)
	for i := 0; i < length; i++ {
		var u = schemes.User{UserId: "uid", Name: "Konrad Zuse", Posts: 5, Followers: 1783, Followed: 1}
		users = append(users, u)
	}
	var response = schemes.UserList{Length: length, Users: users}
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

	// Send the response
	var response = GetCountResult{Count: 5}
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

	// Send the response
	w.WriteHeader(http.StatusNoContent)
}
