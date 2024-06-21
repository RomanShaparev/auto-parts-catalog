package postgres

import (
	"auto-parts-catalog/catalog-service/storage"
	"errors"
	"strings"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

func PgToStorageErr(err error) error {
	if err == nil {
		return nil
	}

	errCode := storage.InternalError
	var pgErr *pgconn.PgError

	if errors.Is(err, pgx.ErrNoRows) {
		errCode = storage.NotFoundError
	} else if errors.As(err, &pgErr) {

		// Class 23 — Integrity Constraint Violation
		if strings.HasPrefix(pgErr.Code, "23") {
			if pgErr.Code == "23505" {
				errCode = storage.AlreadyExistsError
			} else {
				errCode = storage.InvalidArgumentError
			}
		}
	}
	return &storage.StorageError{Code: errCode, Message: errCode.String()}
}
