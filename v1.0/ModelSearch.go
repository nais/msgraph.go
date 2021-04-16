// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// SearchEntity undocumented
type SearchEntity struct {
	// Entity is the base model of SearchEntity
	Entity
}

// SearchHit undocumented
type SearchHit struct {
	// Object is the base model of SearchHit
	Object
	// ContentSource undocumented
	ContentSource *string `json:"contentSource,omitempty"`
	// HitID undocumented
	HitID *string `json:"hitId,omitempty"`
	// Rank undocumented
	Rank *int `json:"rank,omitempty"`
	// Summary undocumented
	Summary *string `json:"summary,omitempty"`
	// Resource undocumented
	Resource *Entity `json:"resource,omitempty"`
}

// SearchHitsContainer undocumented
type SearchHitsContainer struct {
	// Object is the base model of SearchHitsContainer
	Object
	// Hits undocumented
	Hits []SearchHit `json:"hits,omitempty"`
	// MoreResultsAvailable undocumented
	MoreResultsAvailable *bool `json:"moreResultsAvailable,omitempty"`
	// Total undocumented
	Total *int `json:"total,omitempty"`
}

// SearchQuery undocumented
type SearchQuery struct {
	// Object is the base model of SearchQuery
	Object
	// QueryString undocumented
	QueryString *string `json:"queryString,omitempty"`
}

// SearchRequestObject undocumented
type SearchRequestObject struct {
	// Object is the base model of SearchRequestObject
	Object
	// ContentSources undocumented
	ContentSources []string `json:"contentSources,omitempty"`
	// EnableTopResults undocumented
	EnableTopResults *bool `json:"enableTopResults,omitempty"`
	// EntityTypes undocumented
	EntityTypes []EntityType `json:"entityTypes,omitempty"`
	// Fields undocumented
	Fields []string `json:"fields,omitempty"`
	// From undocumented
	From *int `json:"from,omitempty"`
	// Query undocumented
	Query *SearchQuery `json:"query,omitempty"`
	// Size undocumented
	Size *int `json:"size,omitempty"`
}

// SearchResponse undocumented
type SearchResponse struct {
	// Object is the base model of SearchResponse
	Object
	// HitsContainers undocumented
	HitsContainers []SearchHitsContainer `json:"hitsContainers,omitempty"`
	// SearchTerms undocumented
	SearchTerms []string `json:"searchTerms,omitempty"`
}

// SearchResult undocumented
type SearchResult struct {
	// Object is the base model of SearchResult
	Object
	// OnClickTelemetryURL undocumented
	OnClickTelemetryURL *string `json:"onClickTelemetryUrl,omitempty"`
}
