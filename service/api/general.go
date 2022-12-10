package api

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) SearchUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	searchString := r.URL.Query().Get("searchString")
	if searchString == "" {
		//ctx.Logger.WithError(err).Error("enroll: error decoding JSON") TODO figure out how to use those
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	searchType := r.URL.Query().Get("type")
	if searchType == "" {
		searchType = "both"
	}

	// Send the response
	var length = 3
	var contentStream = make([]Post, 0, length)
	for i := 0; i < length; i++ {
		var u = User{UserId: "uid", Name: "Konrad Zuse", Followers: 1783, Followed: 1}
		var p = Post{Poster: u, DateTime: "10-12-2022", Caption: "caption", PictureId: "pid", Likes: 3, Comments: 4}
		contentStream = append(contentStream, p)
	}

	w.Header().Set("content-type", "application/json")
	_ = json.NewEncoder(w).Encode(contentStream)
}

func (rt *_router) GetUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	uid := ps.ByName("uid")
	if len(uid) != 12 {
		//ctx.Logger.WithError(err).Error("enroll: error decoding JSON") TODO figure out how to use those
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	var response = User{UserId: "uid", Name: "name", Followers: 1, Followed: 0}

	w.Header().Set("content-type", "application/json")
	_ = json.NewEncoder(w).Encode(response)
}

func (rt *_router) GetPosts(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	uid := ps.ByName("uid")
	if len(uid) != 12 {
		//ctx.Logger.WithError(err).Error("enroll: error decoding JSON") TODO figure out how to use those
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Send the response
	var length = 3
	var contentStream = make([]Post, 0, length)
	for i := 0; i < length; i++ {
		var u = User{UserId: "uid", Name: "Konrad Zuse", Followers: 1783, Followed: 1}
		var p = Post{Poster: u, DateTime: "10-12-2022", Caption: "caption", PictureId: "pid", Likes: 3, Comments: 4}
		contentStream = append(contentStream, p)
	}

	w.Header().Set("content-type", "application/json")
	_ = json.NewEncoder(w).Encode(contentStream)
}

func (rt *_router) GetStream(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	uid := ps.ByName("uid")
	if len(uid) != 12 {
		//ctx.Logger.WithError(err).Error("enroll: error decoding JSON") TODO figure out how to use those
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Send the response
	var length = 3
	var contentStream = make([]Post, 0, length)
	for i := 0; i < length; i++ {
		var u = User{UserId: "uid", Name: "Konrad Zuse", Followers: 1783, Followed: 1}
		var p = Post{Poster: u, DateTime: "10-12-2022", Caption: "caption", PictureId: "pid", Likes: 3, Comments: 4}
		contentStream = append(contentStream, p)
	}
	w.Header().Set("content-type", "application/json")
	_ = json.NewEncoder(w).Encode(contentStream)
}

func (rt *_router) ChangeUsername(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	uid := ps.ByName("uid")
	if len(uid) != 12 {
		//ctx.Logger.WithError(err).Error("enroll: error decoding JSON") TODO figure out how to use those
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var request ChangeNameRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	_ = r.Body.Close()
	if err != nil {
		//ctx.Logger.WithError(err).Error("enroll: error decoding JSON") TODO figure out how to use those
		w.WriteHeader(http.StatusBadRequest)
		return
	} else if !request.Valid() {
		//ctx.Logger.Error("enroll: error validating JSON")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Send the response
	w.WriteHeader(http.StatusCreated)
}
