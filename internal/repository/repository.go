package repository

import (
	"database/sql"
	"hickmet_whatapp_bot/models"
)

type Repository interface {
	GetUserByIIN(iin string) (models.User, error)

	GetSessionByPhone(phone string) (models.Session, error)
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}
