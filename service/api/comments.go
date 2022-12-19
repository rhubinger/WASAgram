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
	} else if postExists, err := rt.db.PostExists(pid); err != nil {
		rt.baseLogger.Error("CreateComment: Error while checking for post in db")
		w.WriteHeader(http.StatusInternalServerError)
		return
	} else if !postExists {
		rt.baseLogger.Error("CreateComment: Post doesn't exist in db")
		w.WriteHeader(http.StatusNotFound)
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
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// Authentification as user with userId comment.userId
	identifier, err := ParseIdentifier(r)
	if err != nil {
		rt.baseLogger.WithError(err).Error("CreateComment: Failed to parse identifier from reques")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	authorized, err := rt.db.AuthorizeAsUser(identifier, comment.UserId)
	if err != nil {
		rt.baseLogger.WithError(err).Error("CreateComment: Error occured during authorization")
		w.WriteHeader(http.StatusInternalServerError)
		return
	} else if !authorized {
		rt.baseLogger.Error("CreateComment: User unauthorized to access resource")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// Insert into db
	comment.PostId = pid
	if userExists, err := rt.db.UserExists(comment.UserId); err != nil {
		rt.baseLogger.Error("CreateComment: Error while checking for user in db")
		w.WriteHeader(http.StatusInternalServerError)
		return
	} else if !userExists {
		rt.baseLogger.Error("CreateComment: User doesn't exist in db")
		w.WriteHeader(http.StatusNotFound)
		return
	}
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
	} else if postExists, err := rt.db.PostExists(pid); err != nil {
		rt.baseLogger.Error("GetComments: Error while checking for post in db")
		w.WriteHeader(http.StatusInternalServerError)
		return
	} else if !postExists {
		rt.baseLogger.Error("GetComments: Post doesn't exist in db")
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// Authentification as not banned by the user with userId post.UserId
	identifier, err := ParseIdentifier(r)
	if err != nil {
		rt.baseLogger.WithError(err).Error("GetComments: Failed to parse identifier from reques")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	post, err := rt.db.GetPost(pid)
	if err != nil {
		rt.baseLogger.WithError(err).Error("GetComments: Failed to get post from db")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	authorized, err := rt.db.AuthorizeAsNotBanned(identifier, post.UserId)
	if err != nil {
		rt.baseLogger.WithError(err).Error("GetComments: Error occured during authorization")
		w.WriteHeader(http.StatusInternalServerError)
		return
	} else if !authorized {
		rt.baseLogger.Error("GetComments: User unauthorized to access resource")
		w.WriteHeader(http.StatusUnauthorized)
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
	w.WriteHeader(http.StatusOK)
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
	} else if postExists, err := rt.db.PostExists(pid); err != nil {
		rt.baseLogger.Error("GetCommentCount: Error while checking for post in db")
		w.WriteHeader(http.StatusInternalServerError)
		return
	} else if !postExists {
		rt.baseLogger.Error("GetCommentCount: Post doesn't exist in db")
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// Authentification as not banned by the user with userId post.UserId
	identifier, err := ParseIdentifier(r)
	if err != nil {
		rt.baseLogger.WithError(err).Error("GetCommentCount: Failed to parse identifier from reques")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	post, err := rt.db.GetPost(pid)
	if err != nil {
		rt.baseLogger.WithError(err).Error("GetCommentCount: Failed to get post from db")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	authorized, err := rt.db.AuthorizeAsNotBanned(identifier, post.UserId)
	if err != nil {
		rt.baseLogger.WithError(err).Error("GetCommentCount: Error occured during authorization")
		w.WriteHeader(http.StatusInternalServerError)
		return
	} else if !authorized {
		rt.baseLogger.Error("GetCommentCount: User unauthorized to access resource")
		w.WriteHeader(http.StatusUnauthorized)
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
	w.WriteHeader(http.StatusOK)
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
	} else if postExists, err := rt.db.PostExists(pid); err != nil {
		rt.baseLogger.Error("DeleteComment: Error while checking for post in db")
		w.WriteHeader(http.StatusInternalServerError)
		return
	} else if !postExists {
		rt.baseLogger.Error("DeleteComment: Post doesn't exist in db")
		w.WriteHeader(http.StatusNotFound)
		return
	}
	cid := ps.ByName("cid")
	if !schemes.ValidId(cid) {
		rt.baseLogger.Error("CommentId (cid) invalid")
		w.WriteHeader(http.StatusBadRequest)
		return
	} else if commentExists, err := rt.db.CommentExists(cid); err != nil {
		rt.baseLogger.Error("DeleteComment: Error while checking for comment in db")
		w.WriteHeader(http.StatusInternalServerError)
		return
	} else if !commentExists {
		rt.baseLogger.Error("DeleteComment: Comment doesn't exist in db")
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// Authentification as user with userId comment.userId
	identifier, err := ParseIdentifier(r)
	if err != nil {
		rt.baseLogger.WithError(err).Error("DeleteComment: Failed to parse identifier from reques")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	comment, err := rt.db.GetComment(cid)
	if err != nil {
		rt.baseLogger.WithError(err).Error("DeleteComment: Failed to get comment from db")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	authorized, err := rt.db.AuthorizeAsUser(identifier, comment.UserId)
	if err != nil {
		rt.baseLogger.WithError(err).Error("DeleteComment: Error occured during authorization")
		w.WriteHeader(http.StatusInternalServerError)
		return
	} else if !authorized {
		rt.baseLogger.Error("DeleteComment: User unauthorized to access resource")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// Delete comment from db
	err = rt.db.DeleteComment(cid)
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
