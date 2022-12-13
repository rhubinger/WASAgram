package database

import (
	"database/sql"

	"github.com/rhubinger/WASAgram/service/schemes"
)

func (db *appdbimpl) InsertUser(user schemes.User, identifier string) error {
	_, err := db.c.Exec("INSERT INTO users VALUES (?, ?, ?, ?, ?, ?);",
		user.UserId, user.Name, identifier, user.Posts, user.Followers, user.Followed)
	return err
}

func (db *appdbimpl) GetIdentifier(uid string) (string, error) {
	var identifier string
	err := db.c.QueryRow("SELECT identifier FROM users WHERE userId = ?;", uid).Scan(&identifier)
	return identifier, err
}

func (db *appdbimpl) UpdateUsername(name string, uid string) error {
	_, err := db.c.Exec("UPDATE users SET name = ? WHERE userId = ?;", name, uid)
	return err
}

func (db *appdbimpl) GetUser(uid string) (schemes.User, error) {
	var u schemes.User
	err := db.c.QueryRow("SELECT userid, name, posts, followers, followed 
						FROM users 
						WHERE userId = ?;", uid).Scan(&u.UserId, &u.Name, &identifier, &u.Posts, &u.Followers, &u.Followed)
	return u, err
}

func (db *appdbimpl) SearchUser(searchString string, searchType string) ([]schemes.User, error) {
	// TODO maybe order results by relevance somehow
	var err error
	var rows *sql.Rows
	searchString = "%" + searchString + "%"
	switch searchType {
	case "uid":
		rows, err = db.c.Query("SELECT userid, name, posts, followers, followed 
								FROM users 
								WHERE userId LIKE ?;", searchString)
	case "name":
		rows, err = db.c.Query("SELECT userid, name, posts, followers, followed 
								FROM users 
								WHERE name LIKE ?;", searchString)
	default:
		rows, err = db.c.Query("SELECT userid, name, posts, followers, followed
								FROM users 
								WHERE userId LIKE ? OR name LIKE ?;", searchString, searchString)
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	result := []schemes.User{}
	for rows.Next() {
		u := schemes.User{}
		err = rows.Scan(&u.UserId, &u.Name, &u.Posts, &u.Followers, &u.Followed)
		if err != nil {
			return nil, err
		}
		result = append(result, u)
	}
	return result, nil
}
