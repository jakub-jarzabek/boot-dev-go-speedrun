package database

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (db *DB) CreateUser(email, password string) (User, error) {
	dbStructure, err := db.loadDB()
	if err != nil {
		return User{}, err
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return User{}, err
	}

	id := len(dbStructure.Users) + 1
	user := User{
		ID:       id,
		Email:    email,
		Password: string(hash),
	}

	dbStructure.Users[id] = user

	err = db.writeDB(dbStructure)
	if err != nil {
		return User{}, err
	}

	return user, nil
}

func (db *DB) GetUser(id int) (User, error) {
	dbStructure, err := db.loadDB()
	if err != nil {
		return User{}, err
	}

	user, ok := dbStructure.Users[id]
	if !ok {
		return User{}, ErrNotExist
	}

	return user, nil
}

func (db *DB) GetUserByEmail(email string) (User, error) {
	dbStructure, err := db.loadDB()
	if err != nil {
		return User{}, err
	}

	for _, v := range dbStructure.Users {
		if v.Email == email {
			return v, nil
		}
	}

	return User{}, errors.New("User not found")
}
