package database

func (db *appdbimpl) AuthorizeAsUser(identifier string, userId string) (bool, error) {
	rightIdentifier, err := db.GetIdentifier(userId)
	if err != nil {
		return false, err
	}
	return identifier == rightIdentifier, nil
}

func (db *appdbimpl) AuthorizeAsNotBanned(identifier string, userId string) (bool, error) {
	bannedId, err := db.GetUserId(identifier)
	if err != nil {
		return false, err
	}

	banned, err := db.BanExists(userId, bannedId)
	if err != nil {
		return false, err
	}
	return !banned, nil
}
