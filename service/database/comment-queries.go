package database

import "github.com/rhubinger/WASAgram/service/schemes"

func (db *appdbimpl) InsertComment(comment schemes.Comment) error {
	return nil
}

func (db *appdbimpl) DeleteComments(pid string) error {
	return nil
}

func (db *appdbimpl) GetComments(pid string) ([]schemes.Comment, error) {
	return []schemes.Comment{}, nil
}

func (db *appdbimpl) GetCommentCount(pid string) (int, error) {
	return -1, nil
}

func (db *appdbimpl) IncrementCommentCount(pid string) error {
	return nil
}

func (db *appdbimpl) DecrementCommentCount(pid string) error {
	return nil
}
