package database

import "github.com/rhubinger/WASAgram/service/schemes"

func (db *appdbimpl) InsertComment(comment schemes.Comment) (string, error) {
	cid := db.GenerateId("commentId")
	_, err := db.c.Exec("INSERT INTO comments VALUES (?, ?, ?, CURRENT_TIMESTAMP, ?);",
		cid, comment.Poster.UserId, comment.PostId, comment.Comment)
	return cid, err
}

func (db *appdbimpl) DeleteComment(cid string) error {
	_, err := db.c.Exec("DELETE FROM comments where commentId = ?", cid)
	return err
}

func (db *appdbimpl) DeleteComments(pid string) error {
	_, err := db.c.Exec("DELETE FROM comments where postId = ?", pid)
	return err
}

func (db *appdbimpl) GetComments(pid string) ([]schemes.Comment, error) {
	rows, err := db.c.Query(`SELECT u.userId, u.name, u.posts, u.followers, u.followed, 
								 c.postId, c.uploadTime, c.commentText
							 FROM comments c, users u, posts p
							 WHERE c.userId = u.userId AND c.postId = p.postId AND c.postId = ?;`, pid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	comments := []schemes.Comment{}
	for rows.Next() {
		c := schemes.Comment{}
		err = rows.Scan(&c.Poster.UserId, &c.Poster.Name, &c.Poster.Posts, &c.Poster.Followers, &c.Poster.Followed,
			&c.PostId, &c.DateTime, &c.Comment)
		if err != nil {
			return nil, err
		}
		comments = append(comments, c)
	}

	return comments, err
}

func (db *appdbimpl) GetCommentCount(pid string) (int, error) {
	var count int
	err := db.c.QueryRow("SELECT comments FROM posts WHERE postId = ?", pid).Scan(&count)
	return count, err
}

func (db *appdbimpl) IncrementCommentCount(pid string) error {
	_, err := db.c.Exec("UPDATE posts SET comments = comments + 1 WHERE postId = ?;", pid)
	return err
}

func (db *appdbimpl) DecrementCommentCount(pid string) error {
	_, err := db.c.Exec("UPDATE posts SET comments = comments - 1 WHERE postId = ?;", pid)
	return err
}
