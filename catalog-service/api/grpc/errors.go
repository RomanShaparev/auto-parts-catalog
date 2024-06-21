package grpc

import (
	"auto-parts-catalog/catalog-service/storage"
	"errors"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func StorageToGrpcErr(err error) error {
	if err == nil {
		return nil
	}

	var errCode codes.Code
	var message string
	var storageErr *storage.StorageError

	if errors.As(err, &storageErr) {
		switch storageErr.Code {
		case storage.InvalidArgumentError:
			errCode = codes.InvalidArgument
		case storage.NotFoundError:
			errCode = codes.NotFound
		case storage.AlreadyExistsError:
			errCode = codes.AlreadyExists
		case storage.InternalError:
			errCode = codes.Internal
		default:
			errCode = codes.Unknown
		}
		message = storageErr.Error()
	} else {
		errCode = codes.Internal
		message = ""
	}
	return status.Errorf(errCode, message)
}
