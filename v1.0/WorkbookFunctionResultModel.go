// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "encoding/json"

// WorkbookFunctionResult undocumented
type WorkbookFunctionResult struct {
	Entity
	// Error undocumented
	Error *string `json:"error,omitempty"`
	// Value undocumented
	Value json.RawMessage `json:"value,omitempty"`
}