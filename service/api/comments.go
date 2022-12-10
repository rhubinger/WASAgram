package api

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) CreateComment(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	pid := ps.ByName("pid")
	if len(pid) != 12 {
		//ctx.Logger.WithError(err).Error("enroll: error decoding JSON") TODO figure out how to use those
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var comment Comment
	err := json.NewDecoder(r.Body).Decode(&comment)
	_ = r.Body.Close()
	if err != nil {
		//ctx.Logger.WithError(err).Error("enroll: error decoding JSON") TODO figure out how to use those
		w.WriteHeader(http.StatusBadRequest)
		return
	} else if !comment.Valid() {
		//ctx.Logger.Error("enroll: error validating JSON")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Send the response
	w.WriteHeader(http.StatusCreated)
}

func (rt *_router) GetComments(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	pid := ps.ByName("pid")
	if len(pid) != 12 {
		//ctx.Logger.WithError(err).Error("enroll: error decoding JSON") TODO figure out how to use those
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Send the response
	var length = 3
	var comments = make([]Comment, 0, length)
	for i := 0; i < length; i++ {
		var u = User{UserId: "uid", Name: "Konrad Zuse", Followers: 1783, Followed: 1}
		var c = Comment{Poster: u, PostId: "pid", DateTime: "datetime", Comment: "commenting..."}
		comments = append(comments, c)
	}

	var response = CommentList{Length: length, Comments: comments}

	w.Header().Set("content-type", "application/json")
	_ = json.NewEncoder(w).Encode(response)
}

func (rt *_router) GetCommentCount(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	pid := ps.ByName("pid")
	if len(pid) != 12 {
		//ctx.Logger.WithError(err).Error("enroll: error decoding JSON") TODO figure out how to use those
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Send the response
	var response = GetCountResult{Count: 7}
	w.Header().Set("content-type", "application/json")
	_ = json.NewEncoder(w).Encode(response)
}

func (rt *_router) DeleteComment(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	pid := ps.ByName("pid")
	if len(pid) != 12 {
		//ctx.Logger.WithError(err).Error("enroll: error decoding JSON") TODO figure out how to use those
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	cid := ps.ByName("cid")
	if len(cid) != 12 {
		//ctx.Logger.WithError(err).Error("enroll: error decoding JSON") TODO figure out how to use those
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
