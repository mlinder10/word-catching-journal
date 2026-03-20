package types

type APIError struct {
	Message      string `json:"message"`
	Code         int    `json:"code"`
	WrappedError error
}

func (e APIError) Error() string {
	return e.Message
}

type errorList struct {
	None APIError

	// 401
	Unauthenticated        func(error) APIError
	InvalidEmailOrPassword func(error) APIError

	// 400
	InvalidResetCode     func(error) APIError
	PasswordsDoNotMatch  func(error) APIError
	EmailAlreadyInUse    func(error) APIError
	UsernameAlreadyInUse func(error) APIError
	FailedToParseRequest func(error) APIError

	// 500
	InternalServerError func(error) APIError
	FailedToDefineWord  func(error) APIError
	Unimplemented       APIError
}

func newError(code int, message string, err error) APIError {
	return APIError{
		Message:      message,
		Code:         code,
		WrappedError: err,
	}
}

var Errors = errorList{
	None: APIError{},
	// 401
	Unauthenticated:        func(err error) APIError { return newError(401, "unauthenticated", err) },
	InvalidEmailOrPassword: func(err error) APIError { return newError(401, "invalid email or password", err) },

	// 400
	InvalidResetCode:     func(err error) APIError { return newError(400, "invalid reset code", err) },
	PasswordsDoNotMatch:  func(err error) APIError { return newError(400, "passwords do not match", err) },
	EmailAlreadyInUse:    func(err error) APIError { return newError(400, "email already in use", err) },
	UsernameAlreadyInUse: func(err error) APIError { return newError(400, "username already in use", err) },
	FailedToParseRequest: func(err error) APIError { return newError(400, "failed to parse request", err) },
	// 500
	InternalServerError: func(err error) APIError { return newError(500, "internal server error", err) },
	FailedToDefineWord:  func(err error) APIError { return newError(500, "failed to define word", err) },
	Unimplemented:       newError(501, "unimplemented", nil),
}
