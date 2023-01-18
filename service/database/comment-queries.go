package database

import "github.com/rhubinger/WASAgram/service/schemes"

func (db *appdbimpl) InsertComment(comment schemes.Comment) (string, error) {
	cid := db.GenerateId("commentId")
	_, err := db.c.Exec("INSERT INTO comments VALUES (?, ?, ?, CURRENT_TIMESTAMP, ?);",
		cid, comment.UserId, comment.PostId, comment.Comment)
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

func (db *appdbimpl) GetComment(cid string) (schemes.Comment, error) {
	var c schemes.Comment
	err := db.c.QueryRow(`SELECT c.commentId, c.postId, c.userId, u.name, c.uploadTime, c.commentText
							 FROM comments c, users u
							 WHERE c.userId = u.userId AND c.commentId = ?`, cid).Scan(
		&c.CommentId, &c.PostId, &c.UserId, &c.Username, &c.DateTime, &c.Comment)
	return c, err
}

func (db *appdbimpl) GetComments(pid string) ([]schemes.Comment, error) {
	rows, err := db.c.Query(`SELECT c.commentId, c.postId, c.userId, u.name, c.uploadTime, c.commentText
							 FROM comments c, users u
							 WHERE c.userId = u.userId AND c.postId = ?
							 ORDER BY c.uploadTime DESC;`, pid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	comments := []schemes.Comment{}
	for rows.Next() {
		if err = rows.Err(); err != nil {
			return nil, err
		}
		c := schemes.Comment{}
		err = rows.Scan(&c.CommentId, &c.PostId, &c.UserId, &c.Username, &c.DateTime, &c.Comment)
		if err != nil {
			return nil, err
		}
		comments = append(comments, c)
	}

	return comments, err
}

func (db *appdbimpl) IncrementCommentCount(pid string) error {
	_, err := db.c.Exec("UPDATE posts SET comments = comments + 1 WHERE postId = ?;", pid)
	return err
}

func (db *appdbimpl) DecrementCommentCount(pid string) error {
	_, err := db.c.Exec("UPDATE posts SET comments = comments - 1 WHERE postId = ?;", pid)
	return err
}
