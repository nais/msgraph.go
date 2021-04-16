// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// Teamwork undocumented
type Teamwork struct {
	// Entity is the base model of Teamwork
	Entity
	// WorkforceIntegrations undocumented
	WorkforceIntegrations []WorkforceIntegration `json:"workforceIntegrations,omitempty"`
}

// TeamworkActivityTopic undocumented
type TeamworkActivityTopic struct {
	// Object is the base model of TeamworkActivityTopic
	Object
	// Source undocumented
	Source *TeamworkActivityTopicSource `json:"source,omitempty"`
	// Value undocumented
	Value *string `json:"value,omitempty"`
	// WebURL undocumented
	WebURL *string `json:"webUrl,omitempty"`
}

// TeamworkBot undocumented
type TeamworkBot struct {
	// Entity is the base model of TeamworkBot
	Entity
}

// TeamworkHostedContent undocumented
type TeamworkHostedContent struct {
	// Entity is the base model of TeamworkHostedContent
	Entity
	// ContentBytes undocumented
	ContentBytes *Binary `json:"contentBytes,omitempty"`
	// ContentType undocumented
	ContentType *string `json:"contentType,omitempty"`
}

// TeamworkNotificationRecipient undocumented
type TeamworkNotificationRecipient struct {
	// Object is the base model of TeamworkNotificationRecipient
	Object
}