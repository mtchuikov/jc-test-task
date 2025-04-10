package v1handlers

import (
	"context"
	"errors"
	"io"

	"net/http"

	"github.com/go-chi/chi/v5"
	jsoniter "github.com/json-iterator/go"
	"github.com/mtchuikov/jc-test-task/internal/domain/entities"
	"github.com/mtchuikov/jc-test-task/internal/domain/vobjects"
	"github.com/rs/zerolog"
)

type transactor interface {
	Serve(ctx context.Context, tx vobjects.Transaction) (entities.Transaction, error)
}

type transact struct {
	log        zerolog.Logger
	transactor transactor
}

type RegisterTransactParams struct {
	Log        zerolog.Logger
	Transactor transactor
	Router     chi.Router
}

func RegisterTransact(args RegisterTransactParams) {
	handler := transact{
		log:        args.Log,
		transactor: args.Transactor,
	}

	args.Router.Post("/wallet", handler.Handle)
}

const transactMaxPayloadSize = 2048

const transactHandlerOp = "handlers.transact.handle"

func (h *transact) Handle(rw http.ResponseWriter, req *http.Request) {
	resp := &prepareResponseArgs[TransactResponse]{
		statusCode: http.StatusBadRequest,
		response:   &TransactResponse{},
	}
	defer prepareResponse(rw, resp)

	req.Body = http.MaxBytesReader(rw, req.Body, transactMaxPayloadSize)
	payload, err := io.ReadAll(req.Body)
	if err != nil {
		var maxBytesError *http.MaxBytesError
		if errors.As(err, &maxBytesError) {
			resp.response.Msg = payloadTooLarge
		} else {
			resp.response.Msg = invalidRequestBody
		}

		return
	}

	var reqData TransactRequest
	err = jsoniter.Unmarshal(payload, &reqData)
	if err != nil {
		resp.response.Msg = invalidPayloadMsg
		return
	}

	tx, err := vobjects.NewTransaction(
		vobjects.NewTransactionArgs{
			WalletID:      reqData.WalletID,
			OperationType: reqData.OperationType,
			Amount:        reqData.Amount,
		})
	if err != nil {
		statusCode, msg, err := domainErrorToCodeAndMsg(err)
		if err != nil {
			logUnexpectedError(h.log, err, transactHandlerOp)
		}

		resp.statusCode = statusCode
		resp.response.Msg = msg
		return
	}

	ctx := req.Context()
	newTx, err := h.transactor.Serve(ctx, tx)
	if err != nil {
		statusCode, msg, err := serviceErrorToCodeAndMsg(err)
		if err != nil {
			logUnexpectedError(h.log, err, transactHandlerOp)
		}

		resp.statusCode = statusCode
		resp.response.Msg = msg
		return
	}

	resp.statusCode = http.StatusOK
	resp.response.Success = true
	resp.response.Msg = txCompletedMsg
	resp.response.Data = newTransactionResponseData(newTx)
}
