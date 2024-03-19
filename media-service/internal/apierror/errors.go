package apierror

type NotFoundError struct {
	Message string
}

type InternalServerError struct {
	Message string
}

type BadRequestError struct {
	Message string
}

type UnauthorizedRequestError struct {
	Message string
}

func (e NotFoundError) Error() string {
	return e.Message
}

func (e InternalServerError) Error() string {
	return e.Message
}

func (e BadRequestError) Error() string {
	return e.Message
}

func NewNotFoundError(message string) NotFoundError {
	return NotFoundError{Message: message}
}

func NewInternalServerError(message string) InternalServerError {
	return InternalServerError{Message: message}
}

func NewBadRequestError(message string) BadRequestError {
	return BadRequestError{Message: message}
}

func NewUnauthorizedRequestError(message string) BadRequestError {
	return BadRequestError{Message: message}
}
