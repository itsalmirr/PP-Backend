package repositories

import (
	"context"
	"errors"
	"fmt"

	"ppgroup.ppgroup.com/ent"
	"ppgroup.ppgroup.com/ent/realtor"
)

func CreateRealtorRepo(ctx context.Context, entClient *ent.Client, data *ent.Realtor) error {
	exists, err := entClient.Realtor.Query().Where(realtor.Or(realtor.EmailEQ(data.Email), realtor.PhoneEQ(data.Phone))).Exist(ctx)
	if err != nil {
		return fmt.Errorf("checking realtor existence: %w", err)
	}

	if exists {
		return errors.New("realtor with the given email or phone already exists")
	}

	// Start a transaction
	tx, err := entClient.Tx(ctx)
	if err != nil {
		return fmt.Errorf("starting transaction: %w", err)
	}

	_, err = tx.Realtor.Create().SetEmail(data.Email).SetFullName(data.FullName).SetPhone(data.Phone).SetIsMvp(data.IsMvp).Save(ctx)
	if err != nil {
		if rerr := tx.Rollback(); rerr != nil {
			return fmt.Errorf("rollback failed: %v (original: %w)", rerr, err)
		}
		return fmt.Errorf("creating realtor: %w", err)
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("committing transaction: %w", err)
	}

	return nil
}

func GetRealtorRepo(ctx context.Context, entClient *ent.Client, email string) (*ent.Realtor, error) {
	r, err := entClient.Realtor.Query().Where(realtor.EmailEQ(email)).First(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, errors.New("realtor not found")
		}
		return nil, fmt.Errorf("querying realtor: %w", err)
	}

	return r, nil
}

func GetRealtorsRepo(ctx context.Context, entClient *ent.Client) ([]*ent.Realtor, error) {
	realtors, err := entClient.Realtor.Query().All(ctx)
	if err != nil {
		return nil, fmt.Errorf("querying realtors: %w", err)
	}

	return realtors, nil
}
