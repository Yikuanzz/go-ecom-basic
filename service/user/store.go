package user

import (
	"database/sql"
	"fmt"

	"github.com/yikuanzz/ecom/model"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}
func (s *Store) GetUserByID(id int) (*model.User, error) {
	return nil, nil
}

func (s *Store) CreateUser(user model.User) error {
	return nil
}

func (s *Store) GetUserByEmail(email string) (*model.User, error) {
	rows, err := s.db.Query("SELECT * FROM user WHERE email = ?", email)
	if err != nil {
		return nil, err
	}

	user := &model.User{}
	if rows.Next() {
		err := scanRowIntoUser(rows, user)
		if err != nil {
			return nil, err
		}
	} else {
		return nil, fmt.Errorf("user not found")
	}
	return user, nil
}

func scanRowIntoUser(rows *sql.Rows, user *model.User) error {
	err := rows.Scan(
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
	)
	return err
}
