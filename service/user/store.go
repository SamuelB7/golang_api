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

	for rows.Next() {
		user, err = scanRowIntoUser(rows)
		if err != nil {
			return nil, err
		}
	}
	if user.ID == "" {
		return nil, fmt.Errorf("user not found")
	}
	return user, nil
}

func scanRowIntoUser(rows *sql.Rows) (*types.User, error) {
	user := new(types.User)
	err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.Role, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (store *Store) GetUserById(id string) (*types.User, error) {
	rows, err := store.db.Query("SELECT * FROM users WHERE id = $1", id)
	if err != nil {
		return nil, err
	}

	user := new(types.User)
	for rows.Next() {
		user, err = scanRowIntoUser(rows)
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
