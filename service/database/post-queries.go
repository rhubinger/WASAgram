package database

import "github.com/rhubinger/WASAgram/service/schemes"

func (db *appdbimpl) InsertPost(post schemes.Post) (string, error) {
	pid := db.GenerateId("postId")
	_, err := db.c.Exec("INSERT INTO posts VALUES (?, ?, CURRENT_TIMESTAMP, ?, ?, ?, ?);",
		pid, post.UserId, post.Caption, post.PictureId, post.Likes, post.Comments)
	return pid, err
}

func (db *appdbimpl) DeletePost(pid string) error {
	_, err := db.c.Exec("DELETE FROM posts WHERE postId = ?;",
		pid)
	return err
}

func (db *appdbimpl) GetPictureId(pid string) (string, error) {
	var pictureId string
	err := db.c.QueryRow("SELECT pictureId FROM posts WHERE postId = ?", pid).Scan(&pictureId)
	return pictureId, err
}

func (db *appdbimpl) GetPost(pid string) (schemes.Post, error) {
	var p schemes.Post
	err := db.c.QueryRow(`SELECT p.postId, p.userId, u.name, p.uploadTime, p.caption, p.pictureId, p.likes, p.comments
						  FROM posts p, users u
						  WHERE p.userId = u.usersId AND p.postId = ?`,
		pid).Scan(&p.PostId, &p.UserId, &p.Username, &p.DateTime, &p.Caption, &p.PictureId, &p.Likes, &p.Comments)
	return p, err
}

func (db *appdbimpl) GetStream(uid string) ([]schemes.Post, error) {
	rows, err := db.c.Query(`SELECT p.postId, p.userId, u.name, p.uploadTime, p.caption, p.pictureId, p.likes, p.comments
							 FROM followers f, posts p, users u
							 WHERE f.userId = p.userId AND p.userId = u.userId AND f.followerId = ?
							 ORDER BY p.uploadTime DESC;`, uid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	posts := []schemes.Post{}
	for rows.Next() {
		p := schemes.Post{}
		err = rows.Scan(&p.PostId, &p.UserId, &p.Username, &p.DateTime, &p.Caption, &p.PictureId, &p.Likes, &p.Comments)
		if err != nil {
			return nil, err
		}
		posts = append(posts, p)
	}

	return posts, err
}

func (db *appdbimpl) GetPosts(uid string) ([]schemes.Post, error) {
	rows, err := db.c.Query(`SELECT p.postId, p.userId, u.name, p.uploadTime, p.caption, p.pictureId, p.likes, p.comments
							 FROM posts p, users u
							 WHERE p.userId = p.userId AND p.userId = ?
							 ORDER BY p.uploadTime DESC;`, uid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	posts := []schemes.Post{}
	for rows.Next() {
		p := schemes.Post{}
		err = rows.Scan(&p.PostId, &p.UserId, &p.Username, &p.DateTime, &p.Caption, &p.PictureId, &p.Likes, &p.Comments)
		if err != nil {
			return nil, err
		}
		posts = append(posts, p)
	}

	return posts, err
}

func (db *appdbimpl) GetPostCount(uid string) (int, error) {
	var count int
	err := db.c.QueryRow("SELECT posts FROM users WHERE userId = ?", uid).Scan(&count)
	return count, err
}

func (db *appdbimpl) IncrementPostCount(uid string) error {
	_, err := db.c.Exec("UPDATE users SET posts = posts + 1 WHERE userId = ?", uid)
	return err
}

func (db *appdbimpl) DecrementPostCount(uid string) error {
	_, err := db.c.Exec("UPDATE users SET posts = posts - 1 WHERE userId = ?", uid)
	return err
}
