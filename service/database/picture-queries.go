package database

// GetName is an example that shows you how to query data
func (db *appdbimpl) GetPicture(pid string) (string, error) {
	var name string
	err := db.c.QueryRow("SELECT name FROM example_table WHERE id=1").Scan(&name)
	return "picture", err
}
