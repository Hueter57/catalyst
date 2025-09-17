package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Task holds the schema definition for the Task entity.
type Task struct {
	ent.Schema
}

// Fields of the Task.
func (Task) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New),
		field.String("title").
			NotEmpty(),
		field.Text("description").
			Optional(),
		field.String("status").
			Default("in_progress"),
		field.String("importance").
			Default("low"),
		field.Time("due_date"),
		field.String("message_id").
			Optional(),
		field.String("channel_id").
			NotEmpty(),
		field.Time("created_at").
			Default(time.Now),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now),
	}
}

// Edges of the Task.
func (Task) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("category", Category.Type),
		edge.To("assignee", Assignee.Type).
			Unique(),
		edge.To("reminder", Reminder.Type).
			Unique(),
	}
}
