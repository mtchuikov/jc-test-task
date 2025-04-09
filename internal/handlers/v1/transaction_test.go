package v1handlers

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	jsoniter "github.com/json-iterator/go"
	"github.com/mtchuikov/jc-test-task/internal/domain/entities"
	"github.com/mtchuikov/jc-test-task/internal/domain/vobjects"
	"github.com/stretchr/testify/suite"
)

type mockTransactor struct {
	err error
}

func (m *mockTransactor) Serve(
	ctx context.Context,
	tx vobjects.Transaction,
) (
	entities.Transaction, error,
) {
	return entities.Transaction{}, m.err
}

type testTransactSuite struct {
	suite.Suite
	transact transact
	reqData  TransactRequest
	rr       *httptest.ResponseRecorder
}

func TestTransactionSuite(t *testing.T) {
	suite.Run(t, new(testTransactSuite))
}

func (s *testTransactSuite) SetupTest() {
	transactor := &mockTransactor{}
	s.transact = transact{transactor}

	s.reqData = TransactRequest{
		WalletID:      "c1caa508-6a7b-4381-bcca-afca5ab8ad8f",
		OperationType: vobjects.DepositTx,
		Amount:        1337,
	}

	s.rr = httptest.NewRecorder()
}

type testTransactArgs struct {
	ExpecteCode        int
	ExpecteSuccess     bool
	ExpecteMsgContains string
}

func (s *testTransactSuite) testTransact(args testTransactArgs) {
	payload, _ := jsoniter.Marshal(s.reqData)
	body := bytes.NewBuffer(payload)

	req := httptest.NewRequest(http.MethodPost, "/", body)
	s.transact.Handle(s.rr, req)

	s.Require().Equalf(
		args.ExpecteCode, s.rr.Code,
		"handler returned invalid code",
		args.ExpecteCode,
	)

	payload, err := io.ReadAll(s.rr.Body)
	s.Require().NoError(err, "response body must be read")

	var respData TransactResponse
	err = jsoniter.Unmarshal(payload, &respData)
	s.Require().NoError(err, "response body must be marshalled")

	s.Equal(args.ExpecteSuccess, respData.Success)

	s.Require().Contains(
		respData.Msg, args.ExpecteMsgContains,
		"wrong message provided",
	)
}

func (s *testTransactSuite) TestTransaction_Success() {
	s.testTransact(
		testTransactArgs{
			ExpecteCode:        http.StatusOK,
			ExpecteSuccess:     true,
			ExpecteMsgContains: transactionCompletedMsg,
		})
}

func (s *testTransactSuite) TestTransaction_InvalidWalletID() {
	s.reqData.WalletID = ""

	s.testTransact(
		testTransactArgs{
			ExpecteCode:        http.StatusBadRequest,
			ExpecteSuccess:     false,
			ExpecteMsgContains: vobjects.ErrInvalidWalletID.Error(),
		})
}

func (s *testTransactSuite) TestTransaction_InvalidOperationType() {
	s.reqData.OperationType = ""

	s.testTransact(
		testTransactArgs{
			ExpecteCode:        http.StatusBadRequest,
			ExpecteSuccess:     false,
			ExpecteMsgContains: vobjects.ErrInvalidOperationType.Error(),
		})
}

func (s *testTransactSuite) TestTransaction_InvalidLargeBody() {
	s.reqData.WalletID = strings.Repeat("large body", 2048)

	s.testTransact(
		testTransactArgs{
			ExpecteCode:        http.StatusBadRequest,
			ExpecteSuccess:     false,
			ExpecteMsgContains: transactPayloadTooLarge,
		})
}
