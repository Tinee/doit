package validation

const errorMsg = "Validation Error"

// Error represents when a struct is invalid
type Error struct {
	s    string
	errs map[string]string
}

// New returns an ValidationError struct
func New() *Error {
	return &Error{
		s:    errorMsg,
		errs: make(map[string]string),
	}
}

// Add inserts a the validation error that just occured.
func (ve *Error) Add(field, text string) {
	ve.errs[field] = text
}

// Error returns the validation error
func (ve *Error) Error() string {
	return ve.s
}
