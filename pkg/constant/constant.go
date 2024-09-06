package constant

// User Roles
const (
	RoleAdmin     = "admin"
	RoleStudent   = "student"
	RoleTeacher   = "teacher"
	RoleAdminID   = 111
	RoleStudentID = 222
	RoleTeacherID = 333
)

var AllRoles = []string{RoleAdmin, RoleStudent, RoleTeacher}

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
