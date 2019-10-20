// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// DeviceManagementSettingDependency undocumented
type DeviceManagementSettingDependency struct {
	// DefinitionID The setting definition ID of the setting depended on
	DefinitionID *string `json:"definitionId,omitempty"`
	// Constraints Collection of constraints for the dependency setting value
	Constraints []DeviceManagementConstraint `json:"constraints,omitempty"`
}