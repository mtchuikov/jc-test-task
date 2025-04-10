package v1handlers

import (
	"net/http"

	"github.com/mtchuikov/jc-test-task/internal/domain/vobjects"
	"github.com/mtchuikov/jc-test-task/internal/repo/postgres"
	"github.com/rs/zerolog"
)

func logUnexpectedError(log zerolog.Logger, err error, op string) {
	log.Error().Err(err).Str("op", op).
		Msg("unexpected error")
}

func domainErrorToCodeAndMsg(err error) (int, string, error) {
	msg := "something went wrong"

	switch err {
	case vobjects.ErrInvalidOperationType:
		msg = vobjects.ErrInvalidOperationType.Error()
		return http.StatusBadRequest, msg, nil

	case vobjects.ErrInvalidAmount:
		msg = vobjects.ErrInvalidAmount.Error()
		return http.StatusBadRequest, msg, nil

	case vobjects.ErrInvalidTxID:
		msg = vobjects.ErrInvalidTxID.Error()
		return http.StatusBadRequest, msg, nil

	case vobjects.ErrInvalidWalletID:
		msg = vobjects.ErrInvalidWalletID.Error()
		return http.StatusBadRequest, msg, nil

	default:
		msg = "unexpected error"
		return http.StatusInternalServerError, msg, err
	}
}

func serviceErrorToCodeAndMsg(err error) (int, string, error) {
	msg := "something went wrong"

	switch err {
	case postgres.ErrFailedToBeginTx:
		return http.StatusInternalServerError, msg, nil

	case postgres.ErrFailedToInsertTx:
		return http.StatusInternalServerError, msg, nil

	case postgres.ErrFailedToCommitTx:
		return http.StatusInternalServerError, msg, nil

	case postgres.ErrNotEnoughBalance:
		msg = err.Error()
		return http.StatusBadRequest, msg, nil

	case postgres.ErrFailedToGetBalance:
		return http.StatusInternalServerError, msg, nil

	default:
		msg = "unexpected error"
		return http.StatusInternalServerError, msg, nil
	}
}
