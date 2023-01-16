package database

import "github.com/rhubinger/WASAgram/service/schemes"

func (db *appdbimpl) Follow(uid string, fid string) error {
	_, err := db.c.Exec("INSERT INTO followers VALUES (?, ?);", uid, fid)
	return err
}

func (db *appdbimpl) Unfollow(uid string, fid string) error {
	_, err := db.c.Exec("DELETE FROM followers WHERE userId = ? AND followerId = ?", uid, fid)
	return err
}

func (db *appdbimpl) GetFollowers(uid string) ([]schemes.User, error) {
	rows, err := db.c.Query(`SELECT u.userId, u.name, u.posts, u.followers, u.followed 
							 FROM users u, followers f
							 WHERE u.userId = f.followerId AND f.userId = ?
							 ORDER BY u.name DESC;`, uid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := []schemes.User{}
	for rows.Next() {
		if err = rows.Err(); err != nil {
			return nil, err
		}
		u := schemes.User{}
		err = rows.Scan(&u.UserId, &u.Name, &u.Posts, &u.Followers, &u.Followed)
		if err != nil {
			return nil, err
		}
		users = append(users, u)
	}

	return users, err
}

func (db *appdbimpl) GetFollowerCount(uid string) (int, error) {
	var count int
	err := db.c.QueryRow("SELECT followers FROM users WHERE userId = ?", uid).Scan(&count)
	return count, err
}

func (db *appdbimpl) IncrementFollowerCount(uid string) error {
	_, err := db.c.Exec("UPDATE users SET followers = followers + 1 WHERE userId = ?;", uid)
	return err
}

func (db *appdbimpl) DecrementFollowerCount(uid string) error {
	_, err := db.c.Exec("UPDATE users SET followers = followers - 1 WHERE userId = ?;", uid)
	return err
}

func (db *appdbimpl) GetFollowed(uid string) ([]schemes.User, error) {
	rows, err := db.c.Query(`SELECT u.userId, u.name, u.posts, u.followers, u.followed 
							 FROM users u, followers f
							 WHERE u.userId = f.userId AND f.followerId = ?
							 ORDER BY u.name DESC;`, uid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := []schemes.User{}
	for rows.Next() {
		if err = rows.Err(); err != nil {
			return nil, err
		}
		u := schemes.User{}
		err = rows.Scan(&u.UserId, &u.Name, &u.Posts, &u.Followers, &u.Followed)
		if err != nil {
			return nil, err
		}
		users = append(users, u)
	}

	return users, err
}

func (db *appdbimpl) GetFollowedCount(uid string) (int, error) {
	var count int
	err := db.c.QueryRow("SELECT followed FROM users WHERE userId = ?", uid).Scan(&count)
	return count, err
}

func (db *appdbimpl) IncrementFollowedCount(uid string) error {
	_, err := db.c.Exec("UPDATE users SET followed = followed + 1 WHERE userId = ?;", uid)
	return err
}

func (db *appdbimpl) DecrementFollowedCount(uid string) error {
	_, err := db.c.Exec("UPDATE users SET followed = followed - 1 WHERE userId = ?;", uid)
	return err
}
