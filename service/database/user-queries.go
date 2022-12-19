package database

import (
	"github.com/rhubinger/WASAgram/service/schemes"
)

func (db *appdbimpl) InsertUser(user schemes.User) (string, error) {
	identifier := db.GenerateIdentifier()
	_, err := db.c.Exec("INSERT INTO users VALUES (?, ?, ?, ?, ?, ?);",
		user.UserId, user.Name, identifier, user.Posts, user.Followers, user.Followed)
	return identifier, err
}

func (db *appdbimpl) GetIdentifier(uid string) (string, error) {
	var identifier string
	err := db.c.QueryRow("SELECT identifier FROM users WHERE userId = ?;", uid).Scan(&identifier)
	return identifier, err
}

func (db *appdbimpl) GetUserId(identifier string) (string, error) {
	var userId string
	err := db.c.QueryRow("SELECT userId FROM users WHERE identifier = ?;", identifier).Scan(&userId)
	return identifier, err
}

func (db *appdbimpl) UpdateUsername(name string, uid string) error {
	_, err := db.c.Exec("UPDATE users SET name = ? WHERE userId = ?;", name, uid)
	return err
}

func (db *appdbimpl) GetUser(uid string) (schemes.User, error) {
	var u schemes.User
	err := db.c.QueryRow(`SELECT userId, name, posts, followers, followed
						FROM users 
						WHERE userId = ?;`, uid).Scan(&u.UserId, &u.Name, &u.Posts, &u.Followers, &u.Followed)
	return u, err
}

func (db *appdbimpl) SearchUser(searchString string, uid string) ([]schemes.User, error) {
	searchString = "%" + searchString + "%"
	rows, err := db.c.Query(`SELECT userId, name, posts, followers, followed
	FROM users 
	WHERE userId LIKE ? OR name LIKE ?`, searchString, searchString, uid, uid, uid, uid)
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

/* Better query but has bug which needs to be figured out (shows no results when it should be)
`WITH matches AS
	(
		SELECT userId, name, posts, followers, followed
		FROM users
		WHERE userId LIKE ? OR name LIKE ?
	),
	relevance AS
	(
		WITH contacts AS
		(
			SELECT f.followerId AS user, f.userId AS contact
			FROM followers f
			UNION
			SELECT f.userId AS user, f.followerId AS contact
			FROM followers f
		)
			SELECT ?, u.userId as searched, 0 as relevance
			FROM users u
			WHERE u.userId != ?
		UNION
			SELECT c.user as searcher, c.contact as searched, 2147483647 as relevance --MAXINT
			FROM contacts c
			WHERE c.user = ?
		UNION
			SELECT c1.user as searcher, c2.user as searched, COUNT(c2.user) as relevance
			FROM contacts c1, contacts c2
			WHERE c1.user = ? AND c1.contact = c2.contact AND NOT c1.user = c2.user
			GROUP BY c2.user
	)
	SELECT userId, name, posts, followers, followed, relevance
	FROM matches m, relevance r
	WHERE m.userId = r.searched
	GROUP BY m.userId
	ORDER BY MAX(r.relevance) DESC, followers DESC`
*/
