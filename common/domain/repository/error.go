package repository

// ErrorNotAffected
type ErrorNotAffected struct {
	error
}

func NewErrorNotAffected(err error) ErrorNotAffected {
	return ErrorNotAffected{err}
}

// ErrorDuplicateCreating
type ErrorDuplicateCreating struct {
	error
}

func NewErrorDuplicateCreating(err error) ErrorDuplicateCreating {
	return ErrorDuplicateCreating{err}
}

// ErrorResourceNotExists
type ErrorResourceNotExists struct {
	error
}

func NewErrorResourceNotExists(err error) ErrorResourceNotExists {
	return ErrorResourceNotExists{err}
}

// ErrorConcurrentUpdating
type ErrorConcurrentUpdating struct {
	error
}

func NewErrorConcurrentUpdating(err error) ErrorConcurrentUpdating {
	return ErrorConcurrentUpdating{err}
}

// ErrorConstraintViolated
type ErrorConstraintViolated struct {
	error
}

func NewErrorConstraintViolated(err error) ErrorConstraintViolated {
	return ErrorConstraintViolated{err}
}

// helper
func IsErrorNotAffected(err error) bool {
	_, ok := err.(ErrorNotAffected)

	return ok
}

func IsErrorResourceNotExists(err error) bool {
	_, ok := err.(ErrorResourceNotExists)

	return ok
}

func IsErrorDuplicateCreating(err error) bool {
	_, ok := err.(ErrorDuplicateCreating)

	return ok
}

func IsErrorConcurrentUpdating(err error) bool {
	_, ok := err.(ErrorConcurrentUpdating)

	return ok
}

func IsErrorConstraintViolated(err error) bool {
	_, ok := err.(ErrorConstraintViolated)

	return ok
}
