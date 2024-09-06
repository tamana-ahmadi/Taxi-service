package errs

import "errors"

var (
	ErrValidationFailed            = errors.New("ErrValidationFailed")
	ErrPermissionDenied            = errors.New("ErrPermissionDenied")
	ErrUsernameUniquenessFailed    = errors.New("ErrUsernameUniquenessFailed")
	ErrOrdersNotFound              = errors.New("ErrOrdersNotFound")
	ErrRoutesNotFound              = errors.New("ErrRoutesNotFound")
	ErrTaxicompsNotFound           = errors.New("ErrTaxicompsNotFound")
	ErrIncorrectUsernameorPassword = errors.New("ErrIncorrectUsernameorPassword")
	ErrRecordNotFound              = errors.New("ErrRecordNotFound")
	ErrSomethingWentWrong          = errors.New("ErrSomethingWentWrong")
)
