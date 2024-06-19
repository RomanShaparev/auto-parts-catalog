package errors

import (
	"errors"
	"log"
	"strings"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func StorageToGrpcErr(err error) error {
	var errCode codes.Code
	var message string
	var storageErr *StorageError

	if errors.As(err, &storageErr) {
		switch storageErr.Code {
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
		message = storageErr.Error()
	} else {
		errCode = codes.Internal
		message = ""
	}
	return status.Errorf(errCode, message)
}

func PgToStorageErr(err error) error {
	log.Print(err)
	errCode := InternalError
	var pgErr *pgconn.PgError

	if errors.Is(err, pgx.ErrNoRows) {
		errCode = NotFoundError
	} else if errors.As(err, &pgErr) {
		log.Print(err)

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
