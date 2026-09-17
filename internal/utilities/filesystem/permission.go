package filesystem

type FilePermission string

const (
	PermissionRead    FilePermission = "read"
	PermissionWrite   FilePermission = "write"
	PermissionExecute FilePermission = "execute"

	PermissionSetuid FilePermission = "setuid"
	PermissionSetgid FilePermission = "setgid"
	PermissionSticky FilePermission = "sticky"
)
