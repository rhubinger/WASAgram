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
	if !schemes.ValidUserId(uid) {
		rt.baseLogger.Error("UserId (uid) invalid")
		w.WriteHeader(http.StatusBadRequest)
		return
	} else if userExists, err := rt.db.UserExists(uid); err != nil {
		rt.baseLogger.Error("GetFollowed: Error while checking for user in db")
		w.WriteHeader(http.StatusInternalServerError)
		return
	} else if !userExists {
		rt.baseLogger.Error("GetFollowed: User doesn't exist in db")
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// Authentification as not banned by the user with userId uid
	identifier, err := ParseIdentifier(r)
	if err != nil {
		rt.baseLogger.WithError(err).Error("GetFollowed: Failed to parse identifier from reques")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	authorized, err := rt.db.AuthorizeAsNotBanned(identifier, uid)
	if err != nil {
		rt.baseLogger.WithError(err).Error("GetFollowed: Error occured during authorization")
		w.WriteHeader(http.StatusInternalServerError)
		return
	} else if !authorized {
		rt.baseLogger.Error("GetFollowed: User unauthorized to access resource")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// Get list of followed
	followed, err := rt.db.GetFollowed(uid)
	if err != nil {
		rt.baseLogger.WithError(err).Error("GetFollowed: error while getting followed users from db")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	// Shorten if too many followed
	if len(followed) > 1000 {
		followed = followed[0:1000]
	}

	// Send the response
	var response = schemes.UserList{Length: len(followed), Users: followed}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("content-type", "application/json")
	_ = json.NewEncoder(w).Encode(response)
}

func (rt *_router) GetFollowedCount(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// Parse Parameters
	uid := ps.ByName("uid")
	if !schemes.ValidUserId(uid) {
		rt.baseLogger.Error("UserId (uid) invalid")
		w.WriteHeader(http.StatusBadRequest)
		return
	} else if userExists, err := rt.db.UserExists(uid); err != nil {
		rt.baseLogger.Error("GetFollowedCount: Error while checking for user in db")
		w.WriteHeader(http.StatusInternalServerError)
		return
	} else if !userExists {
		rt.baseLogger.Error("GetFollowedCount: User doesn't exist in db")
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// Authentification as not banned by the user with userId uid
	identifier, err := ParseIdentifier(r)
	if err != nil {
		rt.baseLogger.WithError(err).Error("GetFollowedCount: Failed to parse identifier from reques")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	authorized, err := rt.db.AuthorizeAsNotBanned(identifier, uid)
	if err != nil {
		rt.baseLogger.WithError(err).Error("GetFollowedCount: Error occured during authorization")
		w.WriteHeader(http.StatusInternalServerError)
		return
	} else if !authorized {
		rt.baseLogger.Error("GetFollowedCount: User unauthorized to access resource")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// Get number of followed
	count, err := rt.db.GetFollowedCount(uid)
	if err != nil {
		rt.baseLogger.WithError(err).Error("GetFollowedCount: error while getting followed count from db")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Send the response
	var response = GetCountResult{Count: count}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("content-type", "application/json")
	_ = json.NewEncoder(w).Encode(response)
}

func (rt *_router) GetFollowers(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// Parse Parameters
	uid := ps.ByName("uid")
	if !schemes.ValidUserId(uid) {
		rt.baseLogger.Error("UserId (uid) invalid")
		w.WriteHeader(http.StatusBadRequest)
		return
	} else if userExists, err := rt.db.UserExists(uid); err != nil {
		rt.baseLogger.Error("GetFollowers: Error while checking for user in db")
		w.WriteHeader(http.StatusInternalServerError)
		return
	} else if !userExists {
		rt.baseLogger.Error("GetFollowers: User doesn't exist in db")
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// Authentification as not banned by the user with userId uid
	identifier, err := ParseIdentifier(r)
	if err != nil {
		rt.baseLogger.WithError(err).Error("GetFollowers: Failed to parse identifier from reques")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	authorized, err := rt.db.AuthorizeAsNotBanned(identifier, uid)
	if err != nil {
		rt.baseLogger.WithError(err).Error("GetFollowers: Error occured during authorization")
		w.WriteHeader(http.StatusInternalServerError)
		return
	} else if !authorized {
		rt.baseLogger.Error("GetFollowers: User unauthorized to access resource")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// Get followers
	followers, err := rt.db.GetFollowers(uid)
	if err != nil {
		rt.baseLogger.WithError(err).Error("GetFollowers: error while getting followers from db")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	// Shorten if too many followers
	if len(followers) > 1000 {
		followers = followers[0:1000]
	}

	// Send the response
	var response = schemes.UserList{Length: len(followers), Users: followers}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("content-type", "application/json")
	_ = json.NewEncoder(w).Encode(response)
}

func (rt *_router) GetFollowerCount(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// Parse Parameters
	uid := ps.ByName("uid")
	if !schemes.ValidUserId(uid) {
		rt.baseLogger.Error("UserId (uid) invalid")
		w.WriteHeader(http.StatusBadRequest)
		return
	} else if userExists, err := rt.db.UserExists(uid); err != nil {
		rt.baseLogger.Error("GetFollowerCount: Error while checking for user in db")
		w.WriteHeader(http.StatusInternalServerError)
		return
	} else if !userExists {
		rt.baseLogger.Error("GetFollowerCount: User doesn't exist in db")
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// Authentification as not banned by the user with userId uid
	identifier, err := ParseIdentifier(r)
	if err != nil {
		rt.baseLogger.WithError(err).Error("GetFollowerCount: Failed to parse identifier from reques")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	authorized, err := rt.db.AuthorizeAsNotBanned(identifier, uid)
	if err != nil {
		rt.baseLogger.WithError(err).Error("GetFollowerCount: Error occured during authorization")
		w.WriteHeader(http.StatusInternalServerError)
		return
	} else if !authorized {
		rt.baseLogger.Error("GetFollowerCount: User unauthorized to access resource")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// Get count of followers
	count, err := rt.db.GetFollowerCount(uid)
	if err != nil {
		rt.baseLogger.WithError(err).Error("GetFollowerCount: error while getting follower count from db")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Send the response
	var response = GetCountResult{Count: count}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("content-type", "application/json")
	_ = json.NewEncoder(w).Encode(response)
}

func (rt *_router) Follow(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// Parse Parameters
	uid := ps.ByName("uid")
	if !schemes.ValidUserId(uid) {
		rt.baseLogger.Error("UserId (uid) invalid")
		w.WriteHeader(http.StatusBadRequest)
		return
	} else if userExists, err := rt.db.UserExists(uid); err != nil {
		rt.baseLogger.Error("Follow: Error while checking for user in db")
		w.WriteHeader(http.StatusInternalServerError)
		return
	} else if !userExists {
		rt.baseLogger.Error("Follow: User doesn't exist in db")
		w.WriteHeader(http.StatusNotFound)
		return
	}
	fid := ps.ByName("fid")
	if !schemes.ValidUserId(fid) {
		rt.baseLogger.Error("FollowerId (fid) invalid")
		w.WriteHeader(http.StatusBadRequest)
		return
	} else if userExists, err := rt.db.UserExists(fid); err != nil {
		rt.baseLogger.Error("Follow: Error while checking for user in db")
		w.WriteHeader(http.StatusInternalServerError)
		return
	} else if !userExists {
		rt.baseLogger.Error("Follow: User doesn't exist in db")
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// Check wheter users want's to follow himself
	if uid == fid {
		rt.baseLogger.Error("Follow: Users cant't follow themselves")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Authentification as user with userId fid
	identifier, err := ParseIdentifier(r)
	if err != nil {
		rt.baseLogger.WithError(err).Error("Follow: Failed to parse identifier from reques")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	authorized, err := rt.db.AuthorizeAsUser(identifier, fid)
	if err != nil {
		rt.baseLogger.WithError(err).Error("Follow: Error occured during authorization")
		w.WriteHeader(http.StatusInternalServerError)
		return
	} else if !authorized {
		rt.baseLogger.Error("Follow: User unauthorized to access resource")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// Check whether follow allready exists and if so return
	if followExists, err := rt.db.FollowExists(uid, fid); err != nil {
		rt.baseLogger.Error("Follow: Error while checking for follow in db")
		w.WriteHeader(http.StatusInternalServerError)
		return
	} else if followExists {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	// Insert into db
	err = rt.db.Follow(uid, fid)
	if err != nil {
		rt.baseLogger.WithError(err).Error("Follow: failed to insert follow into db")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	// Update following users followed count in db
	err = rt.db.IncrementFollowedCount(fid)
	if err != nil {
		rt.baseLogger.WithError(err).Error("Follow: failed to update following users followed count in db")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	// Update followed users followers count in db
	err = rt.db.IncrementFollowerCount(uid)
	if err != nil {
		rt.baseLogger.WithError(err).Error("Follow: failed to update followed users follower count in db")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Send the response
	w.WriteHeader(http.StatusNoContent)
}

func (rt *_router) Unfollow(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// Parse Parameters
	uid := ps.ByName("uid")
	if !schemes.ValidUserId(uid) {
		rt.baseLogger.Error("UserId (uid) invalid")
		w.WriteHeader(http.StatusBadRequest)
		return
	} else if userExists, err := rt.db.UserExists(uid); err != nil {
		rt.baseLogger.Error("Unfollow: Error while checking for user in db")
		w.WriteHeader(http.StatusInternalServerError)
		return
	} else if !userExists {
		rt.baseLogger.Error("Unfollow: User doesn't exist in db")
		w.WriteHeader(http.StatusNotFound)
		return
	}
	fid := ps.ByName("fid")
	if !schemes.ValidUserId(fid) {
		rt.baseLogger.Error("FollowerId (fid) invalid")
		w.WriteHeader(http.StatusBadRequest)
		return
	} else if userExists, err := rt.db.UserExists(fid); err != nil {
		rt.baseLogger.Error("Unfollow: Error while checking for user in db")
		w.WriteHeader(http.StatusInternalServerError)
		return
	} else if !userExists {
		rt.baseLogger.Error("Unfollow: User doesn't exist in db")
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// Authentification as user with userId fid
	identifier, err := ParseIdentifier(r)
	if err != nil {
		rt.baseLogger.WithError(err).Error("Unfollow: Failed to parse identifier from reques")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	authorized, err := rt.db.AuthorizeAsUser(identifier, fid)
	if err != nil {
		rt.baseLogger.WithError(err).Error("Unfollow: Error occured during authorization")
		w.WriteHeader(http.StatusInternalServerError)
		return
	} else if !authorized {
		rt.baseLogger.Error("Unfollow: User unauthorized to access resource")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// Insert into db
	err = rt.db.Unfollow(uid, fid)
	if err != nil {
		rt.baseLogger.WithError(err).Error("UnFollow: failed to delete follow from db")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	// Update unfollowing users followed count in db
	err = rt.db.DecrementFollowedCount(fid)
	if err != nil {
		rt.baseLogger.WithError(err).Error("Unfollow: failed to update unfollowing users followed count in db")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	// Update unfollowed users followers count in db
	err = rt.db.DecrementFollowerCount(uid)
	if err != nil {
		rt.baseLogger.WithError(err).Error("Unfollow: failed to update unfollowed users follower count in db")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Send the response
	w.WriteHeader(http.StatusNoContent)
}
