package api

import (
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"
	"strings"

	"github.com/julienschmidt/httprouter"
	"github.com/rhubinger/WASAgram/service/schemes"
)

func (rt *_router) Login(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// Parse request body
	var request LoginRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	_ = r.Body.Close()
	if err != nil {
		rt.baseLogger.WithError(err).Error("Login: Failed to parse request body")
		w.WriteHeader(http.StatusBadRequest)
		return
	} else if !request.Valid() {
		rt.baseLogger.Error("Login: Request Body invalid")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Check whether user already exists
	identifier, err := rt.db.GetIdentifier(request.UserId)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		rt.baseLogger.WithError(err).Error("Login: Failed to look up whether user exists db")
		w.WriteHeader(http.StatusInternalServerError)
		return
	} else if identifier == "" {
		// If user doesn't exist create user
		name := strings.Replace(request.UserId, "@", "", -1)
		name = strings.Replace(name, "_", " ", -1)

		var user = schemes.User{UserId: request.UserId, Name: name, Posts: 0, Followers: 0, Followed: 0}
		identifier, err = rt.db.InsertUser(user)
		if err != nil {
			rt.baseLogger.WithError(err).Error("Login: Failed to insert new user into db")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}

	// Send response (identifier)
	var response = LoginResult{Identifier: identifier}
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("content-type", "application/json")
	_ = json.NewEncoder(w).Encode(response)
}

func (rt *_router) SearchUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// Parse Parameters
	searchString := r.URL.Query().Get("searchString")
	if searchString == "" {
		rt.baseLogger.Error("SearchUser: Failed to parse SearchString")
		w.WriteHeader(http.StatusBadRequest)
		return
	} else if !schemes.ValidSearchString(searchString) {
		rt.baseLogger.Error("SearchUser: SearchString is formated incorrectly")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	uid := r.URL.Query().Get("uid")
	if uid == "" {
		rt.baseLogger.Error("SearchUser: Failed to parse uid")
		w.WriteHeader(http.StatusBadRequest)
		return
	} else if !schemes.ValidUserId(uid) {
		rt.baseLogger.Error("SearchUser: uid is formated incorrectly")
		w.WriteHeader(http.StatusBadRequest)
		return
	} else if userExists, err := rt.db.UserExists(uid); err != nil {
		rt.baseLogger.Error("SearchUser: Error while checking for user in db")
		w.WriteHeader(http.StatusInternalServerError)
		return
	} else if userExists {
		rt.baseLogger.Error("SearchUser: User doesn't exist in db")
		w.WriteHeader(http.StatusNotFound)
		return
	}
	// Search user in db
	users, err := rt.db.SearchUser(searchString, uid)
	if err != nil {
		rt.baseLogger.WithError(err).Error("SearchUser: couldn't get search results from db")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Send the response
	w.WriteHeader(http.StatusOK)
	w.Header().Set("content-type", "application/json")
	_ = json.NewEncoder(w).Encode(users)
}

func (rt *_router) GetUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// Parse Parameters
	uid := ps.ByName("uid")
	if !schemes.ValidUserId(uid) {
		rt.baseLogger.Error("UserId (uid) invalid")
		w.WriteHeader(http.StatusBadRequest)
		return
	} else if userExists, err := rt.db.UserExists(uid); err != nil {
		rt.baseLogger.Error("GetUser: Error while checking for user in db")
		w.WriteHeader(http.StatusInternalServerError)
		return
	} else if userExists {
		rt.baseLogger.Error("GetUser: User doesn't exist in db")
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// Get user from db by uid
	user, err := rt.db.GetUser(uid)
	if err != nil {
		rt.baseLogger.WithError(err).Error("GetUser: Failed to get User from db")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Send the response
	w.WriteHeader(http.StatusOK)
	w.Header().Set("content-type", "application/json")
	_ = json.NewEncoder(w).Encode(user)
}

func (rt *_router) GetPosts(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// Parse parameters
	uid := ps.ByName("uid")
	if !schemes.ValidUserId(uid) {
		rt.baseLogger.Error("UserId (uid) invalid")
		w.WriteHeader(http.StatusBadRequest)
		return
	} else if userExists, err := rt.db.UserExists(uid); err != nil {
		rt.baseLogger.Error("GetPosts: Error while checking for user in db")
		w.WriteHeader(http.StatusInternalServerError)
		return
	} else if userExists {
		rt.baseLogger.Error("GetPosts: User doesn't exist in db")
		w.WriteHeader(http.StatusNotFound)
		return
	}

	dateTime := r.URL.Query().Get("dateTime")
	if dateTime == "" {
		rt.baseLogger.Error("SearchUser: Failed to parse dateTime")
		w.WriteHeader(http.StatusBadRequest)
		return
	} else if !schemes.ValidDatetime(dateTime) {
		rt.baseLogger.Error("GetPosts: dateTime is formated incorrectly")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Get the profiles posts in reverse chronological order
	posts, err := rt.db.GetPosts(uid, dateTime)
	if err != nil {
		rt.baseLogger.WithError(err).Error("GetPosts: error while getting users posts")
		w.WriteHeader(http.StatusNotFound)
		return
	}
	// Shorten if too many posts
	if len(posts) > 12 {
		posts = posts[0:12]
	}

	// Send the response
	w.WriteHeader(http.StatusOK)
	w.Header().Set("content-type", "application/json")
	_ = json.NewEncoder(w).Encode(posts)
}

func (rt *_router) GetPostCount(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// Parse Parameters
	uid := ps.ByName("uid")
	if !schemes.ValidUserId(uid) {
		rt.baseLogger.Error("UserId (uid) invalid")
		w.WriteHeader(http.StatusBadRequest)
		return
	} else if userExists, err := rt.db.UserExists(uid); err != nil {
		rt.baseLogger.Error("GetPostCount: Error while checking for user in db")
		w.WriteHeader(http.StatusInternalServerError)
		return
	} else if userExists {
		rt.baseLogger.Error("GetPostCount: User doesn't exist in db")
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// Get count of posts
	count, err := rt.db.GetPostCount(uid)
	if err != nil {
		rt.baseLogger.WithError(err).Error("GetPostCount: error while getting post count from db")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Send the response
	var response = GetCountResult{Count: count}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("content-type", "application/json")
	_ = json.NewEncoder(w).Encode(response)
}

func (rt *_router) GetStream(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// Parse Parameters
	uid := ps.ByName("uid")
	if !schemes.ValidUserId(uid) {
		rt.baseLogger.Error("UserId (uid) invalid")
		w.WriteHeader(http.StatusBadRequest)
		return
	} else if userExists, err := rt.db.UserExists(uid); err != nil {
		rt.baseLogger.Error("GetStream: Error while checking for user in db")
		w.WriteHeader(http.StatusInternalServerError)
		return
	} else if userExists {
		rt.baseLogger.Error("GetStream: User doesn't exist in db")
		w.WriteHeader(http.StatusNotFound)
		return
	}

	dateTime := r.URL.Query().Get("dateTime")
	if dateTime == "" {
		rt.baseLogger.Error("SearchUser: Failed to parse dateTime")
		w.WriteHeader(http.StatusBadRequest)
		return
	} else if !schemes.ValidDatetime(dateTime) {
		rt.baseLogger.Error("GetPosts: dateTime is formated incorrectly")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Get the users stream in reverse chronological order
	stream, err := rt.db.GetStream(uid, dateTime)
	if err != nil {
		rt.baseLogger.WithError(err).Error("GetStream: error while getting posts from db")
		w.WriteHeader(http.StatusNotFound)
		return
	}
	// Shorten if too many posts
	if len(stream) > 12 {
		stream = stream[0:12]
	}

	// Send the response
	w.WriteHeader(http.StatusOK)
	w.Header().Set("content-type", "application/json")
	_ = json.NewEncoder(w).Encode(stream)
}

func (rt *_router) ChangeUsername(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// Parse Parameters
	uid := ps.ByName("uid")
	if !schemes.ValidUserId(uid) {
		rt.baseLogger.Error("UserId (uid) invalid")
		w.WriteHeader(http.StatusBadRequest)
		return
	} else if userExists, err := rt.db.UserExists(uid); err != nil {
		rt.baseLogger.Error("ChangeUsername: Error while checking for user in db")
		w.WriteHeader(http.StatusInternalServerError)
		return
	} else if userExists {
		rt.baseLogger.Error("ChangeUsername: User doesn't exist in db")
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// Get request body
	var request ChangeNameRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	_ = r.Body.Close()
	if err != nil {
		rt.baseLogger.WithError(err).Error("ChangeUsername: Failed to parse request body")
		w.WriteHeader(http.StatusBadRequest)
		return
	} else if !request.Valid() {
		rt.baseLogger.Error("ChangeUsername: Request Body invalid")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Update/Change username in db
	err = rt.db.UpdateUsername(request.Name, uid)
	if err != nil {
		rt.baseLogger.WithError(err).Error("ChangeUsername: Failed to update new name in db")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Send the response
	w.WriteHeader(http.StatusNoContent)
}
