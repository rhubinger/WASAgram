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
	if !schemes.ValidId(pid) {
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
	comment.PostId = pid
	cid, err := rt.db.InsertComment(comment)
	if err != nil {
		rt.baseLogger.WithError(err).Error("CreateComment: failed to insert comment into db")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	// Increment posts comment count
	err = rt.db.IncrementCommentCount(comment.PostId)
	if err != nil {
		rt.baseLogger.WithError(err).Error("CreateComment: failed to update posts comment count in db")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Send the response
	var response = CreateCommentResponse{CommentId: cid}
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("content-type", "application/json")
	_ = json.NewEncoder(w).Encode(response)
}

func (rt *_router) GetComments(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// Get post from db
	pid := ps.ByName("pid")
	if !schemes.ValidId(pid) {
		rt.baseLogger.Error("PostId (pid) invalid")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Get comments from db
	comments, err := rt.db.GetComments(pid)
	if err != nil {
		rt.baseLogger.WithError(err).Error("GetComments: failed to get comments from db")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	// Shorten if too many comments
	if len(comments) > 1000 {
		comments = comments[0:1000]
	}

	// Send the response
	var response = schemes.CommentList{Length: len(comments), Comments: comments}
	w.Header().Set("content-type", "application/json")
	_ = json.NewEncoder(w).Encode(response)
}

func (rt *_router) GetCommentCount(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// Get post from db
	pid := ps.ByName("pid")
	if !schemes.ValidId(pid) {
		rt.baseLogger.Error("PostId (pid) invalid")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Get count of comments from db
	count, err := rt.db.GetCommentCount(pid)
	if err != nil {
		rt.baseLogger.WithError(err).Error("GetCommentCount: failed to get comment count from db")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Send the response
	var response = GetCountResult{Count: count}
	w.Header().Set("content-type", "application/json")
	_ = json.NewEncoder(w).Encode(response)
}

func (rt *_router) DeleteComment(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// Get post from db
	pid := ps.ByName("pid")
	if !schemes.ValidId(pid) {
		rt.baseLogger.Error("PostId (pid) invalid")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	cid := ps.ByName("cid")
	if !schemes.ValidId(cid) {
		rt.baseLogger.Error("CommentId (cid) invalid")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Delete comment from db
	err := rt.db.DeleteComment(cid)
	if err != nil {
		rt.baseLogger.WithError(err).Error("DeleteComment: failed to delete comment from db")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	// Decrement comment count on post
	err = rt.db.DecrementCommentCount(pid)
	if err != nil {
		rt.baseLogger.WithError(err).Error("DeleteComment: failed to update posts comment count in db")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Send the response
	w.WriteHeader(http.StatusNoContent)
}
