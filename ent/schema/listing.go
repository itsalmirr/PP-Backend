package schema

import (
	"regexp"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

// Listing holds the schema definition for the Listing entity.
type Listing struct {
	ent.Schema
}

// Fields of the Listing.
func (Listing) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.String("title").MaxLen(120).MinLen(10).NotEmpty(),
		field.String("address").MaxLen(255).Unique().NotEmpty(),
		field.String("city").MaxLen(255).NotEmpty(),
		field.String("state").MaxLen(3).NotEmpty().Match(regexp.MustCompile(`^[A-Z]{2}$`)),
		field.String("zip_code").MaxLen(6).NotEmpty().Match(regexp.MustCompile(`^\d{5}$`)),
		field.Text("description").Optional(),
		field.Float("price").GoType(decimal.Decimal{}).SchemaType(map[string]string{dialect.Postgres: "numeric"}).Positive(),
		field.Int("bedroom").Positive(),
		field.Float("bathroom").Positive(),
		field.Int("garage").Optional().Nillable().Positive(),
		field.Int("sqft").Positive(),
		field.Enum("type_of_property").Values("house", "apartment", "condo", "townhouse").Default("house"),
		field.Enum("status").Values("DRAFT", "PUBLISHED", "ARCHIVED").Default("DRAFT"),
		field.Int("lot_size").Optional().Nillable().Positive(),
		field.Bool("pool").Optional(),
		field.Int("year_built").Positive().Range(1800, time.Now().Year()),
		field.JSON("media", map[string]interface{}{}).Optional(),
	}
}

// Edges of the Listing.
func (Listing) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("realtor", Realtor.Type).Ref("listings").Unique().Field("realtor_id"),
	}
}
