package schema

import (
	"regexp"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"entgo.io/ent/schema/mixin"
	"github.com/google/uuid"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Mixin of the User.
func (User) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}

// Fields of the User.
func (User) Fields() []ent.Field {
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)

	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.String("avatar").MaxLen(255).Optional(),
		field.String("email").
			MaxLen(120).
			Unique().
			NotEmpty().
			Match(emailRegex),
		field.String("username").
			MaxLen(120).
			Unique().
			NotEmpty().
			MinLen(3),
		field.String("full_name").MaxLen(100).NotEmpty(),
		field.Time("start_date").Default(time.Now).Immutable(),
		field.Bool("is_staff").Default(false),
		field.Bool("is_active").Default(true),
		field.String("password").MaxLen(128).NotEmpty().Sensitive(),
		field.String("provider").Default("email"),
		field.String("provider_id").Optional(),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{}
}

// Indexes of the User.
func (User) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("email").Unique(),
		index.Fields("username").Unique(),
		index.Fields("provider", "provider_id").Unique(),
	}
}
