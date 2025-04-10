package v1handlers

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/mtchuikov/jc-test-task/internal/domain/entities"
	"github.com/mtchuikov/jc-test-task/internal/domain/vobjects"
	"github.com/rs/zerolog"
)

type balanceGetter interface {
	Serve(ctx context.Context, walletID vobjects.WalletID) (entities.Balance, error)
}

type getBalance struct {
	log           zerolog.Logger
	balanceGetter balanceGetter
}

type RegisterGetBalanceParams struct {
	Log             zerolog.Logger
	BabalanceGetter balanceGetter
	Router          chi.Router
}

func RegisterGetBalance(args RegisterGetBalanceParams) {
	handler := getBalance{
		log:           args.Log,
		balanceGetter: args.BabalanceGetter,
	}

	args.Router.Get("/wallet/{wallet_uuid}", handler.Handle)
}

func (h *getBalance) Handle(rw http.ResponseWriter, req *http.Request) {
	resp := &prepareResponseArgs[GetBalanceResponse]{
		statusCode: http.StatusBadRequest,
		response:   &GetBalanceResponse{},
	}
	defer prepareResponse(rw, resp)

	id := chi.URLParam(req, "wallet_uuid")
	walletID, err := vobjects.NewWalletID(id)
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
	balance, err := h.balanceGetter.Serve(ctx, walletID)
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
	resp.response.Msg = balanceFetchedMsg
	resp.response.Data = newGetBalanceData(balance)
}
