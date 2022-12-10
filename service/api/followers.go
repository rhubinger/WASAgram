package api

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) GetFollowed(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	uid := ps.ByName("uid")
	if len(uid) != 12 {
		//ctx.Logger.WithError(err).Error("enroll: error decoding JSON") TODO figure out how to use those
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Send the response
	var length = 3
	var users = make([]User, 0, length)
	for i := 0; i < length; i++ {
		var u = User{UserId: "uid", Name: "Konrad Zuse", Followers: 1783, Followed: 1}
		users = append(users, u)
	}

	var response = UserList{Length: length, Users: users}

	w.Header().Set("content-type", "application/json")
	_ = json.NewEncoder(w).Encode(response)
}

func (rt *_router) GetFollowedCount(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	uid := ps.ByName("uid")
	if len(uid) != 12 {
		//ctx.Logger.WithError(err).Error("enroll: error decoding JSON") TODO figure out how to use those
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Send the response
	var response = GetCountResult{Count: 3}
	w.Header().Set("content-type", "application/json")
	_ = json.NewEncoder(w).Encode(response)
}

func (rt *_router) GetFollowers(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	uid := ps.ByName("uid")
	if len(uid) != 12 {
		//ctx.Logger.WithError(err).Error("enroll: error decoding JSON") TODO figure out how to use those
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Send the response
	var length = 3
	var users = make([]User, 0, length)
	for i := 0; i < length; i++ {
		var u = User{UserId: "uid", Name: "Konrad Zuse", Followers: 1783, Followed: 1}
		users = append(users, u)
	}

	var response = UserList{Length: length, Users: users}

	w.Header().Set("content-type", "application/json")
	_ = json.NewEncoder(w).Encode(response)
}

func (rt *_router) GetFollowerCount(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	uid := ps.ByName("uid")
	if len(uid) != 12 {
		//ctx.Logger.WithError(err).Error("enroll: error decoding JSON") TODO figure out how to use those
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Send the response
	var response = GetCountResult{Count: 4}
	w.Header().Set("content-type", "application/json")
	_ = json.NewEncoder(w).Encode(response)
}

func (rt *_router) Follow(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	uid := ps.ByName("uid")
	if len(uid) != 12 {
		//ctx.Logger.WithError(err).Error("enroll: error decoding JSON") TODO figure out how to use those
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	fid := ps.ByName("fid")
	if len(fid) != 12 {
		//ctx.Logger.WithError(err).Error("enroll: error decoding JSON") TODO figure out how to use those
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Send the response
	w.WriteHeader(http.StatusNoContent)
}

func (rt *_router) Unfollow(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	uid := ps.ByName("uid")
	if len(uid) != 12 {
		//ctx.Logger.WithError(err).Error("enroll: error decoding JSON") TODO figure out how to use those
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	fid := ps.ByName("fid")
	if len(fid) != 12 {
		//ctx.Logger.WithError(err).Error("enroll: error decoding JSON") TODO figure out how to use those
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Send the response
	w.WriteHeader(http.StatusNoContent)
}