// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// TargetedManagedAppPolicyAssignment The type for deployment of groups or apps.
type TargetedManagedAppPolicyAssignment struct {
	Entity
	// Target Identifier for deployment of a group or app
	Target *DeviceAndAppManagementAssignmentTarget `json:"target,omitempty"`
}