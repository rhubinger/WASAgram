package schemes

import (
	"os"
	"regexp"
)

// General Schemas
// User related structs
type User struct {
	UserId    string `json:"userId"`
	Name      string `json:"name"`
	Posts     int    `json:"posts"`
	Followers int    `json:"followers"`
	Followed  int    `json:"followed"`
}

func (s *User) Valid() bool {
	return ValidUserId(s.UserId) &&
		ValidUsername(s.Name) &&
		s.Posts >= 0 &&
		s.Followers >= 0 &&
		s.Followed >= 0
}

type UserList struct {
	Length int    `json:"length"`
	Users  []User `json:"users"`
}

func (s *UserList) Valid() bool {
	if s.Length < 0 {
		return false
	}
	for _, u := range s.Users {
		if !u.Valid() {
			return false
		}
	}
	return true
}

// Post related structs
type Post struct {
	PostId    string `json:"postId"`
	UserId    string `json:"userId"`
	Username  string `json:"username"`
	DateTime  string `json:"date-time"`
	Caption   string `json:"caption"`
	PictureId string `json:"pictureId"`
	Likes     int    `json:"likes"`
	Comments  int    `json:"comments"`
}

func (s *Post) Valid() bool {
	return ValidId(s.PostId) &&
		ValidUserId(s.UserId) &&
		ValidUsername(s.Username) &&
		ValidDatetime(s.DateTime) &&
		len(s.Caption) >= 1 && len(s.Caption) <= 140 &&
		ValidId(s.PictureId) &&
		s.Likes >= 0 &&
		s.Comments >= 0
}

// Comment related structs
type Comment struct {
	CommentId string `json:"commentId"`
	PostId    string `json:"postId"`
	UserId    string `json:"userId"`
	Username  string `json:"username"`
	DateTime  string `json:"date-time"`
	Comment   string `json:"comment"`
}

func (s *Comment) Valid() bool {
	return ValidId(s.CommentId) &&
		ValidId(s.PostId) &&
		ValidUserId(s.UserId) &&
		ValidUsername(s.Username) &&
		ValidDatetime(s.DateTime) &&
		len(s.Comment) >= 1 && len(s.Comment) <= 140
}

type CommentList struct {
	Length   int       `json:"length"`
	Comments []Comment `json:"comments"`
}

func (s *CommentList) Valid() bool {
	if s.Length < 0 {
		return false
	}
	for _, c := range s.Comments {
		if !c.Valid() {
			return false
		}
	}
	return true
}

// Validity Checker for Parameters
func ValidId(id string) bool {
	correctPattern, err := regexp.MatchString("[a-zA-z0-9-_]{11}", id)
	if err != nil {
		os.Stderr.WriteString(err.Error())
		return false
	}
	return len(id) == 11 && correctPattern
}

func ValidUserId(uid string) bool {
	correctPattern, err := regexp.MatchString("@[a-zA-z0-9_.]{3,16}", uid)
	if err != nil {
		os.Stderr.WriteString(err.Error())
		return false
	}
	return len(uid) >= 3 && len(uid) <= 16 && correctPattern
}

func ValidUsername(uid string) bool {
	correctPattern, err := regexp.MatchString("[a-zA-z0-9-. ]{1,30}", uid)
	if err != nil {
		os.Stderr.WriteString(err.Error())
		return false
	}
	return len(uid) >= 1 && len(uid) <= 30 && correctPattern
}

func ValidDatetime(uid string) bool {
	correctPattern, err := regexp.MatchString("[0-9]{4}-[0-9]{2}-[0-9]{2} [0-9]{2}:[0-9]{2}:[0-9]{2}", uid)
	if err != nil {
		os.Stderr.WriteString(err.Error())
		return false
	}
	return len(uid) == 19 && correctPattern
}

func ValidSearchString(searchString string) bool {
	correctPattern, err := regexp.MatchString("[@a-zA-z0-9-_.]{1,30}", searchString)
	if err != nil {
		os.Stderr.WriteString(err.Error())
		return false
	}
	return len(searchString) >= 1 && len(searchString) <= 30 && correctPattern
}
