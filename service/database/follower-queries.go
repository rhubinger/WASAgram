package database

import "github.com/rhubinger/WASAgram/service/schemes"

func (db *appdbimpl) GetFollowers(uid string) ([]schemes.User, error) {
	var name string
	err := db.c.QueryRow("SELECT name FROM example_table WHERE id=1").Scan(&name)
	return []schemes.User{}, err
}
