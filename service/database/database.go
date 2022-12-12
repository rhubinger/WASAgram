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
)

// AppDatabase is the high level interface for the DB
type AppDatabase interface {
	/*// User
	InsertUser(user api.User, identifier string) error

	GetIdentifier(uid string) (string, error)
	UpdateUsername(name string) error
	GetUser(uid string) (api.User, error)
	SearchUser(searchString string, searchType string) ([]api.User, error)

	// Posts
	InsertPost(post api.Post) error
	DeletePost(pid string) error

	GetPictureId(pid string) (string, error)
	GetPost(pid string) (api.Post, error)
	GetStream(uid string) ([]api.Post, error)
	GetPosts(uid string) ([]api.Post, error)
	GetPostCount(uid string) (int, error)
	IncrementPostCount(uid string) error
	DecrementPostCount(uid string) error

	// Comments
	InsertComment(comment api.Comment) error
	DeleteComments(pid string) error

	GetComments(pid string) ([]api.Comment, error)
	GetCommentCount(pid string) (int, error)
	IncrementCommentCount(pid string) error
	DecrementCommentCount(pid string) error

	// Follower
	Follow(uid string, fid string) error
	Unfollow(uid string, fid string) error

	GetFollowers(uid string) ([]api.User, error)
	IncrementFollowerCount(uid string) error
	DecrementFollowerCount(uid string) error
	GetFollowerCount(uid string) (int, error)

	GetFollowed(uid string) ([]api.User, error)
	GetFollowedCount(uid string) (int, error)
	IncrementFollowedCount(uid string) error
	DecrementFollowedCount(uid string) error

	// Banning
	Ban(uid string, bid string) error
	Unban(uid string, bid string) error

	GetBanned(uid string) ([]api.User, error)
	GetBannedCount(uid string) (int, error)

	// Likes
	Like(pid string, uid string) error
	Unlike(pid string, uid string) error
	DeleteLikes(pid string) error

	GetLikes(pid string) ([]api.User, error)
	GetLikeCount(pid string) (int, error)
	IncrementLikeCount(pid string) error
	DecrementLikeCount(pid string) error

	// Pictures
	InsertPicture(pid string, picture string) error // TODO get right datatype for images in BLOB
	GetPicture(pid string) (string, byte)
	DeletePicture(pid string) error
	*/
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
	// User table
	var userTable string
	err := db.QueryRow(`SELECT uid FROM sqlite_master WHERE type='table' AND name='users';`).Scan(&userTable)
	if errors.Is(err, sql.ErrNoRows) {
		sqlStmt := `CREATE TABLE example_table (id INTEGER NOT NULL PRIMARY KEY, name TEXT);` // TODO insert correct SQL statement
		_, err = db.Exec(sqlStmt)
		if err != nil {
			return nil, fmt.Errorf("error creating database structure: %w", err)
		}
	}
	// Posts table
	var postTable string
	err = db.QueryRow(`SELECT pid FROM sqlite_master WHERE type='table' AND name='posts';`).Scan(&postTable)
	if errors.Is(err, sql.ErrNoRows) {
		sqlStmt := `CREATE TABLE example_table (id INTEGER NOT NULL PRIMARY KEY, name TEXT);` // TODO insert correct SQL statement
		_, err = db.Exec(sqlStmt)
		if err != nil {
			return nil, fmt.Errorf("error creating database structure: %w", err)
		}
	}
	// Comments table
	var commentTable string
	err = db.QueryRow(`SELECT cid FROM sqlite_master WHERE type='table' AND name='comments';`).Scan(&commentTable)
	if errors.Is(err, sql.ErrNoRows) {
		sqlStmt := `CREATE TABLE example_table (id INTEGER NOT NULL PRIMARY KEY, name TEXT);` // TODO insert correct SQL statement
		_, err = db.Exec(sqlStmt)
		if err != nil {
			return nil, fmt.Errorf("error creating database structure: %w", err)
		}
	}
	// Followers table
	var followerTable string
	err = db.QueryRow(`SELECT uid FROM sqlite_master WHERE type='table' AND name='followers';`).Scan(&followerTable)
	if errors.Is(err, sql.ErrNoRows) {
		sqlStmt := `CREATE TABLE example_table (id INTEGER NOT NULL PRIMARY KEY, name TEXT);` // TODO insert correct SQL statement
		_, err = db.Exec(sqlStmt)
		if err != nil {
			return nil, fmt.Errorf("error creating database structure: %w", err)
		}
	}
	// Banned table
	var bannedTable string
	err = db.QueryRow(`SELECT uid FROM sqlite_master WHERE type='table' AND name='banned';`).Scan(&bannedTable)
	if errors.Is(err, sql.ErrNoRows) {
		sqlStmt := `CREATE TABLE example_table (id INTEGER NOT NULL PRIMARY KEY, name TEXT);` // TODO insert correct SQL statement
		_, err = db.Exec(sqlStmt)
		if err != nil {
			return nil, fmt.Errorf("error creating database structure: %w", err)
		}
	}
	// Likkes table
	var likeTable string
	err = db.QueryRow(`SELECT pid FROM sqlite_master WHERE type='table' AND name='likes';`).Scan(&likeTable)
	if errors.Is(err, sql.ErrNoRows) {
		sqlStmt := `CREATE TABLE example_table (id INTEGER NOT NULL PRIMARY KEY, name TEXT);` // TODO insert correct SQL statement
		_, err = db.Exec(sqlStmt)
		if err != nil {
			return nil, fmt.Errorf("error creating database structure: %w", err)
		}
	}
	// Picture table
	var pictureTable string
	err = db.QueryRow(`SELECT pid FROM sqlite_master WHERE type='table' AND name='likes';`).Scan(&pictureTable)
	if errors.Is(err, sql.ErrNoRows) {
		sqlStmt := `CREATE TABLE example_table (id INTEGER NOT NULL PRIMARY KEY, name TEXT);` // TODO insert correct SQL statement
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
