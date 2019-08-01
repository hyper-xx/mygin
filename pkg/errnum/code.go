package errnum

var (
	//Common errors
	OK                  = &ErrNum{Code: 0, Message: "OK"}
	InternalServerError = &ErrNum{Code: 10001, Message: "Internal server error."}
	ErrBind             = &ErrNum{Code: 10002, Message: "Error occurred while binding the request body to the struct."}

	//User errors
	ErrUserNotFound = &ErrNum{Code: 20102, Message: "The user was not found."}
)
