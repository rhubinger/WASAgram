package database

import (
	"math/rand"
)

func GenerateRandomString(length int) string {
	chars := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ-_"
	str := make([]byte, length)
	for i := range str {
		str[i] = chars[rand.Intn(len(chars))]
	}
	return string(str)
}

func (db *appdbimpl) GenerateId(idType string) string {
	// Infer database name from idType
	var table string
	switch idType {
	case "userId":
		table = "users"
	case "postId":
		table = "posts"
	case "commentId":
		table = "comments"
	case "pictureId":
		table = "pictures"
	default:
		return GenerateRandomString(11)
	}

	id := GenerateRandomString(11)
	var result string
	err := db.c.QueryRow("SELECT ? FROM ? WHERE ? == ?", idType, table, idType, id).Scan(&result)
	for err == nil {
		id = GenerateRandomString(11)
		err = db.c.QueryRow("SELECT ? FROM ? WHERE ? == ?", idType, table, idType, id).Scan(&result)
	}

	return id
}

func (db *appdbimpl) GenerateIdentifier() string {
	id := GenerateRandomString(11)
	var result string
	err := db.c.QueryRow("SELECT identifier FROM users WHERE identifier == ?", id).Scan(&result)
	for err == nil {
		id = GenerateRandomString(11)
		err = db.c.QueryRow("SELECT identifier FROM users WHERE identifier == ?", id).Scan(&result)
	}

	return id
}
