package repositories

import (
	"context"

	"github.com/google/uuid"
	"ppgroup.ppgroup.com/ent"
)

// ListingRepository defines the contract for listing data access.
type ListingRepository interface {
	Create(ctx context.Context, client *ent.Client, data *ent.Listing) error
	GetAll(ctx context.Context, client *ent.Client, params ListingQueryParams) ([]*ent.Listing, PaginationMeta, error)
	Delete(ctx context.Context, client *ent.Client, id string) error
	Update(ctx context.Context, client *ent.Client, data *ent.Listing) error
}

// UserRepository defines the contract for user data access.
type UserRepository interface {
	Create(ctx context.Context, client *ent.Client, data *CreateUserInput) error
	GetByEmail(ctx context.Context, client *ent.Client, email string) (*ent.User, error)
}

// RealtorRepository defines the contract for realtor data access.
type RealtorRepository interface {
	Create(ctx context.Context, client *ent.Client, data *ent.Realtor) error
	GetByEmail(ctx context.Context, client *ent.Client, email string) (*ent.Realtor, error)
	GetAll(ctx context.Context, client *ent.Client) ([]*ent.Realtor, error)
}

// Ensure package-level functions could be wrapped to satisfy these interfaces.
// This provides a clear contract for mocking in tests.
var _ ListingRepository = (*listingRepo)(nil)
var _ UserRepository = (*userRepo)(nil)
var _ RealtorRepository = (*realtorRepo)(nil)

// listingRepo wraps the package-level functions to satisfy the interface.
type listingRepo struct{}

func (r *listingRepo) Create(ctx context.Context, client *ent.Client, data *ent.Listing) error {
	return CreateListingRepo(ctx, client, data)
}
func (r *listingRepo) GetAll(ctx context.Context, client *ent.Client, params ListingQueryParams) ([]*ent.Listing, PaginationMeta, error) {
	return GetListingsRepo(ctx, client, params)
}
func (r *listingRepo) Delete(ctx context.Context, client *ent.Client, id string) error {
	return DeleteListing(ctx, client, id)
}
func (r *listingRepo) Update(ctx context.Context, client *ent.Client, data *ent.Listing) error {
	return UpdateListingRepo(ctx, client, data)
}

// userRepo wraps the package-level functions to satisfy the interface.
type userRepo struct{}

func (r *userRepo) Create(ctx context.Context, client *ent.Client, data *CreateUserInput) error {
	return CreateUserRepo(ctx, client, data)
}
func (r *userRepo) GetByEmail(ctx context.Context, client *ent.Client, email string) (*ent.User, error) {
	return GetUserRepo(ctx, client, email)
}

// realtorRepo wraps the package-level functions to satisfy the interface.
type realtorRepo struct{}

func (r *realtorRepo) Create(ctx context.Context, client *ent.Client, data *ent.Realtor) error {
	return CreateRealtorRepo(ctx, client, data)
}
func (r *realtorRepo) GetByEmail(ctx context.Context, client *ent.Client, email string) (*ent.Realtor, error) {
	return GetRealtorRepo(ctx, client, email)
}
func (r *realtorRepo) GetAll(ctx context.Context, client *ent.Client) ([]*ent.Realtor, error) {
	return GetRealtorsRepo(ctx, client)
}

// Constructors for the interface-backed implementations.
func NewListingRepository() ListingRepository   { return &listingRepo{} }
func NewUserRepository() UserRepository         { return &userRepo{} }
func NewRealtorRepository() RealtorRepository   { return &realtorRepo{} }

// compile-time check that we didn't accidentally remove the uuid import
var _ = uuid.UUID{}
