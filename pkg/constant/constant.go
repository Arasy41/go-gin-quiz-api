package constant

// User Roles
const (
	RoleAdmin       = "admin"
	RoleUser        = "user"
	RoleModerator   = "moderator"
	RoleUserID      = 12010
	RoleAdminID     = 12020
	RoleModeratorID = 12030
)

var AllRoles = []string{RoleAdmin, RoleUser, RoleModerator}

// Content-Type and Header Keys
const (
	ContentTypeJSON  = "application/json"
	AuthorizationKey = "Authorization"
)

// Application Configuration
const (
	DefaultPageSize = 20
	MaxPageSize     = 100
)

// Validation Constants
const (
	MinPasswordLength = 8
	MaxPasswordLength = 32
)
