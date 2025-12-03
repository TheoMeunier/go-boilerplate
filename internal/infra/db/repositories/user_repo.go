package repositories

import "boilerplate/internal/infra/db"

type UserRepo struct {
	*db.BaseRepo // embedding
}

type User struct {
	ID   int
	Name string
}

func (r *UserRepo) GetUserByID(id int) (*User, error) {
	user := &User{}
	err := r.DB.Conn.Model(user).Where("id = ?", id).Select()

	if err != nil {
		return nil, err
	}

	return user, nil
}
