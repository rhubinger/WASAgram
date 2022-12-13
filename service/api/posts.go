package api

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/rhubinger/WASAgram/service/schemes"
)

func (rt *_router) CreatePost(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// Parse request body
	r.ParseMultipartForm(100000000) // allows for an Imagesize of ~100MB should be good for 8k pictures
	// Parse the metadata
	metadataString := r.FormValue("post")
	var metadata schemes.Post
	err := json.Unmarshal([]byte(metadataString), &metadata)
	if err != nil {
		rt.baseLogger.WithError(err).Error("CreatePost: Request body (JSON) couldn't be parsed")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	// Parse the file
	file, _, err := r.FormFile("image")
	if err != nil {
		rt.baseLogger.Error("CreatePost: Request body (file) couldn't be parsed")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	defer file.Close()
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		log.Println(err)
	}

	// Insert post in db
	// Insert image in db
	pid, err := rt.db.InsertPicture(fileBytes) // TODO change to proper id
	if err != nil {
		rt.baseLogger.WithError(err).Error("CreatePost: Failed to insert picture in db")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	// Increment posts in user

	// Send the response
	var post = CreatePostResponse{PostId: pid}
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("content-type", "application/json")
	_ = json.NewEncoder(w).Encode(post)
}

func (rt *_router) GetPost(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	pid := ps.ByName("pid")
	if !ValidId(pid) {
		rt.baseLogger.Error("PostId (pid) invalid")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Get post from db

	// Send the response
	var user = schemes.User{UserId: "uid", Name: "name", Posts: 5, Followers: 194, Followed: 207}
	var response = schemes.Post{Poster: user, DateTime: "datetime", Caption: "caption", PictureId: "pid", Likes: 497, Comments: 53}
	w.Header().Set("content-type", "application/json")
	_ = json.NewEncoder(w).Encode(response)
}

func (rt *_router) DeletePost(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// Parse parameters
	pid := ps.ByName("pid")
	if !ValidId(pid) {
		rt.baseLogger.Error("PostId (pid) invalid")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Delete post from db
	// Delete picture from db
	// Delete likes from db
	// Delete comments from db
	// Decrement posts from user
	rt.db.DeletePicture(pid)

	// Send the response
	w.WriteHeader(http.StatusNoContent)
}
