package sqlstore

import "github.com/Dennikoff/UserTagApi/internal/app/model"

type UserRepository struct {
	store *Store
}

func (r *UserRepository) Create(user *model.User) error {
	if err := r.store.db.QueryRow("INSERT INTO users (id, email, password, nickname) VALUES (default, $1, $2, $3) RETURNING id",
		user.Email, user.EncryptedPassword, user.NickName,
	).Scan(&user.ID); err != nil {
		return err
	}
	return nil
}
