package database

import "github.com/rhubinger/WASAgram/service/schemes"

func (db *appdbimpl) InsertUser(user schemes.User, identifier string) error {
	return nil
}

func (db *appdbimpl) GetIdentifier(uid string) (string, error) {
	return "user", nil
}

func (db *appdbimpl) UpdateUsername(name string) error {
	return nil
}

func (db *appdbimpl) GetUser(uid string) (schemes.User, error) {
	return schemes.User{}, nil
}

func (db *appdbimpl) SearchUser(searchString string, searchType string) ([]schemes.User, error) {
	return []schemes.User{}, nil
}
