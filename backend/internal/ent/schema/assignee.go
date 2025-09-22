package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Assignee holds the schema definition for the Assignee entity.
type Assignee struct {
	ent.Schema
}

// Fields of the Assignee.
func (Assignee) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New),
		field.UUID("user_id", uuid.UUID{}).
			Unique(),
	}
}

// Edges of the Assignee.
func (Assignee) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("task", Task.Type).
			Ref("assignee").
			Unique(),
	}
}
