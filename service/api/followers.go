package api

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/rhubinger/WASAgram/service/schemes"
)

// Prevent self folow
func (rt *_router) GetFollowed(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// Parse Parameters
	uid := ps.ByName("uid")
	if !ValidUid(uid) {
		rt.baseLogger.Error("UserId (uid) invalid")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Get list of followed
	followers, err := rt.db.GetFollowed(uid)
	if err != nil {
		rt.baseLogger.WithError(err).Error("GetFollowed: error while getting followed users from db")
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// Send the response
	var response = schemes.UserList{Length: len(followers), Users: followers}
	w.Header().Set("content-type", "application/json")
	_ = json.NewEncoder(w).Encode(response)
}

func (rt *_router) GetFollowedCount(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// Parse Parameters
	uid := ps.ByName("uid")
	if !ValidUid(uid) {
		rt.baseLogger.Error("UserId (uid) invalid")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Get number of followed
	count, err := rt.db.GetFollowedCount(uid)
	if err != nil {
		rt.baseLogger.WithError(err).Error("GetFollowedCount: error while getting followed count from db")
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// Send the response
	var response = GetCountResult{Count: count}
	w.Header().Set("content-type", "application/json")
	_ = json.NewEncoder(w).Encode(response)
}

func (rt *_router) GetFollowers(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// Parse Parameters
	uid := ps.ByName("uid")
	if !ValidUid(uid) {
		rt.baseLogger.Error("UserId (uid) invalid")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Get followers
	followers, err := rt.db.GetFollowers(uid)
	if err != nil {
		rt.baseLogger.WithError(err).Error("GetFollowers: error while getting followers from db")
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// Send the response
	var response = schemes.UserList{Length: len(followers), Users: followers}
	w.Header().Set("content-type", "application/json")
	_ = json.NewEncoder(w).Encode(response)
}

func (rt *_router) GetFollowerCount(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// Parse Parameters
	uid := ps.ByName("uid")
	if !ValidUid(uid) {
		rt.baseLogger.Error("UserId (uid) invalid")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Get count of followers
	count, err := rt.db.GetFollowerCount(uid)
	if err != nil {
		rt.baseLogger.WithError(err).Error("GetFollowerCount: error while getting follower count from db")
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// Send the response
	var response = GetCountResult{Count: count}
	w.Header().Set("content-type", "application/json")
	_ = json.NewEncoder(w).Encode(response)
}

func (rt *_router) Follow(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// Parse Parameters
	uid := ps.ByName("uid")
	if !ValidUid(uid) {
		rt.baseLogger.Error("UserId (uid) invalid")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	fid := ps.ByName("fid")
	if !ValidUid(fid) {
		rt.baseLogger.Error("FollowerId (fid) invalid")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Insert into db
	err := rt.db.Follow(uid, fid)
	if err != nil {
		rt.baseLogger.WithError(err).Error("Follow: failed to insert follow into db")
		w.WriteHeader(http.StatusNotFound)
		return
	}
	// Update following users followed count in db
	err = rt.db.IncrementFollowedCount(fid)
	if err != nil {
		rt.baseLogger.WithError(err).Error("Follow: failed to update following users followed count in db")
		w.WriteHeader(http.StatusNotFound)
		return
	}
	// Update followed users followers count in db
	err = rt.db.IncrementFollowerCount(uid)
	if err != nil {
		rt.baseLogger.WithError(err).Error("Follow: failed to update followed users follower count in db")
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// Send the response
	w.WriteHeader(http.StatusNoContent)
}

func (rt *_router) Unfollow(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// Parse Parameters
	uid := ps.ByName("uid")
	if !ValidUid(uid) {
		rt.baseLogger.Error("UserId (uid) invalid")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	fid := ps.ByName("fid")
	if !ValidUid(fid) {
		rt.baseLogger.Error("FollowerId (fid) invalid")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Insert into db
	err := rt.db.Unfollow(uid, fid)
	if err != nil {
		rt.baseLogger.WithError(err).Error("Follow: failed to delete follow from db")
		w.WriteHeader(http.StatusNotFound)
		return
	}
	// Update unfollowing users followed count in db
	err = rt.db.DecrementFollowedCount(fid)
	if err != nil {
		rt.baseLogger.WithError(err).Error("Follow: failed to update unfollowing users followed count in db")
		w.WriteHeader(http.StatusNotFound)
		return
	}
	// Update unfollowed users followers count in db
	err = rt.db.DecrementFollowerCount(uid)
	if err != nil {
		rt.baseLogger.WithError(err).Error("Follow: failed to update unfollowed users follower count in db")
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// Send the response
	w.WriteHeader(http.StatusNoContent)
}
