package security

type RolePermissionMap map[Role]map[Permission]struct{}

var RolePermissions RolePermissionMap = RolePermissionMap{
	Administrator: {
		CREATE_SESSION:       {},
		VIEW_SESSION_OWN:     {},
		VIEW_SESSION_OTHER:   {},
		DELETE_SESSION_OWN:   {},
		DELETE_SESSION_OTHER: {},

		CREATE_USER_OTHER: {},
		VIEW_USER_OWN:     {},
		VIEW_USER_OTHER:   {},
		EDIT_USER_OWN:     {},
		EDIT_USER_OTHER:   {},
		DELETE_USER_OWN:   {},
		DELETE_USER_OTHER: {},
	},
}
