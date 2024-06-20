package grpcimpl

import (
	"auto-parts-catalog/backend/catalogservice"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func grpcToServiceErr(err error) error {
	if err == nil {
		return nil
	}

	var errCode catalogservice.CatalogServiceErrorCode

	switch status.Code(err) {
	case codes.InvalidArgument:
		errCode = catalogservice.InvalidArgumentError
	case codes.NotFound:
		errCode = catalogservice.NotFoundError
	case codes.AlreadyExists:
		errCode = catalogservice.AlreadyExistsError
	case codes.Internal:
		errCode = catalogservice.InternalError
	default:
		errCode = catalogservice.UnknownError
	}

	return &catalogservice.CatalogServiceError{Code: errCode}
}
