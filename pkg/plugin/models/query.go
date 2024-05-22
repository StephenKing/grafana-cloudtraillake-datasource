package models

import (
	"github.com/grafana/sqlds/v2"
)

// FormatQueryOption defines how the user has chosen to represent the data
type FormatQueryOption uint32

// Query is the model that represents the query that users submit from the panel / queryeditor.
// For the sake of backwards compatibility, when making changes to this type, ensure that changes are
// only additive.
type Query struct {
	*sqlds.Query

	// Macros
	EdsId string `json:"edsId,omitempty"`
}
