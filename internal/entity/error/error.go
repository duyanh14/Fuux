package error

import "errors"

type Error struct {
	Code    int
	Message string
	Error   error
}

func NewError(code int, message string) *Error {
	return &Error{Code: code, Message: message, Error: errors.New(message)}
}

func ExposeError(err error, es ...*Error) *Error {
	for _, e := range es {
		if errors.Is(err, e.Error) {
			return e
		}
	}
	return NewError(0, err.Error())
}

var (
	Unknown                       = NewError(10000, "Unknown error")
	FieldRequired                 = NewError(10002, "Field required")
	RecordNotFound                = NewError(10003, "Record not found")
	PermissionDenied              = NewError(10004, "Permission denied")
	PathAlreadyUse                = NewError(10017, "Path already in use")
	NameAlreadyUse                = NewError(10019, "Name already in use")
	PathExist                     = NewError(20000, "Path exist in database")
	NameExist                     = NewError(20001, "Name exist in database")
	PathRecordNotFound            = NewError(20002, "Path record not found")
	ResourceAccessStatusIsDisable = NewError(30002, "Resource Access is disabled")
	UploadDisallow                = NewError(30003, "Resource Access Upload disallow")
	DownloadDisallow              = NewError(30004, "Resource Access Download disallow")
)
