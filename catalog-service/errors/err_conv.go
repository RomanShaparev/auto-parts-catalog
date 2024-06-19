package errors

import (
	"errors"
	"strings"

	"github.com/jackc/pgx/v5/pgconn"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func StorageToGrpcErr(err error) error {
	var errCode codes.Code
	var message string
	var pgErr *StorageError

	if errors.As(err, &pgErr) {
		switch pgErr.Code {
		case InvalidArgumentError:
			errCode = codes.InvalidArgument
		case NotFoundError:
			errCode = codes.NotFound
		case AlreadyExistsError:
			errCode = codes.AlreadyExists
		case InternalError:
			errCode = codes.Internal
		default:
			errCode = codes.Unknown
		}
		message = pgErr.Error()
	} else {
		errCode = codes.Internal
		message = ""
	}
	return status.Errorf(errCode, message)
}

func PgToStorageErr(err error) error {
	errCode := InternalError
	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		// Class 23 — Integrity Constraint Violation
		if strings.HasPrefix(pgErr.Code, "23") {
			if pgErr.Code == "23505" {
				errCode = AlreadyExistsError
			} else {
				errCode = InvalidArgumentError
			}
		}
	}
	return &StorageError{Code: errCode, Message: errCode.String()}
}
