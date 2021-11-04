package commonschema

type IdentityType string

const (
	none                       IdentityType = "None"
	systemAssigned             IdentityType = "SystemAssigned"
	userAssigned               IdentityType = "UserAssigned"
	systemAssignedUserAssigned IdentityType = "SystemAssigned, UserAssigned"
)
