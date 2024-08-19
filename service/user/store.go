package user

import (
	"database/sql"
	"fmt"
	"go-api/types"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db}
}

func (store *Store) GetUserByEmail(email string) (*types.User, error) {
	rows, err := store.db.Query("SELECT * FROM users WHERE email = $1", email)

	if err != nil {
		return nil, err
	}

	user := new(types.User)
	found := false

	for rows.Next() {
		err = scanRowIntoUser(rows, user)
		if err != nil {
			return nil, err
		}
		found = true
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	if !found {
		return nil, fmt.Errorf("user not found")
	}

	return user, nil
}

func scanRowIntoUser(rows *sql.Rows, user *types.User) error {
	return rows.Scan(&user.ID, &user.Email, &user.Name, &user.CreatedAt, &user.UpdatedAt)
}

func (store *Store) GetUserById(id string) (*types.User, error) {
	rows, err := store.db.Query("SELECT * FROM users WHERE id = $1", id)
	if err != nil {
		return nil, err
	}

	user := new(types.User)
	for rows.Next() {
		err = scanRowIntoUser(rows, user)
		if err != nil {
			return nil, err
		}
	}

	if user.ID == "" {
		return nil, fmt.Errorf("user not found")
	}

	return user, nil
}

func (store *Store) CreateUser(user types.User) error {
	_, err := store.db.Exec("INSERT INTO users (name, email, role, password) VALUES ($1, $2, $3, $4)", user.Name, user.Email, user.Role, user.Password)
	if err != nil {
		return err
	}

	return nil
}
