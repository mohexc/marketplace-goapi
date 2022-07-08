package auth

import (
	"github.com/google/uuid"
)

type UserRequest struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type UserResponse struct {
	ID    uuid.UUID `json:"id"`
	Email string    `json:"email"`
}

func CreateUser(userReq *UserRequest) (*UserResponse, error) {
	user := User{
		Email:    userReq.Email,
		Password: userReq.Password,
	}
	result := postgresqlDB.Create(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &UserResponse{
		ID:    user.ID,
		Email: user.Email,
	}, nil
}
