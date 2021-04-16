// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// Extension undocumented
type Extension struct {
	// Entity is the base model of Extension
	Entity
}

// ExtensionProperty undocumented
type ExtensionProperty struct {
	// DirectoryObject is the base model of ExtensionProperty
	DirectoryObject
	// AppDisplayName undocumented
	AppDisplayName *string `json:"appDisplayName,omitempty"`
	// DataType undocumented
	DataType *string `json:"dataType,omitempty"`
	// IsSyncedFromOnPremises undocumented
	IsSyncedFromOnPremises *bool `json:"isSyncedFromOnPremises,omitempty"`
	// Name undocumented
	Name *string `json:"name,omitempty"`
	// TargetObjects undocumented
	TargetObjects []string `json:"targetObjects,omitempty"`
}

// ExtensionSchemaProperty undocumented
type ExtensionSchemaProperty struct {
	// Object is the base model of ExtensionSchemaProperty
	Object
	// Name undocumented
	Name *string `json:"name,omitempty"`
	// Type undocumented
	Type *string `json:"type,omitempty"`
}
