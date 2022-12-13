package database

func (db *appdbimpl) InsertPicture(picture []byte) (string, error) {
	pid := db.GenerateId("pictureId")
	_, err := db.c.Exec("INSERT INTO pictures VALUES (?, ?);", pid, picture)
	return pid, err
}

func (db *appdbimpl) GetPicture(pid string) ([]byte, error) {
	var picture []byte
	err := db.c.QueryRow("SELECT picture FROM pictures WHERE pictureId == ?", pid).Scan(&picture)
	return picture, err
}

func (db *appdbimpl) DeletePicture(pid string) error {
	_, err := db.c.Exec("DELETE FROM pictures WHERE pictureId == ?", pid)
	return err
}
