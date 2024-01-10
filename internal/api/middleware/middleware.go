package middleware

var (
	Resource      = resource()
	Auth          = auth()
	AllowUpload   = allowUpload()
	AllowDownload = allowDownload()

	/////////

	InternalPermission = internalPermission()
)
