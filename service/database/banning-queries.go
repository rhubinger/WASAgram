package database

import "github.com/rhubinger/WASAgram/service/schemes"

func (db *appdbimpl) Ban(uid string, bid string) error {
	return nil
}

func (db *appdbimpl) Unban(uid string, bid string) error {
	return nil
}

func (db *appdbimpl) GetBanned(uid string) ([]schemes.User, error) {
	return []schemes.User{}, nil
}

func (db *appdbimpl) GetBannedCount(uid string) (int, error) {
	return -1, nil
}
