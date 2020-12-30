package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
)

type Installment struct {
	ent.schema
}

func (Installment) Fields() []ent.Field {
	return []ent.Field {
		field.String("Paymoneyperoid").NotEmty(),
		field.Int("Down").Positive(),
		field.Int("Paymoneypermount").Positive(),
	}
}

func (Playlist_Video) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("customer", Customer.Type).Ref("installment").Unique(),
		edge.From("saler", Saler.Type).Ref("installment").Unique(),
	}
}

