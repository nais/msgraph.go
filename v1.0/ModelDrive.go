// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// Drive undocumented
type Drive struct {
	// BaseItem is the base model of Drive
	BaseItem
	// DriveType undocumented
	DriveType *string `json:"driveType,omitempty"`
	// Owner undocumented
	Owner *IdentitySet `json:"owner,omitempty"`
	// Quota undocumented
	Quota *Quota `json:"quota,omitempty"`
	// SharePointIDs undocumented
	SharePointIDs *SharepointIDs `json:"sharePointIds,omitempty"`
	// System undocumented
	System *SystemFacet `json:"system,omitempty"`
	// Following undocumented
	Following []DriveItem `json:"following,omitempty"`
	// Items undocumented
	Items []DriveItem `json:"items,omitempty"`
	// List undocumented
	List *List `json:"list,omitempty"`
	// Root undocumented
	Root *DriveItem `json:"root,omitempty"`
	// Special undocumented
	Special []DriveItem `json:"special,omitempty"`
}

// DriveItem undocumented
type DriveItem struct {
	// BaseItem is the base model of DriveItem
	BaseItem
	// Audio undocumented
	Audio *Audio `json:"audio,omitempty"`
	// Content undocumented
	Content *Stream `json:"content,omitempty"`
	// CTag undocumented
	CTag *string `json:"cTag,omitempty"`
	// Deleted undocumented
	Deleted *Deleted `json:"deleted,omitempty"`
	// File undocumented
	File *File `json:"file,omitempty"`
	// FileSystemInfo undocumented
	FileSystemInfo *FileSystemInfo `json:"fileSystemInfo,omitempty"`
	// Folder undocumented
	Folder *Folder `json:"folder,omitempty"`
	// Image undocumented
	Image *Image `json:"image,omitempty"`
	// Location undocumented
	Location *GeoCoordinates `json:"location,omitempty"`
	// Package undocumented
	Package *Package `json:"package,omitempty"`
	// PendingOperations undocumented
	PendingOperations *PendingOperations `json:"pendingOperations,omitempty"`
	// Photo undocumented
	Photo *Photo `json:"photo,omitempty"`
	// Publication undocumented
	Publication *PublicationFacet `json:"publication,omitempty"`
	// RemoteItem undocumented
	RemoteItem *RemoteItem `json:"remoteItem,omitempty"`
	// Root undocumented
	Root *Root `json:"root,omitempty"`
	// SearchResult undocumented
	SearchResult *SearchResult `json:"searchResult,omitempty"`
	// Shared undocumented
	Shared *Shared `json:"shared,omitempty"`
	// SharepointIDs undocumented
	SharepointIDs *SharepointIDs `json:"sharepointIds,omitempty"`
	// Size undocumented
	Size *int `json:"size,omitempty"`
	// SpecialFolder undocumented
	SpecialFolder *SpecialFolder `json:"specialFolder,omitempty"`
	// Video undocumented
	Video *Video `json:"video,omitempty"`
	// WebDavURL undocumented
	WebDavURL *string `json:"webDavUrl,omitempty"`
	// Workbook undocumented
	Workbook *Workbook `json:"workbook,omitempty"`
	// Analytics undocumented
	Analytics *ItemAnalytics `json:"analytics,omitempty"`
	// Children undocumented
	Children []DriveItem `json:"children,omitempty"`
	// ListItem undocumented
	ListItem *ListItem `json:"listItem,omitempty"`
	// Permissions undocumented
	Permissions []Permission `json:"permissions,omitempty"`
	// Subscriptions undocumented
	Subscriptions []Subscription `json:"subscriptions,omitempty"`
	// Thumbnails undocumented
	Thumbnails []ThumbnailSet `json:"thumbnails,omitempty"`
	// Versions undocumented
	Versions []DriveItemVersion `json:"versions,omitempty"`
}

// DriveItemUploadableProperties undocumented
type DriveItemUploadableProperties struct {
	// Object is the base model of DriveItemUploadableProperties
	Object
	// Description undocumented
	Description *string `json:"description,omitempty"`
	// FileSize undocumented
	FileSize *int `json:"fileSize,omitempty"`
	// FileSystemInfo undocumented
	FileSystemInfo *FileSystemInfo `json:"fileSystemInfo,omitempty"`
	// Name undocumented
	Name *string `json:"name,omitempty"`
}

// DriveItemVersion undocumented
type DriveItemVersion struct {
	// BaseItemVersion is the base model of DriveItemVersion
	BaseItemVersion
	// Content undocumented
	Content *Stream `json:"content,omitempty"`
	// Size undocumented
	Size *int `json:"size,omitempty"`
}

// DriveRecipient undocumented
type DriveRecipient struct {
	// Object is the base model of DriveRecipient
	Object
	// Alias undocumented
	Alias *string `json:"alias,omitempty"`
	// Email undocumented
	Email *string `json:"email,omitempty"`
	// ObjectID undocumented
	ObjectID *string `json:"objectId,omitempty"`
}
