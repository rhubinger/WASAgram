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
	err := r.ParseMultipartForm(100000000) // Allows for an image size of ~100MB; should be good for 8k pictures
	if err != nil {
		rt.baseLogger.WithError(err).Error("CreatePost: Error occured while parsing multipart form")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	// Parse the metadata
	metadataString := r.FormValue("post")
	var metadata schemes.Post
	err = json.Unmarshal([]byte(metadataString), &metadata)
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

	// Insert image in db
	pictureId, err := rt.db.InsertPicture(fileBytes)
	if err != nil {
		rt.baseLogger.WithError(err).Error("CreatePost: Failed to insert picture in db")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	// Format metadata and insert post in db
	metadata.PictureId = pictureId
	metadata.Likes = 0
	metadata.Comments = 0
	if userExists, err := rt.db.UserExists(metadata.UserId); err != nil {
		rt.baseLogger.Error("CreatePost: Error while checking for user in db")
		w.WriteHeader(http.StatusInternalServerError)
		return
	} else if userExists {
		rt.baseLogger.Error("CreatePost: User doesn't exist in db")
		w.WriteHeader(http.StatusNotFound)
		return
	}
	pid, err := rt.db.InsertPost(metadata)
	if err != nil {
		rt.baseLogger.WithError(err).Error("Create Post: failed insert post (metadata) into db")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	// Increment posts in user
	err = rt.db.IncrementPostCount(metadata.UserId)
	if err != nil {
		rt.baseLogger.WithError(err).Error("CreatePost: Failed to increment post count of posting user in db")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Send the response
	var response = CreatePostResponse{PostId: pid}
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("content-type", "application/json")
	_ = json.NewEncoder(w).Encode(response)
}

func (rt *_router) GetPost(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	pid := ps.ByName("pid")
	if !schemes.ValidId(pid) {
		rt.baseLogger.Error("PostId (pid) invalid")
		w.WriteHeader(http.StatusBadRequest)
		return
	} else if postExists, err := rt.db.PostExists(pid); err != nil {
		rt.baseLogger.Error("GetPost: Error while checking for post in db")
		w.WriteHeader(http.StatusInternalServerError)
		return
	} else if postExists {
		rt.baseLogger.Error("GetPost: Post doesn't exist in db")
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// Get post from db
	post, err := rt.db.GetPost(pid)
	if err != nil {
		rt.baseLogger.WithError(err).Error("GetPost: Failed to get post from db")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Send the response
	w.WriteHeader(http.StatusOK)
	w.Header().Set("content-type", "application/json")
	_ = json.NewEncoder(w).Encode(post)
}

func (rt *_router) DeletePost(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// Parse parameters
	pid := ps.ByName("pid")
	if !schemes.ValidId(pid) {
		rt.baseLogger.Error("PostId (pid) invalid")
		w.WriteHeader(http.StatusBadRequest)
		return
	} else if postExists, err := rt.db.PostExists(pid); err != nil {
		rt.baseLogger.Error("DeletePost: Error while checking for post in db")
		w.WriteHeader(http.StatusInternalServerError)
		return
	} else if postExists {
		rt.baseLogger.Error("DeletePost: Post doesn't exist in db")
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// Get from database post to get pictureId and userId of poster
	post, err := rt.db.GetPost(pid)
	if err != nil {
		rt.baseLogger.WithError(err).Error("DeletePost: failed to get Post from db")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	// Delete post from db
	err = rt.db.DeletePost(pid)
	if err != nil {
		rt.baseLogger.WithError(err).Error("DeletePost: failed to delete post from db")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	// Delete picture from db
	err = rt.db.DeletePicture(post.PictureId)
	if err != nil {
		rt.baseLogger.WithError(err).Error("DeletePost: failed to delete picture from db")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	// Delete likes from db
	err = rt.db.DeleteLikes(pid)
	if err != nil {
		rt.baseLogger.WithError(err).Error("DeletePost: failed to delete posts likes from db")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	// Delete comments from db
	err = rt.db.DeleteComments(pid)
	if err != nil {
		rt.baseLogger.WithError(err).Error("DeletePost: failed to delete posts comments from db")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	// Decrement post count from user
	err = rt.db.DecrementPostCount(post.UserId)
	if err != nil {
		rt.baseLogger.WithError(err).Error("DeletePost: failed to decrement the posters post count in db")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Send the response
	w.WriteHeader(http.StatusNoContent)
}
