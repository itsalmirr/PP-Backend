package repositories

import (
	"context"
	"errors"
	"fmt"

	"github.com/alexedwards/argon2id"
	"ppgroup.ppgroup.com/ent"
	"ppgroup.ppgroup.com/ent/user"
)

// CreateUserInput represents the input data for creating a new user
type CreateUserInput struct {
	Email    string `json:"email" binding:"required,email"`
	Username string `json:"username" binding:"required,min=3"`
	FullName string `json:"full_name" binding:"required"`
	Password string `json:"password" binding:"required,min=6"` // Plain text password - will be hashed
	Avatar   string `json:"avatar,omitempty"`
}

func CreateUserRepo(ctx context.Context, entClient *ent.Client, data *CreateUserInput) error {
	exists, err := entClient.User.Query().Where(user.Or(user.EmailEQ(data.Email), user.UsernameEQ(data.Username))).Exist(ctx)
	if err != nil {
		return fmt.Errorf("checking user existence: %w", err)
	}

	if exists {
		return errors.New("user with the given email or username already exists")
	}

	// Hash the password
	hashedPassword, err := argon2id.CreateHash(data.Password, argon2id.DefaultParams)
	if err != nil {
		return fmt.Errorf("hashing password: %w", err)
	}

	// Start a transaction
	tx, err := entClient.Tx(ctx)
	if err != nil {
		return fmt.Errorf("starting transaction: %w", err)
	}

	// For email-based users, use email as provider_id to ensure uniqueness
	providerToUse := "email"
	providerIDToUse := data.Email

	_, err = tx.User.Create().
		SetAvatar(data.Avatar).
		SetEmail(data.Email).
		SetUsername(data.Username).
		SetFullName(data.FullName).
		SetPassword(hashedPassword).
		SetIsStaff(false).
		SetIsActive(true).
		SetProvider(providerToUse).
		SetProviderID(providerIDToUse).
		Save(ctx)
	if err != nil {
		if rerr := tx.Rollback(); rerr != nil {
			return fmt.Errorf("rollback failed: %v (original: %w)", rerr, err)
		}
		return fmt.Errorf("creating user: %w", err)
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("committing transaction: %w", err)
	}

	return nil
}

func GetUserRepo(ctx context.Context, entClient *ent.Client, identifier string) (*ent.User, error) {
	u, err := entClient.User.Query().Where(user.EmailEQ(identifier)).Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, errors.New("user not found")
		}
		return nil, fmt.Errorf("querying user: %w", err)
	}

	return u, nil
}
