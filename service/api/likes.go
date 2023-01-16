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
	if !schemes.ValidId(pid) {
		rt.baseLogger.Error("PostId (pid) invalid")
		w.WriteHeader(http.StatusBadRequest)
		return
	} else if postExists, err := rt.db.PostExists(pid); err != nil {
		rt.baseLogger.Error("GetLikes: Error while checking for post in db")
		w.WriteHeader(http.StatusInternalServerError)
		return
	} else if !postExists {
		rt.baseLogger.Error("GetLikes: Post doesn't exist in db")
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// Authentification as not banned by the user with userId post.UserId
	identifier, err := ParseIdentifier(r)
	if err != nil {
		rt.baseLogger.WithError(err).Error("GetLikes: Failed to parse identifier from reques")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	post, err := rt.db.GetPost(pid)
	if err != nil {
		rt.baseLogger.WithError(err).Error("GetLikes: Failed to get post from db")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	authorized, err := rt.db.AuthorizeAsNotBanned(identifier, post.UserId)
	if err != nil {
		rt.baseLogger.WithError(err).Error("GetLikes: Error occured during authorization")
		w.WriteHeader(http.StatusInternalServerError)
		return
	} else if !authorized {
		rt.baseLogger.Error("GetLikes: User unauthorized to access resource")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// Get list of likes
	likes, err := rt.db.GetLikes(pid)
	if err != nil {
		rt.baseLogger.WithError(err).Error("GetLikes: failed to get likes from db")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	// Shorten if too many followed
	if len(likes) > 1000 {
		likes = likes[0:1000]
	}

	// Send the response
	var response = schemes.UserList{Length: len(likes), Users: likes}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("content-type", "application/json")
	_ = json.NewEncoder(w).Encode(response)
}

func (rt *_router) GetLikeCount(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// Get post from db
	pid := ps.ByName("pid")
	if !schemes.ValidId(pid) {
		rt.baseLogger.Error("PostId (pid) invalid")
		w.WriteHeader(http.StatusBadRequest)
		return
	} else if postExists, err := rt.db.PostExists(pid); err != nil {
		rt.baseLogger.Error("GetLikeCount: Error while checking for post in db")
		w.WriteHeader(http.StatusInternalServerError)
		return
	} else if !postExists {
		rt.baseLogger.Error("GetLikeCount: Post doesn't exist in db")
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// Authentification as not banned by the user with userId post.UserId
	identifier, err := ParseIdentifier(r)
	if err != nil {
		rt.baseLogger.WithError(err).Error("GetLikeCount: Failed to parse identifier from reques")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	post, err := rt.db.GetPost(pid)
	if err != nil {
		rt.baseLogger.WithError(err).Error("GetLikeCount: Failed to get post from db")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	authorized, err := rt.db.AuthorizeAsNotBanned(identifier, post.UserId)
	if err != nil {
		rt.baseLogger.WithError(err).Error("GetLikeCount: Error occured during authorization")
		w.WriteHeader(http.StatusInternalServerError)
		return
	} else if !authorized {
		rt.baseLogger.Error("GetLikeCount: User unauthorized to access resource")
		w.WriteHeader(http.StatusUnauthorized)
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
	w.WriteHeader(http.StatusOK)
	w.Header().Set("content-type", "application/json")
	_ = json.NewEncoder(w).Encode(response)
}

func (rt *_router) hasLikedPost(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// Get post from db
	pid := ps.ByName("pid")
	if !schemes.ValidId(pid) {
		rt.baseLogger.Error("PostId (pid) invalid")
		w.WriteHeader(http.StatusBadRequest)
		return
	} else if postExists, err := rt.db.PostExists(pid); err != nil {
		rt.baseLogger.Error("hasLikedPost: Error while checking for post in db")
		w.WriteHeader(http.StatusInternalServerError)
		return
	} else if !postExists {
		rt.baseLogger.Error("hasLikedPost: Post doesn't exist in db")
		w.WriteHeader(http.StatusNotFound)
		return
	}
	uid := ps.ByName("uid")
	if !schemes.ValidUserId(uid) {
		rt.baseLogger.Error("UserId (uid) invalid")
		w.WriteHeader(http.StatusBadRequest)
		return
	} else if userExists, err := rt.db.UserExists(uid); err != nil {
		rt.baseLogger.Error("hasLikedPost: Error while checking for user in db")
		w.WriteHeader(http.StatusInternalServerError)
		return
	} else if !userExists {
		rt.baseLogger.Error("hasLikedPost: User doesn't exist in db")
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// Authentification as not banned by the user with userId post.UserId
	identifier, err := ParseIdentifier(r)
	if err != nil {
		rt.baseLogger.WithError(err).Error("hasLikedPost: Failed to parse identifier from reques")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	post, err := rt.db.GetPost(pid)
	if err != nil {
		rt.baseLogger.WithError(err).Error("hasLikedPost: Failed to get post from db")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	authorized, err := rt.db.AuthorizeAsNotBanned(identifier, post.UserId)
	if err != nil {
		rt.baseLogger.WithError(err).Error("hasLikedPost: Error occured during authorization")
		w.WriteHeader(http.StatusInternalServerError)
		return
	} else if !authorized {
		rt.baseLogger.Error("hasLikedPost: User unauthorized to access resource")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// Check whether like exists in db
	likeExists, err := rt.db.LikeExists(pid, uid)
	if err != nil {
		rt.baseLogger.Error("hasLikedPost: Error while checking for like in db")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Send the response
	var response = struct {hasLiked bool}{likeExists}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("content-type", "application/json")
	_ = json.NewEncoder(w).Encode(response)
}

func (rt *_router) LikePost(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// Get post from db
	pid := ps.ByName("pid")
	if !schemes.ValidId(pid) {
		rt.baseLogger.Error("PostId (pid) invalid")
		w.WriteHeader(http.StatusBadRequest)
		return
	} else if postExists, err := rt.db.PostExists(pid); err != nil {
		rt.baseLogger.Error("LikePost: Error while checking for post in db")
		w.WriteHeader(http.StatusInternalServerError)
		return
	} else if !postExists {
		rt.baseLogger.Error("LikePost: Post doesn't exist in db")
		w.WriteHeader(http.StatusNotFound)
		return
	}
	uid := ps.ByName("uid")
	if !schemes.ValidUserId(uid) {
		rt.baseLogger.Error("UserId (uid) invalid")
		w.WriteHeader(http.StatusBadRequest)
		return
	} else if userExists, err := rt.db.UserExists(uid); err != nil {
		rt.baseLogger.Error("LikePost: Error while checking for user in db")
		w.WriteHeader(http.StatusInternalServerError)
		return
	} else if !userExists {
		rt.baseLogger.Error("LikePost: User doesn't exist in db")
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// Authentification as user with userId uid
	identifier, err := ParseIdentifier(r)
	if err != nil {
		rt.baseLogger.WithError(err).Error("LikePost: Failed to parse identifier from reques")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	authorized, err := rt.db.AuthorizeAsUser(identifier, uid)
	if err != nil {
		rt.baseLogger.WithError(err).Error("LikePost: Error occured during authorization")
		w.WriteHeader(http.StatusInternalServerError)
		return
	} else if !authorized {
		rt.baseLogger.Error("LikePost: User unauthorized to access resource")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// Check whether like allready exists and if so return
	if likeExists, err := rt.db.LikeExists(pid, uid); err != nil {
		rt.baseLogger.Error("Like: Error while checking for like in db")
		w.WriteHeader(http.StatusInternalServerError)
		return
	} else if likeExists {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	// Create new like in db
	err = rt.db.Like(pid, uid)
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
	if !schemes.ValidId(pid) {
		rt.baseLogger.Error("PostId (pid) invalid")
		w.WriteHeader(http.StatusBadRequest)
		return
	} else if postExists, err := rt.db.PostExists(pid); err != nil {
		rt.baseLogger.Error("UnlikePost: Error while checking for post in db")
		w.WriteHeader(http.StatusInternalServerError)
		return
	} else if !postExists {
		rt.baseLogger.Error("UnlikePost: Post doesn't exist in db")
		w.WriteHeader(http.StatusNotFound)
		return
	}
	uid := ps.ByName("uid")
	if !schemes.ValidUserId(uid) {
		rt.baseLogger.Error("UserId (uid) invalid")
		w.WriteHeader(http.StatusBadRequest)
		return
	} else if userExists, err := rt.db.UserExists(uid); err != nil {
		rt.baseLogger.Error("UnlikePost: Error while checking for user in db")
		w.WriteHeader(http.StatusInternalServerError)
		return
	} else if !userExists {
		rt.baseLogger.Error("UnlikePost: User doesn't exist in db")
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// Authentification as user with userId uid
	identifier, err := ParseIdentifier(r)
	if err != nil {
		rt.baseLogger.WithError(err).Error("UnlikePost: Failed to parse identifier from reques")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	authorized, err := rt.db.AuthorizeAsUser(identifier, uid)
	if err != nil {
		rt.baseLogger.WithError(err).Error("UnlikePost: Error occured during authorization")
		w.WriteHeader(http.StatusInternalServerError)
		return
	} else if !authorized {
		rt.baseLogger.Error("UnlikePost: User unauthorized to access resource")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// Delete like from db
	err = rt.db.Unlike(pid, uid)
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
