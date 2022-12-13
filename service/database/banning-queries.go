package database

import "github.com/rhubinger/WASAgram/service/schemes"

func (db *appdbimpl) Ban(uid string, bid string) error {
	_, err := db.c.Exec("INSERT INTO bans VALUES (?, ?);", uid, bid)
	return err
}

func (db *appdbimpl) Unban(uid string, bid string) error {
	_, err := db.c.Exec("IDELETE FROM bans WHERE userId = ? AND bannedId = ?", uid, bid)
	return err
}

func (db *appdbimpl) GetBanned(uid string) ([]schemes.User, error) {
	rows, err := db.c.Query(`SELECT u.userId, u.name, u.posts, u.followers, u.followed 
							 FROM users u, bans b 
							 WHERE u.userId = b.bannedId AND b.userId = ?;`, uid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := []schemes.User{}
	for rows.Next() {
		u := schemes.User{}
		err = rows.Scan(&u.UserId, &u.Name, &u.Posts, &u.Followers, &u.Followed)
		if err != nil {
			return nil, err
		}
		users = append(users, u)
	}

	return users, err
}

func (db *appdbimpl) GetBannedCount(uid string) (int, error) {
	var count int
	err := db.c.QueryRow("SELECT COUNT(*) FROM bans WHERE userId = ?", uid).Scan(&count)
	return count, err
}
