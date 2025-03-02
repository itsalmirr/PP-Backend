package repositories

import (
	"context"
	"errors"

	"github.com/alexedwards/argon2id"
	"ppgroup.i0sys.com/ent"
	"ppgroup.i0sys.com/ent/user"
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
func CreateUserRepo(entClient *ent.Client, data *ent.User) error {
	ctx := context.Background()

	exists, err := entClient.User.Query().Where(user.Or(user.EmailEQ(data.Email), user.UsernameEQ(data.Username))).Exist(ctx)
	if err != nil {
		return err
	}

	if exists {
		return errors.New("user with the given email or username already exists")
	}

	// Hash the password
	hashedPassword, err := argon2id.CreateHash(data.Password, argon2id.DefaultParams)
	if err != nil {
		return errors.New("failed to hash password")
	}

	// Start a transaction
	tx, err := entClient.Tx(ctx)
	if err != nil {
		return err
	}

	_, err = tx.User.Create().
		SetAvatar(data.Avatar).
		SetEmail(data.Email).
		SetUsername(data.Username).
		SetFullName(data.FullName).
		SetPassword(hashedPassword).
		SetStartDate(data.StartDate).
		SetIsStaff(data.IsStaff).
		SetIsActive(data.IsActive).
		SetProvider(data.Provider).
		SetProviderID(data.ProviderID).
		Save(context.Background())
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
func GetUserRepo(entClient *ent.Client, identifier string) (*ent.User, error) {
	ctx := context.Background()

	user, err := entClient.User.Query().Where(user.EmailEQ(identifier)).Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, errors.New("user not found")
		}
		return nil, errors.New("failed to get user")
	}

	return user, nil
}
