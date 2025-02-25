package repositories

import (
	"context"
	"errors"

	"backend.com/go-backend/ent"
	"backend.com/go-backend/ent/user"
	"github.com/alexedwards/argon2id"
)

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
func CreateUserRepository(entClient *ent.Client, data *ent.User) error {
	ctx := context.Background()

	exists, err := entClient.User.Query().Where(user.Or(user.EmailEQ(data.Email), user.UsernameEQ(data.Username))).Exist(ctx)
	if err != nil {
		return err
	}

	if exists {
		return errors.New("user with the given email or username already exists")
	}

	// Hash the password
	hash, err := argon2id.CreateHash(data.Password, argon2id.DefaultParams)
	if err != nil {
		return errors.New("failed to hash password")
	}

	// Start a transaction
	tx, err := entClient.Tx(ctx)
	if err != nil {
		return err
	}

	_, err = tx.User.Create().SetAvatar(data.Avatar).SetEmail(data.Email).SetUsername(data.Username).SetFullName(data.FullName).SetPassword(hash).SetStartDate(data.StartDate).SetIsStaff(data.IsStaff).SetIsActive(data.IsActive).SetProvider(data.Provider).SetProviderID(data.ProviderID).Save(context.Background())
	if err != nil {
		tx.Rollback()
		return errors.New("failed to create user" + err.Error())
	}

	// Commit the transaction
	if err := tx.Commit(); err != nil {
		return errors.New("failed to commit transaction")
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
// func GetUserRepository(identifier string) (models.User, error) {
// 	var user models.User

// 	if err := config.DB.Session(&gorm.Session{PrepareStmt: false}).Select("id", "avatar", "email", "username", "full_name", "password", "start_date", "is_staff", "is_active", "provider", "provider_id").Where("email = ? OR provider_id = ?", identifier, identifier).First(&user).Error; err != nil {
// 		if err == gorm.ErrRecordNotFound {
// 			return user, errors.New("user not found")
// 		}
// 		return user, errors.New("failed to get user")
// 	}
// 	return user, nil
// }
