package database

import "github.com/rhubinger/WASAgram/service/schemes"

func (db *appdbimpl) Like(pid string, uid string) error {
	return nil
}

func (db *appdbimpl) Unlike(pid string, uid string) error {
	return nil
}

func (db *appdbimpl) DeleteLikes(pid string) error {
	return nil
}

func (db *appdbimpl) GetLikes(pid string) ([]schemes.User, error) {
	return []schemes.User{}, nil
}

func (db *appdbimpl) GetLikeCount(pid string) (int, error) {
	return -1, nil
}

func (db *appdbimpl) IncrementLikeCount(pid string) error {
	return nil
}

func (db *appdbimpl) DecrementLikeCount(pid string) error {
	return nil
}
