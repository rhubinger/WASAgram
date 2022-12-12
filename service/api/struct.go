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

// Validity Checker for Parameters
func ValidUid(uid string) bool {
	correctPattern, err := regexp.MatchString("@[a-zA-z0-9_.]{3,16}", uid)
	if err != nil {
		os.Stderr.WriteString(err.Error())
		return false
	}
	return len(uid) >= 3 && len(uid) <= 16 && correctPattern
}

func ValidId(id string) bool {
	correctPattern, err := regexp.MatchString("[a-zA-z0-9-_]{11}", id)
	if err != nil {
		os.Stderr.WriteString(err.Error())
		return false
	}
	return len(id) == 11 && correctPattern
}

func ValidSearchString(searchString string) bool {
	correctPattern, err := regexp.MatchString("[@a-zA-z0-9-_.]{1,30}", searchString)
	if err != nil {
		os.Stderr.WriteString(err.Error())
		return false
	}
	return len(searchString) >= 1 && len(searchString) <= 30 && correctPattern
}
