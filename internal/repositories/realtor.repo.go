package repositories

import (
	"context"
	"errors"

	"ppgroup.m0chi.com/ent"
	"ppgroup.m0chi.com/ent/realtor"
)

// CreateRealtorRepository creates a new realtor record in the database.
// It first checks if a realtor with the given email or phone already exists.
// If such a realtor exists, it returns an error.
// If not, it creates a new realtor record with the provided data.
// The operation is performed within a transaction to ensure atomicity.
//
// Parameters:
//   - data: CreateRealtorInput containing the details of the realtor to be created.
//
// Returns:
//   - error: An error if the realtor already exists or if the creation fails, otherwise nil.
func CreateRealtorRepo(entClient *ent.Client, data *ent.Realtor) error {
	ctx := context.Background()

	exists, err := entClient.Realtor.Query().Where(realtor.Or(realtor.EmailEQ(data.Email), realtor.PhoneEQ(data.Phone))).Exist(ctx)
	if err != nil {
		return err
	}

	if exists {
		return errors.New("realtor with the given email or phone already exists")
	}

	// Start a transaction
	tx, err := entClient.Tx(ctx)
	if err != nil {
		return err
	}

	// Create a new realtor
	_, err = tx.Realtor.Create().SetEmail(data.Email).SetFullName(data.FullName).SetPhone(data.Phone).SetIsMvp(data.IsMvp).Save(context.Background())
	if err != nil {
		tx.Rollback()
		return errors.New("failed to create realtor" + err.Error())
	}

	// Commit the transaction
	if err := tx.Commit(); err != nil {
		return errors.New("failed to commit transaction")
	}

	return nil
}

// GetRealtorRepository retrieves a realtor record from the database based on the provided email.
// It returns a models.Realtor object and an error if any occurred during the query.
// If the realtor is not found, it returns an error indicating "user not found".
// If there is any other error during the query, it returns an error indicating "failed to get user".
func GetRealtorRepo(entClient *ent.Client, email string) (*ent.Realtor, error) {
	ctx := context.Background()

	realtor, err := entClient.Realtor.Query().Where(realtor.EmailEQ(email)).First(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, errors.New("realtor not found")
		}
		return nil, errors.New("failed to get realtor")
	}

	return realtor, nil
}

func GetRealtorsRepo(entClient *ent.Client) ([]*ent.Realtor, error) {
	ctx := context.Background()

	// Get all realtors from database
	realtors, err := entClient.Realtor.Query().All(ctx)

	if err != nil {
		return nil, err
	}

	return realtors, nil
}
