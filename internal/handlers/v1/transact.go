package v1handlers

import (
	"context"
	"errors"
	"io"

	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	jsoniter "github.com/json-iterator/go"
	"github.com/mtchuikov/jc-test-task/internal/domain/entities"
	"github.com/mtchuikov/jc-test-task/internal/domain/vobjects"
)

type transactor interface {
	Serve(ctx context.Context, tx vobjects.Transaction) (entities.Transaction, error)
}

type transact struct {
	transactor transactor
}

func RegisterTransact(router chi.Router, transactor transactor) {
	handler := transact{transactor}
	router.Post("/wallet", handler.Handle)
}

const transactMaxPayloadSize = 2048

var transactPayloadTooLarge = fmt.Sprintf(payloadTooLarge, transactMaxPayloadSize)

func (h *transact) Handle(rw http.ResponseWriter, req *http.Request) {
	resp := &prepareResponseArgs[TransactResponse]{
		statusCode: http.StatusBadRequest,
		respData:   &TransactResponse{},
	}
	defer prepareResponse(rw, resp)

	req.Body = http.MaxBytesReader(rw, req.Body, transactMaxPayloadSize)
	payload, err := io.ReadAll(req.Body)
	if err != nil {
		var maxBytesError *http.MaxBytesError
		if errors.As(err, &maxBytesError) {
			resp.respData.Msg = transactPayloadTooLarge
		} else {
			resp.respData.Msg = invalidRequestBody
		}

		return
	}

	var reqData TransactRequest
	err = jsoniter.Unmarshal(payload, &reqData)
	if err != nil {
		resp.respData.Msg = invalidPayloadFormatMsg
		return
	}

	tx, err := vobjects.NewTransaction(
		vobjects.NewTransactionArgs{
			WalletID:      reqData.WalletID,
			OperationType: reqData.OperationType,
			Amount:        reqData.Amount,
		})
	if err != nil {
		resp.respData.Msg = err.Error()
		return
	}

	ctx := req.Context()
	newTx, err := h.transactor.Serve(ctx, tx)
	if err != nil {
		statusCode, msg := serviceErrorToCodeAndMsg(err)
		resp.statusCode = statusCode
		resp.respData.Msg = msg
		return
	}

	resp.statusCode = http.StatusOK
	resp.respData.Success = true
	resp.respData.Msg = transactionCompletedMsg
	resp.respData.Data = newTransactionResponseData(newTx)
}
