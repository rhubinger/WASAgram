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
	correctUidPattern, err := regexp.MatchString("@[a-zA-z0-9_.]{3,16}", s.UserId)
	if err != nil {
		os.Stderr.WriteString(err.Error())
		return false
	}
	correctNamePattern, err := regexp.MatchString("[a-zA-z0-9-. ]{1,30}", s.Name)
	if err != nil {
		os.Stderr.WriteString(err.Error())
		return false
	}
	return len(s.UserId) >= 3 && len(s.UserId) <= 16 && correctUidPattern &&
		len(s.Name) >= 1 && len(s.Name) <= 30 && correctNamePattern &&
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
	correctPostIdPattern, err := regexp.MatchString("[a-zA-z0-9-_]{11}", s.PictureId)
	if err != nil {
		os.Stderr.WriteString(err.Error())
		return false
	}
	correctUidPattern, err := regexp.MatchString("@[a-zA-z0-9_.]{3,16}", s.UserId)
	if err != nil {
		os.Stderr.WriteString(err.Error())
		return false
	}
	correctUsernamePattern, err := regexp.MatchString("[a-zA-z0-9-. ]{1,30}", s.UserId)
	if err != nil {
		os.Stderr.WriteString(err.Error())
		return false
	}
	correctDateTimePattern, err := regexp.MatchString("[0-9]{4}-[0-9]{2}-[0-9]{2} [0-9]{2}:[0-9]{2}:[0-9]{2}", s.DateTime)
	if err != nil {
		os.Stderr.WriteString(err.Error())
		return false
	}
	correctPictureIdPattern, err := regexp.MatchString("[a-zA-z0-9-_]{11}", s.PictureId)
	if err != nil {
		os.Stderr.WriteString(err.Error())
		return false
	}
	return len(s.PostId) == 11 && correctPostIdPattern &&
		len(s.UserId) >= 3 && len(s.UserId) <= 16 && correctUidPattern &&
		len(s.Username) >= 1 && len(s.Username) <= 30 && correctUsernamePattern &&
		len(s.DateTime) == 19 && correctDateTimePattern &&
		len(s.Caption) >= 1 && len(s.Caption) <= 140 &&
		len(s.PictureId) == 11 && correctPictureIdPattern &&
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
	correctCommentIdPattern, err := regexp.MatchString("[a-zA-z0-9-_]{11}", s.CommentId)
	if err != nil {
		os.Stderr.WriteString(err.Error())
		return false
	}
	correctUidPattern, err := regexp.MatchString("@[a-zA-z0-9_.]{3,16}", s.UserId)
	if err != nil {
		os.Stderr.WriteString(err.Error())
		return false
	}
	correctUsernamePattern, err := regexp.MatchString("[a-zA-z0-9-. ]{1,30}", s.UserId)
	if err != nil {
		os.Stderr.WriteString(err.Error())
		return false
	}
	correctPostIdPattern, err := regexp.MatchString("[a-zA-z0-9-_]{11}", s.PostId)
	if err != nil {
		os.Stderr.WriteString(err.Error())
		return false
	}
	correctDateTimePattern, err := regexp.MatchString("[0-9]{4}-[0-9]{2}-[0-9]{2} [0-9]{2}:[0-9]{2}:[0-9]{2}", s.DateTime)
	if err != nil {
		os.Stderr.WriteString(err.Error())
		return false
	}
	return len(s.CommentId) == 11 && correctCommentIdPattern &&
		len(s.PostId) == 11 && correctPostIdPattern &&
		len(s.UserId) >= 3 && len(s.UserId) <= 16 && correctUidPattern &&
		len(s.Username) >= 1 && len(s.Username) <= 30 && correctUsernamePattern &&
		len(s.DateTime) == 19 && correctDateTimePattern &&
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
