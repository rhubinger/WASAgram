package api

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/rhubinger/WASAgram/service/schemes"
)

func (rt *_router) GetLikes(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// Get post from db
	pid := ps.ByName("pid")
	if !ValidId(pid) {
		rt.baseLogger.Error("PostId (pid) invalid")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Get list of likes

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

func (rt *_router) GetLikeCount(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// Get post from db
	pid := ps.ByName("pid")
	if !ValidId(pid) {
		rt.baseLogger.Error("PostId (pid) invalid")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Get like count

	// Send the response
	var response = GetCountResult{Count: 6}
	w.Header().Set("content-type", "application/json")
	_ = json.NewEncoder(w).Encode(response)
}

func (rt *_router) LikePost(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// Get post from db
	pid := ps.ByName("pid")
	if !ValidId(pid) {
		rt.baseLogger.Error("PostId (pid) invalid")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	lid := ps.ByName("lid")
	if !ValidId(lid) {
		rt.baseLogger.Error("LikeId (lid) invalid")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Create new like in db
	// Increment posts like count

	// Send the response
	w.WriteHeader(http.StatusNoContent)
}

func (rt *_router) DeleteLike(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// Get post from db
	pid := ps.ByName("pid")
	if !ValidId(pid) {
		rt.baseLogger.Error("PostId (pid) invalid")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	lid := ps.ByName("lid")
	if !ValidId(lid) {
		rt.baseLogger.Error("LikeId (lid) invalid")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Delete like from db
	// Decrement posts like count

	// Send the response
	w.WriteHeader(http.StatusNoContent)
}
