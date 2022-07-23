package security

type Permission int32

const (
	// USER PERMISSIONS
	CREATE_USER_OWN Permission = iota + 1
	CREATE_USER_OTHER

	VIEW_USER_OWN
	VIEW_USER_OTHER

	EDIT_USER_OWN
	EDIT_USER_OTHER

	DELETE_USER_OWN
	DELETE_USER_OTHER

	// SESSION PERMISSIONS
	CREATE_SESSION

	VIEW_SESSION_OWN
	VIEW_SESSION_OTHER

	DELETE_SESSION_OWN
	DELETE_SESSION_OTHER
)

type PermissionInfo struct {
	Label       string
	Description string
}

var PermissionDetails = map[Permission]PermissionInfo{
	// USER PERMISSION DETAILS
	CREATE_USER_OWN: {
		Label:       "Sign-Up",
		Description: "Register a new personal user account",
	},
	CREATE_USER_OTHER: {
		Label:       "Create User",
		Description: "Create a user account for someone else with a specific role",
	},

	VIEW_USER_OWN: {
		Label:       "View Own Account",
		Description: "View user's own account",
	},
	VIEW_USER_OTHER: {
		Label:       "View Other's Accounts",
		Description: "View other user accounts",
	},

	EDIT_USER_OWN: {
		Label:       "Edit Own Account",
		Description: "Edit information of user's own account",
	},
	EDIT_USER_OTHER: {
		Label:       "Edit Other's Accounts",
		Description: "Edit information of other user accounts",
	},

	DELETE_USER_OWN: {
		Label:       "Delete Own Account",
		Description: "Delete user's own account",
	},
	DELETE_USER_OTHER: {
		Label:       "Delete Other's Accounts",
		Description: "Delete other user accounts",
	},

	// SESSION PERMISSION DETAILS
	CREATE_SESSION: {
		Label:       "Login",
		Description: "Login and create a new session",
	},

	VIEW_SESSION_OWN: {
		Label:       "View Own Sessions",
		Description: "View sessions of users own account",
	},
	VIEW_SESSION_OTHER: {
		Label:       "View Other's Sessions",
		Description: "View sessions of other user accounts",
	},

	DELETE_SESSION_OWN: {
		Label:       "Logout",
		Description: "Logout from a session in your own account",
	},
	DELETE_SESSION_OTHER: {
		Label:       "Delete Other Sessions",
		Description: "Delete sessions of other user accounts",
	},
}
