package usersRepository

import (
	"database/sql"

	"users/usersModels"
)

type IRepository interface {
	PotsRegister(register usersModels.RequestRegister) error
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) IRepository {
	return &repository{db: db}
}

func (r *repository) PotsRegister(register usersModels.RequestRegister) error {
	

	return nil
}
