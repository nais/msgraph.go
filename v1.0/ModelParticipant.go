// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// Participant undocumented
type Participant struct {
	// Entity is the base model of Participant
	Entity
	// Info undocumented
	Info *ParticipantInfo `json:"info,omitempty"`
	// IsInLobby undocumented
	IsInLobby *bool `json:"isInLobby,omitempty"`
	// IsMuted undocumented
	IsMuted *bool `json:"isMuted,omitempty"`
	// MediaStreams undocumented
	MediaStreams []MediaStream `json:"mediaStreams,omitempty"`
	// RecordingInfo undocumented
	RecordingInfo *RecordingInfo `json:"recordingInfo,omitempty"`
}

// ParticipantInfo undocumented
type ParticipantInfo struct {
	// Object is the base model of ParticipantInfo
	Object
	// CountryCode undocumented
	CountryCode *string `json:"countryCode,omitempty"`
	// EndpointType undocumented
	EndpointType *EndpointType `json:"endpointType,omitempty"`
	// Identity undocumented
	Identity *IdentitySet `json:"identity,omitempty"`
	// LanguageID undocumented
	LanguageID *string `json:"languageId,omitempty"`
	// Region undocumented
	Region *string `json:"region,omitempty"`
}
