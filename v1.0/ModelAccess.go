// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "time"

// AccessAction undocumented
type AccessAction struct {
	// Object is the base model of AccessAction
	Object
}

// AccessReviewApplyAction undocumented
type AccessReviewApplyAction struct {
	// Object is the base model of AccessReviewApplyAction
	Object
}

// AccessReviewInactiveUsersQueryScope undocumented
type AccessReviewInactiveUsersQueryScope struct {
	// AccessReviewQueryScope is the base model of AccessReviewInactiveUsersQueryScope
	AccessReviewQueryScope
	// InactiveDuration undocumented
	InactiveDuration *Duration `json:"inactiveDuration,omitempty"`
}

// AccessReviewInstance undocumented
type AccessReviewInstance struct {
	// Entity is the base model of AccessReviewInstance
	Entity
	// EndDateTime undocumented
	EndDateTime *time.Time `json:"endDateTime,omitempty"`
	// Scope undocumented
	Scope *AccessReviewScope `json:"scope,omitempty"`
	// StartDateTime undocumented
	StartDateTime *time.Time `json:"startDateTime,omitempty"`
	// Status undocumented
	Status *string `json:"status,omitempty"`
	// Decisions undocumented
	Decisions []AccessReviewInstanceDecisionItem `json:"decisions,omitempty"`
}

// AccessReviewInstanceDecisionItem undocumented
type AccessReviewInstanceDecisionItem struct {
	// Entity is the base model of AccessReviewInstanceDecisionItem
	Entity
	// AccessReviewID undocumented
	AccessReviewID *string `json:"accessReviewId,omitempty"`
	// AppliedBy undocumented
	AppliedBy *UserIdentity `json:"appliedBy,omitempty"`
	// AppliedDateTime undocumented
	AppliedDateTime *time.Time `json:"appliedDateTime,omitempty"`
	// ApplyResult undocumented
	ApplyResult *string `json:"applyResult,omitempty"`
	// Decision undocumented
	Decision *string `json:"decision,omitempty"`
	// Justification undocumented
	Justification *string `json:"justification,omitempty"`
	// Principal undocumented
	Principal *Identity `json:"principal,omitempty"`
	// PrincipalLink undocumented
	PrincipalLink *string `json:"principalLink,omitempty"`
	// Recommendation undocumented
	Recommendation *string `json:"recommendation,omitempty"`
	// Resource undocumented
	Resource *AccessReviewInstanceDecisionItemResource `json:"resource,omitempty"`
	// ResourceLink undocumented
	ResourceLink *string `json:"resourceLink,omitempty"`
	// ReviewedBy undocumented
	ReviewedBy *UserIdentity `json:"reviewedBy,omitempty"`
	// ReviewedDateTime undocumented
	ReviewedDateTime *time.Time `json:"reviewedDateTime,omitempty"`
}

// AccessReviewInstanceDecisionItemResource undocumented
type AccessReviewInstanceDecisionItemResource struct {
	// Object is the base model of AccessReviewInstanceDecisionItemResource
	Object
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
	// ID undocumented
	ID *string `json:"id,omitempty"`
	// Type undocumented
	Type *string `json:"type,omitempty"`
}

// AccessReviewQueryScope undocumented
type AccessReviewQueryScope struct {
	// AccessReviewScope is the base model of AccessReviewQueryScope
	AccessReviewScope
	// Query undocumented
	Query *string `json:"query,omitempty"`
	// QueryRoot undocumented
	QueryRoot *string `json:"queryRoot,omitempty"`
	// QueryType undocumented
	QueryType *string `json:"queryType,omitempty"`
}

// AccessReviewReviewerScope undocumented
type AccessReviewReviewerScope struct {
	// Object is the base model of AccessReviewReviewerScope
	Object
	// Query undocumented
	Query *string `json:"query,omitempty"`
	// QueryRoot undocumented
	QueryRoot *string `json:"queryRoot,omitempty"`
	// QueryType undocumented
	QueryType *string `json:"queryType,omitempty"`
}

// AccessReviewScheduleDefinition undocumented
type AccessReviewScheduleDefinition struct {
	// Entity is the base model of AccessReviewScheduleDefinition
	Entity
	// CreatedBy undocumented
	CreatedBy *UserIdentity `json:"createdBy,omitempty"`
	// CreatedDateTime undocumented
	CreatedDateTime *time.Time `json:"createdDateTime,omitempty"`
	// DescriptionForAdmins undocumented
	DescriptionForAdmins *string `json:"descriptionForAdmins,omitempty"`
	// DescriptionForReviewers undocumented
	DescriptionForReviewers *string `json:"descriptionForReviewers,omitempty"`
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
	// FallbackReviewers undocumented
	FallbackReviewers []AccessReviewReviewerScope `json:"fallbackReviewers,omitempty"`
	// InstanceEnumerationScope undocumented
	InstanceEnumerationScope *AccessReviewScope `json:"instanceEnumerationScope,omitempty"`
	// LastModifiedDateTime undocumented
	LastModifiedDateTime *time.Time `json:"lastModifiedDateTime,omitempty"`
	// Reviewers undocumented
	Reviewers []AccessReviewReviewerScope `json:"reviewers,omitempty"`
	// Scope undocumented
	Scope *AccessReviewScope `json:"scope,omitempty"`
	// Settings undocumented
	Settings *AccessReviewScheduleSettings `json:"settings,omitempty"`
	// Status undocumented
	Status *string `json:"status,omitempty"`
	// Instances undocumented
	Instances []AccessReviewInstance `json:"instances,omitempty"`
}

// AccessReviewScheduleSettings undocumented
type AccessReviewScheduleSettings struct {
	// Object is the base model of AccessReviewScheduleSettings
	Object
	// ApplyActions undocumented
	ApplyActions []AccessReviewApplyAction `json:"applyActions,omitempty"`
	// AutoApplyDecisionsEnabled undocumented
	AutoApplyDecisionsEnabled *bool `json:"autoApplyDecisionsEnabled,omitempty"`
	// DefaultDecision undocumented
	DefaultDecision *string `json:"defaultDecision,omitempty"`
	// DefaultDecisionEnabled undocumented
	DefaultDecisionEnabled *bool `json:"defaultDecisionEnabled,omitempty"`
	// InstanceDurationInDays undocumented
	InstanceDurationInDays *int `json:"instanceDurationInDays,omitempty"`
	// JustificationRequiredOnApproval undocumented
	JustificationRequiredOnApproval *bool `json:"justificationRequiredOnApproval,omitempty"`
	// MailNotificationsEnabled undocumented
	MailNotificationsEnabled *bool `json:"mailNotificationsEnabled,omitempty"`
	// RecommendationsEnabled undocumented
	RecommendationsEnabled *bool `json:"recommendationsEnabled,omitempty"`
	// Recurrence undocumented
	Recurrence *PatternedRecurrence `json:"recurrence,omitempty"`
	// ReminderNotificationsEnabled undocumented
	ReminderNotificationsEnabled *bool `json:"reminderNotificationsEnabled,omitempty"`
}

// AccessReviewScope undocumented
type AccessReviewScope struct {
	// Object is the base model of AccessReviewScope
	Object
}

// AccessReviewSet undocumented
type AccessReviewSet struct {
	// Entity is the base model of AccessReviewSet
	Entity
	// Definitions undocumented
	Definitions []AccessReviewScheduleDefinition `json:"definitions,omitempty"`
}
