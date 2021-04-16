// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "time"

// Schedule undocumented
type Schedule struct {
	// Entity is the base model of Schedule
	Entity
	// Enabled undocumented
	Enabled *bool `json:"enabled,omitempty"`
	// OfferShiftRequestsEnabled undocumented
	OfferShiftRequestsEnabled *bool `json:"offerShiftRequestsEnabled,omitempty"`
	// OpenShiftsEnabled undocumented
	OpenShiftsEnabled *bool `json:"openShiftsEnabled,omitempty"`
	// ProvisionStatus undocumented
	ProvisionStatus *OperationStatus `json:"provisionStatus,omitempty"`
	// ProvisionStatusCode undocumented
	ProvisionStatusCode *string `json:"provisionStatusCode,omitempty"`
	// SwapShiftsRequestsEnabled undocumented
	SwapShiftsRequestsEnabled *bool `json:"swapShiftsRequestsEnabled,omitempty"`
	// TimeClockEnabled undocumented
	TimeClockEnabled *bool `json:"timeClockEnabled,omitempty"`
	// TimeOffRequestsEnabled undocumented
	TimeOffRequestsEnabled *bool `json:"timeOffRequestsEnabled,omitempty"`
	// TimeZone undocumented
	TimeZone *string `json:"timeZone,omitempty"`
	// WorkforceIntegrationIDs undocumented
	WorkforceIntegrationIDs []string `json:"workforceIntegrationIds,omitempty"`
	// OfferShiftRequests undocumented
	OfferShiftRequests []OfferShiftRequestObject `json:"offerShiftRequests,omitempty"`
	// OpenShiftChangeRequests undocumented
	OpenShiftChangeRequests []OpenShiftChangeRequestObject `json:"openShiftChangeRequests,omitempty"`
	// OpenShifts undocumented
	OpenShifts []OpenShift `json:"openShifts,omitempty"`
	// SchedulingGroups undocumented
	SchedulingGroups []SchedulingGroup `json:"schedulingGroups,omitempty"`
	// Shifts undocumented
	Shifts []Shift `json:"shifts,omitempty"`
	// SwapShiftsChangeRequests undocumented
	SwapShiftsChangeRequests []SwapShiftsChangeRequestObject `json:"swapShiftsChangeRequests,omitempty"`
	// TimeOffReasons undocumented
	TimeOffReasons []TimeOffReason `json:"timeOffReasons,omitempty"`
	// TimeOffRequests undocumented
	TimeOffRequests []TimeOffRequestObject `json:"timeOffRequests,omitempty"`
	// TimesOff undocumented
	TimesOff []TimeOff `json:"timesOff,omitempty"`
}

// ScheduleChangeRequestObject undocumented
type ScheduleChangeRequestObject struct {
	// ChangeTrackedEntity is the base model of ScheduleChangeRequestObject
	ChangeTrackedEntity
	// AssignedTo undocumented
	AssignedTo *ScheduleChangeRequestActor `json:"assignedTo,omitempty"`
	// ManagerActionDateTime undocumented
	ManagerActionDateTime *time.Time `json:"managerActionDateTime,omitempty"`
	// ManagerActionMessage undocumented
	ManagerActionMessage *string `json:"managerActionMessage,omitempty"`
	// ManagerUserID undocumented
	ManagerUserID *string `json:"managerUserId,omitempty"`
	// SenderDateTime undocumented
	SenderDateTime *time.Time `json:"senderDateTime,omitempty"`
	// SenderMessage undocumented
	SenderMessage *string `json:"senderMessage,omitempty"`
	// SenderUserID undocumented
	SenderUserID *string `json:"senderUserId,omitempty"`
	// State undocumented
	State *ScheduleChangeState `json:"state,omitempty"`
}

// ScheduleEntity undocumented
type ScheduleEntity struct {
	// Object is the base model of ScheduleEntity
	Object
	// EndDateTime undocumented
	EndDateTime *time.Time `json:"endDateTime,omitempty"`
	// StartDateTime undocumented
	StartDateTime *time.Time `json:"startDateTime,omitempty"`
	// Theme undocumented
	Theme *ScheduleEntityTheme `json:"theme,omitempty"`
}

// ScheduleInformation undocumented
type ScheduleInformation struct {
	// Object is the base model of ScheduleInformation
	Object
	// AvailabilityView undocumented
	AvailabilityView *string `json:"availabilityView,omitempty"`
	// Error undocumented
	Error *FreeBusyError `json:"error,omitempty"`
	// ScheduleID undocumented
	ScheduleID *string `json:"scheduleId,omitempty"`
	// ScheduleItems undocumented
	ScheduleItems []ScheduleItem `json:"scheduleItems,omitempty"`
	// WorkingHours undocumented
	WorkingHours *WorkingHours `json:"workingHours,omitempty"`
}

// ScheduleItem undocumented
type ScheduleItem struct {
	// Object is the base model of ScheduleItem
	Object
	// End undocumented
	End *DateTimeTimeZone `json:"end,omitempty"`
	// IsPrivate undocumented
	IsPrivate *bool `json:"isPrivate,omitempty"`
	// Location undocumented
	Location *string `json:"location,omitempty"`
	// Start undocumented
	Start *DateTimeTimeZone `json:"start,omitempty"`
	// Status undocumented
	Status *FreeBusyStatus `json:"status,omitempty"`
	// Subject undocumented
	Subject *string `json:"subject,omitempty"`
}
