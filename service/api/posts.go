package api

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) CreatePost(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// Parse request body
	r.ParseMultipartForm(100000000) // allows for an Imagesize of ~100MB should be good for 8k pictures
	// Parse the metadata
	metadataString := r.FormValue("post")
	var metadata Post
	err := json.Unmarshal([]byte(metadataString), &metadata)
	if err != nil {
		//ctx.Logger.WithError(err).Error("enroll: error decoding JSON") TODO figure out how to use those
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	// Parse the file
	file, _, err := r.FormFile("image")
	if err != nil {
		//ctx.Logger.WithError(err).Error("enroll: error decoding JSON") TODO figure out how to use those
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	defer file.Close()
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		log.Println(err)
	}

	// Save picture in folder
	var fid = "dlfi4986gknd"
	filepath := "pictures/" + fid + ".png"
	err = ioutil.WriteFile(filepath, fileBytes, 0777)
	if err != nil {
		//ctx.Logger.WithError(err).Error("enroll: error decoding JSON") TODO figure out how to use those
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Insert post in db
	// Increment posts in user

	// Send the response
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("content-type", "application/json")
	_ = json.NewEncoder(w).Encode("fid")
}

func (rt *_router) GetPost(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	pid := ps.ByName("pid")
	if len(pid) != 12 {
		//ctx.Logger.WithError(err).Error("enroll: error decoding JSON") TODO figure out how to use those
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Get post from db

	// Send the response
	var user = User{UserId: "uid", Name: "name", Posts: 5, Followers: 194, Followed: 207}
	var response = Post{Poster: user, DateTime: "datetime", Caption: "caption", PictureId: "pid", Likes: 497, Comments: 53}
	w.Header().Set("content-type", "application/json")
	_ = json.NewEncoder(w).Encode(response)
}

func (rt *_router) DeletePost(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// Parse parameters
	pid := ps.ByName("pid")
	if len(pid) != 12 {
		//ctx.Logger.WithError(err).Error("enroll: error decoding JSON") TODO figure out how to use those
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Delete post from db
	// Decrement posts from user

	// Send the response
	w.WriteHeader(http.StatusNoContent)
}
