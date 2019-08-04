package errnum

var (
	//Common errors
	OK                  = &ErrNum{Code: 0, Message: "OK"}
	InternalServerError = &ErrNum{Code: 10001, Message: "Internal server error."}
	ErrBind             = &ErrNum{Code: 10002, Message: "Error occurred while binding the request body to the struct."}

	//User errors
	ErrValidation        = &ErrNum{Code: 20001, Message: "Validation failed."}
	ErrDatabase          = &ErrNum{Code: 20002, Message: "Database error."}
	ErrToken             = &ErrNum{Code: 20003, Message: "Error occurred while signing the JSON web token."}
	ErrEncrypt           = &ErrNum{Code: 20101, Message: "Error occurred while encrypting the user password."}
	ErrUserNotFound      = &ErrNum{Code: 20102, Message: "The user was not found."}
	ErrTokenInvalid      = &ErrNum{Code: 20103, Message: "The token was invalid."}
	ErrPasswordIncorrect = &ErrNum{Code: 20104, Message: "The password was incorrect."}
)
