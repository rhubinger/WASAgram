package api

import (
	"os"
	"regexp"
)

// Get number result
type GetCountResult struct {
	Count int `json:"count"`
}

// Login
type LoginRequest struct {
	UserId string `json:"userId"`
}

func (s *LoginRequest) Valid() bool {
	correctPattern, err := regexp.MatchString("@[a-zA-z0-9_.]{3,16}", s.UserId)
	if err != nil {
		os.Stderr.WriteString(err.Error())
		return false
	}
	return len(s.UserId) >= 3 && len(s.UserId) <= 16 && correctPattern
}

type LoginResult struct {
	Identifier string `json:"identifier"`
}

// Change username
type ChangeNameRequest struct {
	Name string `json:"name"`
}

func (s *ChangeNameRequest) Valid() bool {
	correctPattern, err := regexp.MatchString("[a-zA-z0-9-. ]{1,30}", s.Name)
	if err != nil {
		os.Stderr.WriteString(err.Error())
		return false
	}
	return len(s.Name) >= 1 && len(s.Name) <= 30 && correctPattern
}

// Create post response
type CreatePostResponse struct {
	PostId string `json:"postId"`
}

// Create comment response
type CreateCommentResponse struct {
	CommentId string `json:"commentId"`
}
