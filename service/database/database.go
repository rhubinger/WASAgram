/*
Package database is the middleware between the app database and the code. All data (de)serialization (save/load) from a
persistent database are handled here. Database specific logic should never escape this package.

To use this package you need to apply migrations to the database if needed/wanted, connect to it (using the database
data source name from config), and then initialize an instance of AppDatabase from the DB connection.

For example, this code adds a parameter in `webapi` executable for the database data source name (add it to the
main.WebAPIConfiguration structure):

	DB struct {
		Filename string `conf:""`
	}

This is an example on how to migrate the DB and connect to it:

	// Start Database
	logger.Println("initializing database support")
	db, err := sql.Open("sqlite3", "./foo.db")
	if err != nil {
		logger.WithError(err).Error("error opening SQLite DB")
		return fmt.Errorf("opening SQLite: %w", err)
	}
	defer func() {
		logger.Debug("database stopping")
		_ = db.Close()
	}()

Then you can initialize the AppDatabase and pass it to the api package.
*/
package database

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/rhubinger/WASAgram/service/schemes"
)

// AppDatabase is the high level interface for the DB
type AppDatabase interface {
	// User
	InsertUser(user schemes.User) (string, error)

	GetIdentifier(uid string) (string, error)
	UpdateUsername(name string, uid string) error
	GetUser(uid string) (schemes.User, error)
	SearchUser(searchString string, uid string) ([]schemes.User, error)

	// Posts
	InsertPost(post schemes.Post) (string, error)
	DeletePost(pid string) error

	GetPictureId(pid string) (string, error)
	GetPost(pid string) (schemes.Post, error)
	GetStream(uid string, dateTime string) ([]schemes.Post, error)
	GetPosts(uid string, dateTime string) ([]schemes.Post, error)
	GetPostCount(uid string) (int, error)
	IncrementPostCount(uid string) error
	DecrementPostCount(uid string) error

	// Comments
	InsertComment(comment schemes.Comment) (string, error)
	DeleteComment(cid string) error
	DeleteComments(pid string) error

	GetComments(pid string) ([]schemes.Comment, error)
	GetCommentCount(pid string) (int, error)
	IncrementCommentCount(pid string) error
	DecrementCommentCount(pid string) error

	// Follower
	Follow(uid string, fid string) error
	Unfollow(uid string, fid string) error

	GetFollowers(uid string) ([]schemes.User, error)
	GetFollowerCount(uid string) (int, error)
	IncrementFollowerCount(uid string) error
	DecrementFollowerCount(uid string) error

	GetFollowed(uid string) ([]schemes.User, error)
	GetFollowedCount(uid string) (int, error)
	IncrementFollowedCount(uid string) error
	DecrementFollowedCount(uid string) error

	// Banning
	Ban(uid string, bid string) error
	Unban(uid string, bid string) error

	GetBanned(uid string) ([]schemes.User, error)
	GetBannedCount(uid string) (int, error)

	// Likes
	Like(pid string, uid string) error
	Unlike(pid string, uid string) error
	DeleteLikes(pid string) error

	GetLikes(pid string) ([]schemes.User, error)
	GetLikeCount(pid string) (int, error)
	IncrementLikeCount(pid string) error
	DecrementLikeCount(pid string) error

	// Pictures
	InsertPicture(picture []byte) (string, error)
	GetPicture(pid string) ([]byte, error)
	DeletePicture(pid string) error

	// Generates unique Ids for posts, comments and pictures
	GenerateId(idType string) string

	// Ping
	Ping() error
}

type appdbimpl struct {
	c *sql.DB
}

// New returns a new instance of AppDatabase based on the SQLite connection `db`.
// `db` is required - an error will be returned if `db` is `nil`.
func New(db *sql.DB) (AppDatabase, error) {
	if db == nil {
		return nil, errors.New("database is required when building a AppDatabase")
	}

	// Check if tables exist. If not, the tables get created
	var str string
	// User table
	err := db.QueryRow(`SELECT * FROM sqlite_master WHERE type='table' AND name='users';`).Scan(&str)
	if errors.Is(err, sql.ErrNoRows) {
		sqlStmt := `CREATE TABLE users (
						userId VARCHAR(16) NOT NULL PRIMARY KEY,
						name VARCHAR(30) NOT NULL, 
						identifier CHAR(11) NOT NULL,
						posts INTEGER NOT NULL,
						followers INTEGER NOT NULL,
						followed INTEGER NOT NULL
					);`
		_, err = db.Exec(sqlStmt)
		if err != nil {
			return nil, fmt.Errorf("error creating database structure: %w", err)
		}
	}
	// Posts table
	err = db.QueryRow(`SELECT * FROM sqlite_master WHERE type='table' AND name='posts';`).Scan(&str)
	if errors.Is(err, sql.ErrNoRows) {
		sqlStmt := `CREATE TABLE posts (
						postId CHAR(11) NOT NULL PRIMARY KEY,
						userId CHAR(11),
						uploadTime DATETIME,
						caption VARCHAR(140),
						pictureId CHAR(11),
						likes INTEGER,
						comments INTEGER
					);`
		_, err = db.Exec(sqlStmt)
		if err != nil {
			return nil, fmt.Errorf("error creating database structure: %w", err)
		}
	}
	// Comments table
	err = db.QueryRow(`SELECT * FROM sqlite_master WHERE type='table' AND name='comments';`).Scan(&str)
	if errors.Is(err, sql.ErrNoRows) {
		sqlStmt := `CREATE TABLE comments (
						commentId CHAR(11) NOT NULL PRIMARY KEY,
						userId CHAR(11),
						postId CHAR(11),
						uploadTime DATETIME,
						commentText VARCHAR(140)
					);`
		_, err = db.Exec(sqlStmt)
		if err != nil {
			return nil, fmt.Errorf("error creating database structure: %w", err)
		}
	}
	// Followers table
	err = db.QueryRow(`SELECT * FROM sqlite_master WHERE type='table' AND name='followers';`).Scan(&str)
	if errors.Is(err, sql.ErrNoRows) {
		sqlStmt := `CREATE TABLE followers (
						userId CHAR(11) NOT NULL,
						followerId CHAR(11) NOT NULL,
						PRIMARY KEY (userId, followerId)
					);`
		_, err = db.Exec(sqlStmt)
		if err != nil {
			return nil, fmt.Errorf("error creating database structure: %w", err)
		}
	}
	// Banned table
	err = db.QueryRow(`SELECT * FROM sqlite_master WHERE type='table' AND name='bans';`).Scan(&str)
	if errors.Is(err, sql.ErrNoRows) {
		sqlStmt := `CREATE TABLE bans (
						userId CHAR(11) NOT NULL,
						bannedId CHAR(11) NOT NULL,
						PRIMARY KEY (userId, bannedId)
					);`
		_, err = db.Exec(sqlStmt)
		if err != nil {
			return nil, fmt.Errorf("error creating database structure: %w", err)
		}
	}
	// Likkes table
	err = db.QueryRow(`SELECT * FROM sqlite_master WHERE type='table' AND name='likes';`).Scan(&str)
	if errors.Is(err, sql.ErrNoRows) {
		sqlStmt := `CREATE TABLE likes (
						postId CHAR(11) NOT NULL,
						userId CHAR(11) NOT NULL,
						PRIMARY KEY (postId, userId)
					);`
		_, err = db.Exec(sqlStmt)
		if err != nil {
			return nil, fmt.Errorf("error creating database structure: %w", err)
		}
	}
	// Picture table
	err = db.QueryRow(`SELECT * FROM sqlite_master WHERE type='table' AND name='pictures';`).Scan(&str)
	if errors.Is(err, sql.ErrNoRows) {
		sqlStmt := `CREATE TABLE pictures (
						pictureId CHAR(11) NOT NULL PRIMARY KEY,
						picture BLOB
					);`
		_, err = db.Exec(sqlStmt)
		if err != nil {
			return nil, fmt.Errorf("error creating database structure: %w", err)
		}
	}

	return &appdbimpl{
		c: db,
	}, nil
}

func (db *appdbimpl) Ping() error {
	return db.c.Ping()
}
