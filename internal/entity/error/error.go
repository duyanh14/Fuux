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
	Unknown                         = NewError(10000, "Unknown error")
	FieldRequired                   = NewError(10002, "Field required")
	RecordNotFound                  = NewError(10003, "Record not found")
	UserAccountBanned               = NewError(10013, "User account not found")
	PermissionDenied                = NewError(10004, "Permission denied")
	PaymentMethodNotSupported       = NewError(10005, "Payment method not supported")
	PaymentNotFound                 = NewError(10006, "Payment not found")
	PaymentTransactionNotFound      = NewError(10007, "Payment transaction not found")
	PaymentDone                     = NewError(10008, "Payment is done")
	PaymentAmountNotMatch           = NewError(10009, "Payment amount not match")
	PaymentAmountNotReachedMinimum  = NewError(10010, "Payment amount has not been reached minimum")
	UserAccountWrongPassword        = NewError(10011, "Wrong password")
	UserAccountSaveFailed           = NewError(10012, "Account save failed")
	UserAccountResetPasswordExpired = NewError(10014, "Reset password request has expired")
	EmailInvalidate                 = NewError(10015, "Email invalidate")
	PhoneNumberInvalidate           = NewError(10016, "Phone number invalidate")
	PathAlreadyUse                  = NewError(10017, "Path already in use")
	PhoneNumberAlreadyUse           = NewError(10018, "Phone number already in use")
	NameAlreadyUse                  = NewError(10019, "Name already in use")
	SomethingWrong                  = NewError(10020, "Something wrong")
	HardwareSyncFailed              = NewError(10021, "Hardware sync failed")
	HardwareNotFound                = NewError(10022, "Hardware not found")
	UserAccountVerifyRequired       = NewError(10023, "Account verify required")
	UserAccountNotEnoughBalance     = NewError(10024, "Account not enough balance")
	GameNotForRent                  = NewError(10025, "Game not for rent")
	GameOutOfRentalAccount          = NewError(10026, "Game out of rental account")
	GameNotPlay                     = NewError(10027, "Game not play")
	UserRoleSaveFailed              = NewError(10028, "User role save failed")
	UserRoleNotFound                = NewError(10029, "User role not found")
	SomeRoleAlreadyAssign           = NewError(10030, "Some role already assign")
	SomeRoleNotAssign               = NewError(10031, "Some role not assign")
	UserVerifyRequest               = NewError(10032, "User verify request")
	UserVerifyRequestExpired        = NewError(10033, "User verify request expired")
	GameGenreNotFound               = NewError(10034, "Game genre not found")
	GameGenreSaveFailed             = NewError(10035, "Game genre save failed")
	AccessDenied                    = NewError(10036, "Access denied")
	NotEnoughCoin                   = NewError(10037, "Not enough coin")
	NotFound                        = NewError(10038, "Not found")

	PathExist          = NewError(20000, "Path exist in database")
	NameExist          = NewError(20001, "Name exist in database")
	PathRecordNotFound = NewError(20002, "Path record not found")
)
