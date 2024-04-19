package database

func (db *DB) AddRevokedToken(token string) (bool, error) {
	dbStructure, err := db.loadDB()
	if err != nil {
		return false, err
	}

	tokens := dbStructure.RevokedTokens

	for _, t := range tokens {
		if t == token {
			return false, nil
		}
	}

	tokens = append(tokens, token)
	dbStructure.RevokedTokens = tokens
	err = db.writeDB(dbStructure)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (db *DB) IsTokenRevoked(token string) (bool, error) {
	dbStructure, err := db.loadDB()
	if err != nil {
		return false, err
	}

	tokens := dbStructure.RevokedTokens

	for _, t := range tokens {
		if t == token {
			return true, nil
		}
	}

	return false, nil
}
