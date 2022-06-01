package service_error

type WrappedError struct {
	Message   string
	ErrorCode string
}

func (w *WrappedError) Error() string {
	return w.ErrorCode
}

var (
	NoDiaryFoundError = &WrappedError{
		Message:   "No diary with this id.",
		ErrorCode: "NO_DIARY_FOUND_ERROR",
	}
)
