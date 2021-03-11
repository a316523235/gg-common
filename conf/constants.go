package conf

const (
	DEV          = 1
	TEST         = 2
	YF           = 3
	PROD         = 4
	PasswordHalt = "123^_^456"
	TokenHalt    = "456^_^789"

	//redis key
	AdminTokenKey = "gin_admin_token_%s"
)
