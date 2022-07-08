package auth

import (
	"errors"
	"marketplace-goapi/pkg/base"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type CreateUserRequest struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type UserResponse struct {
	ID    uuid.UUID `json:"id"`
	Email string    `json:"email"`
}

var validate = validator.New()

func ValidateStruct(userReq CreateUserRequest) []*base.ErrorResponse {
	var errors []*base.ErrorResponse
	err := validate.Struct(userReq)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element base.ErrorResponse
			element.FailedField = err.StructNamespace()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, &element)
		}
	}
	return errors
}

func CreateUser(userReq *CreateUserRequest) (*UserResponse, error) {
	user := User{
		Email:    userReq.Email,
		Password: userReq.Password,
	}
	result := postgresqlDB.Where("email = ?", user.Email).First(&User{})
	if result.Error == nil {
		return nil, errors.New("User already exists")
	}
	result = postgresqlDB.Create(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &UserResponse{
		ID:    user.ID,
		Email: user.Email,
	}, nil
}

func GetUsers() (*[]UserResponse, error) {
	var users []User
	result := postgresqlDB.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	var usersResp []UserResponse
	for _, user := range users {
		usersResp = append(usersResp, UserResponse{
			ID:    user.ID,
			Email: user.Email,
		})
	}
	return &usersResp, nil
}

func GetUserById(id uuid.UUID) (*UserResponse, error) {
	var user User
	result := postgresqlDB.First(&user, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &UserResponse{
		ID:    user.ID,
		Email: user.Email,
	}, nil
}

func UpdateUserById(id uuid.UUID, userReq *CreateUserRequest) (*UserResponse, error) {
	result := postgresqlDB.Model(&User{}).Where("id = ?", id).Updates(User{
		Email:    userReq.Email,
		Password: userReq.Password,
	})
	if result.Error != nil {
		return nil, result.Error
	}
	return &UserResponse{
		ID:    id,
		Email: userReq.Email,
	}, nil
}

func DeleteUserById(id uuid.UUID) error {
	var user User
	result := postgresqlDB.First(&user, id)
	if result.Error != nil {
		return result.Error
	}
	result = postgresqlDB.Delete(&user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
