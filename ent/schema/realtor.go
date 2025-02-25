package schema

import (
	"regexp"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"entgo.io/ent/schema/mixin"
	"github.com/google/uuid"
)

// Realtor holds the schema definition for the Realtor entity.
type Realtor struct {
	ent.Schema
}

func (Realtor) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{}, // CreatedAt/UpdatedAt
	}
}

// Fields of the Realtor.
func (Realtor) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New).StorageKey("id"),

		field.String("full_name").MaxLen(100).NotEmpty().StructTag(`json:"full_name" validate:"required,min=5,max=100"`),

		field.JSON("photo", map[string]interface{}{}).Optional().StructTag(`json:"photo,omitempty" validate:"omitempty, json"`),

		field.Text("description").Optional().MaxLen(500).StructTag(`json:"description,omitempty" validate:"max=500"`),

		field.String("phone").MaxLen(20).NotEmpty().StructTag(`json:"phone"`).Match(regexp.MustCompile(`^\+[1-9]\d{1,14}$`)),

		field.String("email").MaxLen(255).NotEmpty().Unique().StructTag(`json:"email" validate:"required,email"`),

		field.Bool("is_mvp").Default(false).StructTag(`json:"is_mvp"`),

		field.Time("hire_date").Immutable().Default(time.Now).StructTag(`json:"hire_date"`),
	}
}

// Edges of the Realtor.
func (Realtor) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("listings", Listing.Type),
	}
}

// Indexes of the Realtor.
func (Realtor) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("full_name"),
		index.Fields("email").Unique(),
		index.Fields("phone").Unique(),
	}
}
