// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// RbacApplication undocumented
type RbacApplication struct {
	// Entity is the base model of RbacApplication
	Entity
	// RoleAssignments undocumented
	RoleAssignments []UnifiedRoleAssignment `json:"roleAssignments,omitempty"`
	// RoleDefinitions undocumented
	RoleDefinitions []UnifiedRoleDefinition `json:"roleDefinitions,omitempty"`
}