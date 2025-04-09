package v1handlers

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/mtchuikov/jc-test-task/internal/domain/entities"
	"github.com/mtchuikov/jc-test-task/internal/domain/vobjects"
)

type balanceGetter interface {
	Serve(ctx context.Context, walletID vobjects.WalletID) (entities.Balance, error)
}

type getBalance struct {
	balanceGetter balanceGetter
}

func RegisterGetBalance(router chi.Router, babalanceGetter balanceGetter) {
	handler := getBalance{babalanceGetter}
	router.Get("/wallet/{wallet_uuid}", handler.Handle)
}

func (h *getBalance) Handle(rw http.ResponseWriter, req *http.Request) {
	resp := &prepareResponseArgs[GetBalanceResponse]{
		statusCode: http.StatusBadRequest,
		respData:   &GetBalanceResponse{},
	}
	defer prepareResponse(rw, resp)

	id := chi.URLParam(req, "wallet_uuid")
	walletID, err := vobjects.NewWalletID(id)
	if err != nil {
		resp.respData.Msg = vobjects.ErrInvalidWalletID.Error()
		return
	}

	ctx := req.Context()
	balance, err := h.balanceGetter.Serve(ctx, walletID)
	if err != nil {
		statusCode, msg := serviceErrorToCodeAndMsg(err)
		resp.statusCode = statusCode
		resp.respData.Msg = msg
		return
	}

	resp.statusCode = http.StatusOK
	resp.respData.Success = true
	resp.respData.Msg = balanceGotMsg
	resp.respData.Data = newGetBalanceData(balance)
}
