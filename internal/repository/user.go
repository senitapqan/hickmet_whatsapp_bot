package repository

import "hickmet_whatapp_bot/models"

func (r *repository) GetUserByIIN(iin string) (models.User, error) {
	return models.User{}, nil
}
