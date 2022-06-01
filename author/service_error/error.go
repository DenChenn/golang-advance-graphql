package service_error

type WrappedError struct {
	Message   string
	ErrorCode string
}

func (w *WrappedError) Error() string {
	return w.ErrorCode
}

var (
	NoAuthorFoundError = &WrappedError{
		Message:   "No author with this id.",
		ErrorCode: "NO_AUTHOR_FOUND_ERROR",
	}
)
