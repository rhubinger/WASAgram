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
	likes, err := rt.db.GetLikes(pid)
	if err != nil {
		rt.baseLogger.WithError(err).Error("GetLikes: failed to get likes from db")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Send the response
	var response = schemes.UserList{Length: len(likes), Users: likes}
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
	count, err := rt.db.GetLikeCount(pid)
	if err != nil {
		rt.baseLogger.WithError(err).Error("GetLikeCount: failed to get like Count from db")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Send the response
	var response = GetCountResult{Count: count}
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
	uid := ps.ByName("uid")
	if !ValidUid(uid) {
		rt.baseLogger.Error("UserId (uid) invalid")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Create new like in db
	err := rt.db.Like(pid, uid)
	if err != nil {
		rt.baseLogger.WithError(err).Error("LikePost: failed to insert like into db")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	// Increment posts like count
	err = rt.db.IncrementLikeCount(pid)
	if err != nil {
		rt.baseLogger.WithError(err).Error("LikePost: failed to update posts like count in db")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Send the response
	w.WriteHeader(http.StatusNoContent)
}

func (rt *_router) UnlikePost(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// Get post from db
	pid := ps.ByName("pid")
	if !ValidId(pid) {
		rt.baseLogger.Error("PostId (pid) invalid")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	uid := ps.ByName("uid")
	if !ValidUid(uid) {
		rt.baseLogger.Error("UserId (uid) invalid")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Delete like from db
	err := rt.db.Unlike(pid, uid)
	if err != nil {
		rt.baseLogger.WithError(err).Error("UnikePost: failed to delete like from db")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	// Decrement posts like count
	err = rt.db.DecrementLikeCount(pid)
	if err != nil {
		rt.baseLogger.WithError(err).Error("UnikePost: failed to update posts like count in db")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Send the response
	w.WriteHeader(http.StatusNoContent)
}
