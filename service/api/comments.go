package api

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/rhubinger/WASAgram/service/schemes"
)

func (rt *_router) CreateComment(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// Get post from db
	pid := ps.ByName("pid")
	if !ValidId(pid) {
		rt.baseLogger.Error("PostId (pid) invalid")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Parse request body
	var comment schemes.Comment
	err := json.NewDecoder(r.Body).Decode(&comment)
	_ = r.Body.Close()
	if err != nil {
		rt.baseLogger.WithError(err).Error("CreateComment: Failed to parse request body")
		w.WriteHeader(http.StatusBadRequest)
		return
	} else if !comment.Valid() {
		rt.baseLogger.Error("CreateComment: Request Body invalid")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Insert into db
	cid := "tesCID"
	// Increment posts comment count

	// Send the response
	var response = CreateCommentResponse{CommentId: cid}
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("content-type", "application/json")
	_ = json.NewEncoder(w).Encode(response)
}

func (rt *_router) GetComments(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// Get post from db
	pid := ps.ByName("pid")
	if !ValidId(pid) {
		rt.baseLogger.Error("PostId (pid) invalid")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Get comments from db

	// Send the response
	var length = 3
	var comments = make([]schemes.Comment, 0, length)
	for i := 0; i < length; i++ {
		var u = schemes.User{UserId: "uid", Name: "Konrad Zuse", Posts: 5, Followers: 1783, Followed: 1}
		var c = schemes.Comment{Poster: u, PostId: "pid", DateTime: "datetime", Comment: "commenting..."}
		comments = append(comments, c)
	}
	var response = schemes.CommentList{Length: length, Comments: comments}
	w.Header().Set("content-type", "application/json")
	_ = json.NewEncoder(w).Encode(response)
}

func (rt *_router) GetCommentCount(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// Get post from db
	pid := ps.ByName("pid")
	if !ValidId(pid) {
		rt.baseLogger.Error("PostId (pid) invalid")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Get count of comments from db

	// Send the response
	var response = GetCountResult{Count: 7}
	w.Header().Set("content-type", "application/json")
	_ = json.NewEncoder(w).Encode(response)
}

func (rt *_router) DeleteComment(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// Get post from db
	pid := ps.ByName("pid")
	if !ValidId(pid) {
		rt.baseLogger.Error("PostId (pid) invalid")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	cid := ps.ByName("cid")
	if !ValidId(cid) {
		rt.baseLogger.Error("PostId (pid) invalid")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Delete comment from db
	// Decrement comment count on post

	// Send the response
	w.WriteHeader(http.StatusNoContent)
}
