package api

import (
	"encoding/json"
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
	if identifier == "" {
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
	w.Header().Set("content-type", "application/json")
	_ = json.NewEncoder(w).Encode(response)
}

func (rt *_router) SearchUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// Parse Parameters
	searchString := r.URL.Query().Get("searchString")
	if searchString == "" {
		rt.baseLogger.Error("Failed to parse SearchString")
		w.WriteHeader(http.StatusBadRequest)
		return
	} else if !ValidSearchString(searchString) {
		rt.baseLogger.Error("SearchString is formated incorrectly")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	searchType := r.URL.Query().Get("type")

	// Search user in db
	users, err := rt.db.SearchUser(searchString, searchType)
	if err != nil {
		rt.baseLogger.WithError(err).Error("SearchUser: no users found in db")
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// Send the response
	w.Header().Set("content-type", "application/json")
	_ = json.NewEncoder(w).Encode(users)
}

func (rt *_router) GetUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// Parse Parameters
	uid := ps.ByName("uid")
	if !ValidUid(uid) {
		rt.baseLogger.Error("UserId (uid) invalid")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Get user from db by uid
	user, err := rt.db.GetUser(uid)
	if err != nil {
		rt.baseLogger.WithError(err).Error("GetUser: Failed to get User from db")
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// Send the response
	w.Header().Set("content-type", "application/json")
	_ = json.NewEncoder(w).Encode(user)
}

func (rt *_router) GetPosts(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	uid := ps.ByName("uid")
	if !ValidUid(uid) {
		rt.baseLogger.Error("UserId (uid) invalid")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Get the profiles posts in reverse chronological order

	// Send the response
	var length = 3
	var contentStream = make([]schemes.Post, 0, length)
	for i := 0; i < length; i++ {
		var u = schemes.User{UserId: "uid", Name: "Konrad Zuse", Posts: 5, Followers: 1783, Followed: 1}
		var p = schemes.Post{Poster: u, DateTime: "10-12-2022", Caption: "caption", PictureId: "pid", Likes: 3, Comments: 4}
		contentStream = append(contentStream, p)
	}
	w.Header().Set("content-type", "application/json")
	_ = json.NewEncoder(w).Encode(contentStream)
}

func (rt *_router) GetPostCount(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// Parse Parameters
	uid := ps.ByName("uid")
	if !ValidUid(uid) {
		rt.baseLogger.Error("UserId (uid) invalid")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Get count of posts

	// Send the response
	var response = GetCountResult{Count: 4}
	w.Header().Set("content-type", "application/json")
	_ = json.NewEncoder(w).Encode(response)
}

func (rt *_router) GetStream(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// Parse Parameters
	uid := ps.ByName("uid")
	if !ValidUid(uid) {
		rt.baseLogger.Error("UserId (uid) invalid")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Get the users stream in reverse chronological order

	// Send the response
	var length = 3
	var contentStream = make([]schemes.Post, 0, length)
	for i := 0; i < length; i++ {
		var u = schemes.User{UserId: "uid", Name: "Konrad Zuse", Posts: 5, Followers: 1783, Followed: 1}
		var p = schemes.Post{Poster: u, DateTime: "10-12-2022", Caption: "caption", PictureId: "pid", Likes: 3, Comments: 4}
		contentStream = append(contentStream, p)
	}
	w.Header().Set("content-type", "application/json")
	_ = json.NewEncoder(w).Encode(contentStream)
}

func (rt *_router) ChangeUsername(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// Parse Parameters
	uid := ps.ByName("uid")
	if !ValidUid(uid) {
		rt.baseLogger.Error("UserId (uid) invalid")
		w.WriteHeader(http.StatusBadRequest)
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
	w.WriteHeader(http.StatusCreated)
}
