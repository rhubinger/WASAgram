package database

import "github.com/rhubinger/WASAgram/service/schemes"

func (db *appdbimpl) Follow(uid string, fid string) error {
	return nil
}

func (db *appdbimpl) Unfollow(uid string, fid string) error {
	return nil
}

func (db *appdbimpl) GetFollowers(uid string) ([]schemes.User, error) {
	return []schemes.User{}, nil
}

func (db *appdbimpl) GetFollowerCount(uid string) (int, error) {
	return -1, nil
}

func (db *appdbimpl) IncrementFollowerCount(uid string) error {
	return nil
}

func (db *appdbimpl) DecrementFollowerCount(uid string) error {
	return nil
}

func (db *appdbimpl) GetFollowed(uid string) ([]schemes.User, error) {
	return []schemes.User{}, nil
}

func (db *appdbimpl) GetFollowedCount(uid string) (int, error) {
	return -1, nil
}

func (db *appdbimpl) IncrementFollowedCount(uid string) error {
	return nil
}

func (db *appdbimpl) DecrementFollowedCount(uid string) error {
	return nil
}
