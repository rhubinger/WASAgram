package database

import "github.com/rhubinger/WASAgram/service/schemes"

// GetName is an example that shows you how to query data
func (db *appdbimpl) GetBanned(uid string) ([]schemes.User, error) {
	var name string
	err := db.c.QueryRow("SELECT name FROM example_table WHERE id=1").Scan(&name)
	return []schemes.User{}, err
}
