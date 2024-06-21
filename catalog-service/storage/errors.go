package storage

type StorageError struct {
	Code    StorageErrorCode
	Message string
}

func (e *StorageError) Error() string {
	return e.Code.String()
}

type StorageErrorCode int

const (
	InvalidArgumentError StorageErrorCode = iota
	NotFoundError
	InternalError
	AlreadyExistsError
)

func (c StorageErrorCode) String() string {
	return [...]string{"InvalidArgumentError", "NotFoundError", "InternalError", "AlreadyExistsError"}[c]
}
