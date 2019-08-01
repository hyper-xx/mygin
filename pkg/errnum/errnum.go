package errnum

import "fmt"

type ErrNum struct {
	Code    int
	Message string
}

//ErrNum.Error() function
func (err ErrNum) Error() string {
	return err.Message
}

type Err struct {
	Code    int
	Message string
	Err     error
}

func (err *Err) Error() string {
	return fmt.Sprintf("Err - code: %d, message: %s, error: %s", err.Code, err.Message, err.Err)
}

func New(errnum *ErrNum, err error) *Err {
	return &Err{Code: errnum.Code, Message: errnum.Message, Err: err}
}

func (err *Err) Add(message string) error {
	err.Message += " " + message
	return err
}

func DecodeErr(err error) (int, string) {
	if err == nil {
		return OK.Code, OK.Message
	}

	switch typed := err.(type) {
	case *Err:
		return typed.Code, typed.Message
	case *ErrNum:
		return typed.Code, typed.Message
	default:
	}

	return InternalServerError.Code, err.Error()
}

func IsErrUserNotFound(err error) bool {
	code, _ := DecodeErr(err)
	return code == ErrUserNotFound.Code
}
