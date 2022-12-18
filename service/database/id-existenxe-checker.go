package database

import (
	"database/sql"
	"errors"
	"os"
)

func (db *appdbimpl) UserExists(uid string) (bool, error) {
	_, err := db.GetUser(uid)
	if errors.Is(err, sql.ErrNoRows) {
		os.Stderr.WriteString("UidExists: uid doesn't exist: " + err.Error())
		return false, nil
	} else if err != nil {
		os.Stderr.WriteString(err.Error())
		return false, err
	}
	return true, nil
}

func (db *appdbimpl) PostExists(pid string) (bool, error) {
	_, err := db.GetPost(pid)
	if errors.Is(err, sql.ErrNoRows) {
		os.Stderr.WriteString("PidExists: pid doesn't exist: " + err.Error())
		return false, nil
	} else if err != nil {
		os.Stderr.WriteString(err.Error())
		return false, err
	}
	return true, nil
}

func (db *appdbimpl) CommentExists(cid string) (bool, error) {
	_, err := db.GetComment(cid)
	if errors.Is(err, sql.ErrNoRows) {
		os.Stderr.WriteString("CidExists: cid doesn't exist: " + err.Error())
		return false, nil
	} else if err != nil {
		os.Stderr.WriteString(err.Error())
		return false, err
	}
	return true, nil
}

func (db *appdbimpl) PictureExists(pid string) (bool, error) {
	var pictureId string
	err := db.c.QueryRow(`SELECT pictureId
						FROM pictures 
						WHERE pictureId = ?;`, pid).Scan(&pictureId)
	if errors.Is(err, sql.ErrNoRows) {
		return false, nil
	} else if err != nil {
		return true, err
	}
	return true, nil
}

func (db *appdbimpl) FollowExists(uid string, fid string) (bool, error) {
	var userId string
	var followerId string
	err := db.c.QueryRow(`SELECT userId, followerId
						FROM followers 
						WHERE userId = ? AND followerId = ?;`, uid, fid).Scan(&userId, &followerId)
	if errors.Is(err, sql.ErrNoRows) {
		return false, nil
	} else if err != nil {
		return true, err
	}
	return true, nil
}

func (db *appdbimpl) BanExists(uid string, bid string) (bool, error) {
	var userId string
	var bannedId string
	err := db.c.QueryRow(`SELECT userId, bannedId
						FROM banned
						WHERE userId = ? AND bannedId = ?;`, uid, bid).Scan(&userId, &bannedId)
	if errors.Is(err, sql.ErrNoRows) {
		return false, nil
	} else if err != nil {
		return true, err
	}
	return true, nil
}

func (db *appdbimpl) LikeExists(pid string, uid string) (bool, error) {
	var postId string
	var userId string
	err := db.c.QueryRow(`SELECT postId, userId
						FROM likes 
						WHERE postId = ? AND userId = ?;`, pid, uid).Scan(&postId, &userId)
	if errors.Is(err, sql.ErrNoRows) {
		return false, nil
	} else if err != nil {
		return true, err
	}
	return true, nil
}
