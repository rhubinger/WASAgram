package api

import (
	"encoding/json"
	"net/http"

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
	// If no create user (set posts, followers & followed to 0)
	// return identifier

	// Send the response
	var response = LoginResult{Identifier: "gocwRvLhDf8"}
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
	if searchType == "" {
		searchType = "both"
	}

	// Search user in db

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

func (rt *_router) GetUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// Parse Parameters
	uid := ps.ByName("uid")
	if !ValidUid(uid) {
		rt.baseLogger.Error("UserId (uid) invalid")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Get user from db by uid

	// Send the response
	var response = schemes.User{UserId: "uid", Name: "name", Posts: 5, Followers: 1, Followed: 0}
	w.Header().Set("content-type", "application/json")
	_ = json.NewEncoder(w).Encode(response)
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

	// Send the response
	w.WriteHeader(http.StatusCreated)
}
