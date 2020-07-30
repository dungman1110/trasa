package consts

const (
	CREATE_USER          = "CREATE_USER"
	CREATE_ADMIN_USER    = "CREATE_ADMIN_USER"
	UPDATE_USER          = "UPDATE_USER"
	DELETE_USER          = "DELETE_USER"
	DELETE_MULTIPLE_USER = "DELETE_MULTIPLE_USER"

	// Admins
	ADMIN_ACTIVITY         = "ADMIN_ACTIVITY"
	DELETE_ADMIN_USER      = "DELETE_ADMIN_USER"
	GRANT_ADMIN_PRIVILEGE  = "GRANT_ADMIN_PRIVILEGE"
	ADMIN_PROFILE_EDITED   = "ADMIN_PROFILE_EDITED"
	REVOKE_ADMIN_PRIVILEGE = "REVOKE_ADMIN_PRIVILEGE"

	CREATE_APP = "CREATE_APP"
	UPDATE_APP = "UPDATE_APP"
	DELETE_APP = "DELETE_APP"

	CREATE_GROUP = "CREATE_GROUP"
	UPDATE_GROUP = "UPDATE_GROUP"
	DELETE_GROUP = "DELETE_GROUP"

	CREATE_POLICY = "CREATE_POLICY"
	UPDATE_POLICY = "UPDATE_POLICY"
	DELETE_POLICY = "DELETE_POLICY"

	ASSIGN_USER_TO_APP             = "ASSIGN_USER_TO_APP"
	ASSIGN_USER_TO_APPGROUP        = "ASSIGN_USER_TO_APPGROUP"
	ASSIGN_USERGROUP_TO_APP        = "ASSIGN_USERGROUP_TO_APP"
	ASSIGN_USERGROUP_TO_APPGROUP   = "ASSIGN_USERGROUP_TO_APPGROUP"
	REMOVE_USER_FROM_APP           = "REMOVE_USER_FROM_APP"
	REMOVE_USERGROUP_FROM_APP      = "REMOVE_USER_FROM_APP"
	REMOVE_USERGROUP_FROM_APPGROUP = "REMOVE_USER_FROM_APP"

	SET_GLOBAL_SETTING = "SET_GLOBAL_SETTING"

	ENROLL_USER_DEVICE = "ENROLL_USER_DEVICE"
	REMOVE_USER_DEVICE = "REMOVE_USER_DEVICE"

	CREATE_HTTP_PROXY = "CREATE_HTTP_PROXY"
	UPDATE_HTTP_PROXY = "UPDATE_HTTP_PROXY"
	DELETE_HTTP_PROXY = "DELETE_HTTP_PROXY"

	UPDATE_PASSWORD_POLICY = "UPDATE_PASSWORD_POLICY"
	UPDATE_SSH_CERT_POLICY = "UPDATE_SSH_CERT_POLICY"

	VAULT_INITIALIZED     = "VAULT_INITIALIZED"
	VAULT_RESET           = "VAULT_RESET"
	VAULT_UNSEALED        = "VAULT_UNSEALED"
	VIEW_VAULT_PASSWORD   = "VIEW_VAULT_PASSWORD"
	SAVE_VAULT_PASSWORD   = "SAVE_VAULT_PASSWORD"
	REMOVE_VAULT_PASSWORD = "REMOVE_VAULT_PASSWORD"
	VAULT_DECRYPTED       = "VAULT_DECRYPTED"

	TAKE_BACKUP        = "TAKE_BACKUP"
	CREATE_BACKUP_PLAN = "CREATE_BACKUP_PLAN"

	FILE_UPLOAD   = "FILE_UPLOAD"
	FILE_DOWNLOAD = "FILE_DOWNLOAD"

	UPDATE_APP_CERTS    = "UPDATE_APP_CERTS"
	DOWNLOAD_HOST_CERTS = "DOWNLOAD_HOST_CERTS"
	GENERATE_USER_CERTS = "GENERATE_USER_CERTS"

	CREATE_IDP            = "CREATE_IDP"
	UPDATE_IDP            = "UPDATE_IDP"
	IMPORT_USERS_FROM_IDP = "IMPORT_USERS_FROM_IDP"
	GENERATE_SCIM_TOKEN   = "GENERATE_SCIM_TOKEN"

	ADHOC_REQUEST_ACK = "ADHOC_REQUEST_ACK"
)
