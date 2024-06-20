package catalogservice

type CatalogServiceError struct {
	Code CatalogServiceErrorCode
}

func (e *CatalogServiceError) Error() string {
	return e.Code.String()
}

type CatalogServiceErrorCode int

const (
	InvalidArgumentError CatalogServiceErrorCode = iota
	NotFoundError
	InternalError
	AlreadyExistsError
	UnknownError
)

func (c CatalogServiceErrorCode) String() string {
	return [...]string{"InvalidArgumentError", "NotFoundError", "InternalError", "AlreadyExistsError", "UnknownError"}[c]
}
