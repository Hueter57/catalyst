package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Reminder holds the schema definition for the Reminder entity.
type Reminder struct {
	ent.Schema
}

// Fields of the Reminder.
func (Reminder) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New),
		field.Time("remind_at"),
	}
}

// Edges of the Reminder.
func (Reminder) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("task", Task.Type).
			Ref("reminder").
			Unique(),
	}
}
