package repository

import (
	"database/sql"
	"fmt"
	"hickmet_whatapp_bot/consts"
	"hickmet_whatapp_bot/models"
)

func (r *repository) GetSessionByPhone(phone string) (models.Session, error) {
	query := `SELECT phone, status FROM sessions WHERE phone = $1`
	session := &models.Session{}

	err := r.db.QueryRow(query, phone).Scan(
		&session.Phone,
		&session.Status,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return r.CreateSession(phone)
		}
		return models.Session{}, err
	}

	return *session, nil
}

func (r *repository) CreateSession(phone string) (models.Session, error) {
	var sessionID int
	query := `
        INSERT INTO sessions (phone, status, created_at, updated_at)
        VALUES ($1, $2, NOW(), NOW())
    `
	err := r.db.QueryRow(query, phone, consts.Start).Scan(&sessionID)
	if err != nil {
		return models.Session{}, fmt.Errorf("ошибка создания сессии: %v", err)
	}

	return models.Session{Phone: phone, Status: consts.Start}, nil
}
