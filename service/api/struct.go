package api

// General Schemas
// User related structs
type User struct {
	UserId    string `json:"userId"`
	Name      string `json:"name"`
	Followers int    `json:"followers"`
	Followed  int    `json:"followed"`
}

func (s *User) Valid() bool {
	return len(s.UserId) == 12 &&
		len(s.Name) >= 3 && len(s.Name) <= 16 &&
		s.Followers >= 0 &&
		s.Followed >= 0
}

type UserList struct {
	Length int    `json:"length"`
	Users  []User `json:"users"`
}

func (s *UserList) Valid() bool {
	return s.Length >= 0 &&
		true // TODO proper validation of the users
}

// Post related structs
type Post struct {
	Poster    User   `json:"poster"`
	DateTime  string `json:"date-time"`
	Caption   string `json:"caption"`
	PictureId string `json:"pictureId"`
	Likes     int    `json:"likes"`
	Comments  int    `json:"comments"`
}

func (s *Post) Valid() bool {
	return s.Poster.Valid() &&
		len(s.DateTime) == 24 && // TODO Better validation for date time
		len(s.Caption) >= 1 && len(s.Caption) <= 140 &&
		len(s.PictureId) == 12 &&
		s.Likes >= 0 &&
		s.Comments >= 0
}

// Comment related structs
type Comment struct {
	Poster   User   `json:"poster"`
	PostId   string `json:"postId"`
	DateTime string `json:"date-time"`
	Comment  string `json:"comment"`
}

func (s *Comment) Valid() bool {
	return s.Poster.Valid() &&
		len(s.PostId) == 12 &&
		len(s.DateTime) == 24 && // TODO Better validation for date time
		len(s.Comment) >= 1 && len(s.Comment) <= 140
}

type CommentList struct {
	Length   int       `json:"length"`
	Comments []Comment `json:"comments"`
}

func (s *CommentList) Valid() bool {
	return s.Length >= 0 &&
		true // TODO proper validation of the comments
}

// Login
type LoginRequest struct {
	Name string `json:"name"`
}

func (s *LoginRequest) Valid() bool {
	return len(s.Name) >= 3 && len(s.Name) <= 16
}

type LoginResult struct {
	Identifier string `json:"identifier"`
}

// General

type ChangeNameRequest struct {
	Name string `json:"name"`
}

func (s *ChangeNameRequest) Valid() bool {
	return len(s.Name) >= 3 && len(s.Name) <= 16
}

// Get number result
type GetCountResult struct {
	Count int `json:"count"`
}
