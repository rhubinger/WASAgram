package database

import "github.com/rhubinger/WASAgram/service/schemes"

func (db *appdbimpl) Like(pid string, uid string) error {
	_, err := db.c.Exec("INSERT INTO likes VALUES (?, ?);", pid, uid)
	return err
}

func (db *appdbimpl) Unlike(pid string, uid string) error {
	_, err := db.c.Exec("DELETE FROM likes WHERE postId = ? AND userId = ?", pid, uid)
	return err
}

func (db *appdbimpl) DeleteLikes(pid string) error {
	_, err := db.c.Exec("DELETE FROM likes WHERE postId = ?", pid)
	return err
}

func (db *appdbimpl) GetLikes(pid string) ([]schemes.User, error) {
	rows, err := db.c.Query(`SELECT u.userId, u.name, u.posts, u.followers, u.followed 
							 FROM users u, likes l
							 WHERE u.userId = l.userId AND l.postId = ?
							 ORDER BY u.name DESC;`, pid)
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

func (db *appdbimpl) GetLikeCount(pid string) (int, error) {
	var count int
	err := db.c.QueryRow("SELECT likes FROM posts WHERE postId = ?", pid).Scan(&count)
	return count, err
}

func (db *appdbimpl) IncrementLikeCount(pid string) error {
	_, err := db.c.Exec("UPDATE posts SET likes = likes + 1 WHERE postId = ?;", pid)
	return err
}

func (db *appdbimpl) DecrementLikeCount(pid string) error {
	_, err := db.c.Exec("UPDATE posts SET likes = likes - 1 WHERE postId = ?;", pid)
	return err
}
