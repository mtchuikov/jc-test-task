package v1handlers

import (
	"net/http"

	"github.com/mtchuikov/jc-test-task/internal/repo/postgres"
)

func serviceErrorToCodeAndMsg(err error) (int, string) {
	msg := "something went wrong"

	switch err {
	case postgres.ErrFailedToBeginTx:
		return http.StatusInternalServerError, msg

	case postgres.ErrFailedToInsertTx:
		return http.StatusInternalServerError, msg

	case postgres.ErrFailedToCommitTx:
		return http.StatusInternalServerError, msg

	case postgres.ErrNotEnoughBalance:
		msg = err.Error()
		return http.StatusBadRequest, msg

	case postgres.ErrFailedToGetBalance:
		return http.StatusInternalServerError, msg

	default:
		msg = "unexpected error"
		return http.StatusInternalServerError, msg
	}
}
