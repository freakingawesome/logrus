package logrus

import "context"

type ContextFielder interface {
	// SetFieldsFromContext receives a context and a field map and should pull usage-specific
	// fields from Context.Value(), adding them to the field map.
	SetFieldsFromContext(context.Context, Fields)
}
