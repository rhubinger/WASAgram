package database

import "github.com/rhubinger/WASAgram/service/schemes"

func (db *appdbimpl) GetComments(pid string) ([]schemes.Comment, error) {
	var name string
	err := db.c.QueryRow("SELECT name FROM example_table WHERE id=1").Scan(&name)
	return []schemes.Comment{}, err
}
