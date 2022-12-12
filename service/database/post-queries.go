package database

import "github.com/rhubinger/WASAgram/service/schemes"

func (db *appdbimpl) InsertPost(post schemes.Post) error {
	return nil
}

func (db *appdbimpl) DeletePost(pid string) error {
	return nil
}

func (db *appdbimpl) GetPictureId(pid string) (string, error) {
	return "pid", nil
}

func (db *appdbimpl) GetPost(pid string) (schemes.Post, error) {
	return schemes.Post{}, nil
}

func (db *appdbimpl) GetStream(uid string) ([]schemes.Post, error) {
	return []schemes.Post{}, nil
}

func (db *appdbimpl) GetPosts(uid string) ([]schemes.Post, error) {
	return []schemes.Post{}, nil
}

func (db *appdbimpl) GetPostCount(uid string) (int, error) {
	return -1, nil
}

func (db *appdbimpl) IncrementPostCount(uid string) error {
	return nil
}

func (db *appdbimpl) DecrementPostCount(uid string) error {
	return nil
}
