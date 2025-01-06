package repositories

import (
	"errors"

	"backend.com/go-backend/src/config"
	"backend.com/go-backend/src/models"
	"github.com/alexedwards/argon2id"
	"gorm.io/gorm"
)

type CreateUserInput struct {
	Avatar   string `json:"avatar"`
	Email    string `json:"email" binding:"required"`
	Username string `json:"username" binding:"required"`
	FullName string `json:"full_name" binding:"required"`
	Password string `json:"password" binding:"required"`
	IsStaff  bool   `json:"is_staff"`
}

// CreateUserRepository creates a new user in the database.
// It first checks if a user with the given email or username already exists.
// If such a user exists, it returns an error.
// If the user does not exist, it hashes the password and creates a new user record in the database.
//
// Parameters:
//   - data: CreateUserInput containing the user's details.
//
// Returns:
//   - error: An error if the user already exists, if password hashing fails, or if the user creation fails.
func CreateUserRepository(data CreateUserInput) error {
	// check if user already exists
	var existingUser models.User
	if err := config.DB.Where("email = ? OR username = ?", data.Email, data.Username).
		First(&existingUser).Error; err == nil {
		return errors.New("user with the given email or username already exists")
	}

	// Create a new user
	hashedPassword, err := argon2id.CreateHash(data.Password, argon2id.DefaultParams)
	println("Generated Hash:", hashedPassword)
	if err != nil {
		return errors.New("failed to hash password")
	}

	user := models.User{
		Avatar:   data.Avatar,
		Email:    data.Email,
		Username: data.Username,
		FullName: data.FullName,
		Password: hashedPassword,
		IsStaff:  data.IsStaff,
	}

	if err := config.DB.Create(&user).Error; err != nil {
		return errors.New("failed to create user")
	}
	return nil
}

// GetUserRepository retrieves a user from the database based on the provided email.
// It returns a models.User object and an error. If the user is not found, it returns
// an error indicating that the user was not found. If there is any other issue during
// the retrieval process, it returns a generic error indicating the failure to get the user.
//
// Parameters:
//   - email: The email of the user to be retrieved.
//
// Returns:
//   - models.User: The user object containing the user's details.
//   - error: An error object if there is an issue during the retrieval process.
func GetUserRepository(email string) (models.User, error) {
	var user models.User

	if err := config.DB.Session(&gorm.Session{PrepareStmt: false}).Select("id", "avatar", "email", "username", "password", "full_name", "start_date", "is_staff", "is_active").Where("email = ?", email).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return user, errors.New("user not found")
		}
		return user, errors.New("failed to get user")
	}
	return user, nil
}
