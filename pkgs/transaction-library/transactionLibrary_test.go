package transactionLibrary

import (
	"testing"

	"github.com/osamamosaad/paytabs/entities"
	accountLibrary "github.com/osamamosaad/paytabs/pkgs/account-library"
	"github.com/osamamosaad/paytabs/storage"
	"github.com/osamamosaad/paytabs/utils"
)

func setup() {
	storage.DBMemory = make(map[string]map[string]interface{})
	storage.DBMemory["Account"] = make(map[string]interface{})
	storage.DBMemory["Account"]["17f904c1-806f-4252-9103-74e7a5d3e340"] = &entities.Account{
		ID:      "17f904c1-806f-4252-9103-74e7a5d3e340",
		Balance: 5000.50,
		Name:    "Osama",
	}
	storage.DBMemory["Account"]["fd796d75-1bcf-4a95-bf1a-f7b296adb79f"] = &entities.Account{
		ID:      "fd796d75-1bcf-4a95-bf1a-f7b296adb79f",
		Balance: 3708.11,
		Name:    "Wikizz",
	}
}

func TestEntityDataValidationFromMissing(t *testing.T) {
	setup()

	storage := storage.New()
	transactionLib := New(
		mockTransactionRepo{},
		accountLibrary.New(&mockAccountRepo{storage: storage}),
	)

	transaction := entities.Transaction{
		From:   "",
		To:     "17f904c1-806f-4252-9103-74e7a5d3e340",
		Amount: 100,
	}
	_, err := transactionLib.Transfer(transaction)

	if err == nil || err.Error() != ERR_MISSING_FROM {
		t.Error("[Form] Validation error is missing")
	}
}

func TestEntityDataValidationToMissing(t *testing.T) {
	setup()
	storage := storage.New()
	transactionLib := New(
		mockTransactionRepo{},
		accountLibrary.New(&mockAccountRepo{storage: storage}),
	)
	transaction := entities.Transaction{
		From:   "17f904c1-806f-4252-9103-74e7a5d3e340",
		To:     "",
		Amount: 100,
	}
	_, err := transactionLib.Transfer(transaction)

	if err == nil || err.Error() != ERR_MISSING_TO {
		t.Error("[To] Validation error is missing")
	}
}

func TestEntityDataValidationAmountExists(t *testing.T) {
	setup()

	storage := storage.New()
	transactionLib := New(
		mockTransactionRepo{},
		accountLibrary.New(&mockAccountRepo{storage: storage}),
	)

	transaction := entities.Transaction{
		From:   "17f904c1-806f-4252-9103-74e7a5d3e340",
		To:     "17f904c1-806f-4252-9103-74e7a526sd52",
		Amount: -1,
	}
	_, err := transactionLib.Transfer(transaction)

	if err == nil || err.Error() != ERR_INVALID_AMOUNT {
		t.Error("[Amount] is missing or not positive number")
	}
}

func TestFiringErrorWhenTryingToTransferForTheSameAccount(t *testing.T) {
	setup()

	storage := storage.New()
	transactionLib := New(
		mockTransactionRepo{},
		accountLibrary.New(&mockAccountRepo{storage: storage}),
	)

	transaction := entities.Transaction{
		From:   "17f904c1-806f-4252-9103-74e7a5d3e340",
		To:     "17f904c1-806f-4252-9103-74e7a5d3e340",
		Amount: 100,
	}
	_, err := transactionLib.Transfer(transaction)

	if err == nil || err.Error() != ERR_INVALID_TRANSFER_SAME_ACC {
		t.Error("Missing error same account transfer operation")
	}
}

func TestErrorFromAccountHasNoEenoughBalance(t *testing.T) {
	setup()

	transactionLib := New(
		mockTransactionRepo{},
		accountLibrary.New(getMockAccountRepo(storage.New(), false)),
	)

	transaction := entities.Transaction{
		From:   "Fake-ID",
		To:     "17f904c1-806f-4252-9103-74e7a5d3e340",
		Amount: 100,
	}

	_, err := transactionLib.Transfer(transaction)

	if err == nil || err.Error() != "Account not found" {
		t.Error("Error 'Account not found' not exists")
	}
}

func TestErrorTooAccountHasNoEenoughBalance(t *testing.T) {
	setup()

	transactionLib := New(
		mockTransactionRepo{},
		accountLibrary.New(getMockAccountRepo(storage.New(), false)),
	)

	transaction := entities.Transaction{
		From:   "17f904c1-806f-4252-9103-74e7a5d3e340",
		To:     "Fake-ID",
		Amount: 100,
	}

	_, err := transactionLib.Transfer(transaction)

	if err == nil || err.Error() != "Account not found" {
		t.Error("Error 'Account not found' not exists")
	}
}

func TestTranfareWorkCorrectlyWithNoError(t *testing.T) {
	setup()

	transactionLib := New(
		mockTransactionRepo{},
		accountLibrary.New(getMockAccountRepo(storage.New(), false)),
	)

	transaction := entities.Transaction{
		From:   "17f904c1-806f-4252-9103-74e7a5d3e340",
		To:     "fd796d75-1bcf-4a95-bf1a-f7b296adb79f",
		Amount: 100,
	}

	_, err := transactionLib.Transfer(transaction)

	if err != nil {
		t.Error("Transfare doesn't work correctly")
	}
}

func TestTranfareReturnTransactionEntityAsExpect(t *testing.T) {
	setup()

	transactionLib := New(
		mockTransactionRepo{},
		accountLibrary.New(getMockAccountRepo(storage.New(), false)),
	)

	transaction := entities.Transaction{
		From:   "17f904c1-806f-4252-9103-74e7a5d3e340",
		To:     "fd796d75-1bcf-4a95-bf1a-f7b296adb79f",
		Amount: 100,
	}

	actual, err := transactionLib.Transfer(transaction)

	if err != nil {
		t.Error("Transfare doesn't work correctly")
	}

	expect := entities.Transaction{
		From:   "17f904c1-806f-4252-9103-74e7a5d3e340",
		To:     "fd796d75-1bcf-4a95-bf1a-f7b296adb79f",
		Amount: 100,
	}

	if actual.ID == "" {
		t.Error("Transfare doesn't work as expected, ID is empty")
	}

	if actual.Amount != expect.Amount {
		t.Errorf("Transfare doesn't work as expected, [Amount] is not the same as expected; \n Expected: %v \n Actual: %v", expect.Amount, actual.Amount)
	}

	if actual.From != expect.From {
		t.Errorf("Transfare doesn't work as expected, [From] is not the same as expected; \n Expected: %v \n Actual: %v", expect.From, actual.From)
	}

	if actual.To != expect.To {
		t.Errorf("Transfare doesn't work as expected, [To] is not the same as expected; \n Expected: %v \n Actual: %v", expect.To, actual.To)
	}
}

func TestTransfareMoreThanAccountBalance(t *testing.T) {
	setup()

	accLib := accountLibrary.New(getMockAccountRepo(storage.New(), true))
	transactionLib := New(
		mockTransactionRepo{},
		accLib,
	)

	transaction := entities.Transaction{
		From:   "17f904c1-806f-4252-9103-74e7a5d3e340",
		To:     "fd796d75-1bcf-4a95-bf1a-f7b296adb79f",
		Amount: 1000000,
	}

	_, _ = transactionLib.Transfer(transaction)

	accFrom, err := accLib.FindById(transaction.From)
	if err == nil && accFrom.Balance < 0 {
		t.Error("Account that transferred from, transfere an amount that is more than that he has in his account")
	}
}

func TestAccountsAfterTransfare(t *testing.T) {
	setup()

	accLib := accountLibrary.New(getMockAccountRepo(storage.New(), true))
	transactionLib := New(
		mockTransactionRepo{},
		accLib,
	)

	transaction := entities.Transaction{
		From:   "17f904c1-806f-4252-9103-74e7a5d3e340",
		To:     "fd796d75-1bcf-4a95-bf1a-f7b296adb79f",
		Amount: 5000,
	}

	accFromBefore, _ := accLib.FindById(transaction.From)
	accTooBefore, _ := accLib.FindById(transaction.To)

	accFromBalanceExpect := utils.RoundFloat(accFromBefore.Balance-transaction.Amount, 2)
	accTooBalanceExpect := utils.RoundFloat(accTooBefore.Balance+transaction.Amount, 2)

	transactionLib.Transfer(transaction)

	accFrom, _ := accLib.FindById(transaction.From)
	accTo, _ := accLib.FindById(transaction.To)
	if accFrom.Balance != accFromBalanceExpect && accTo.Balance != accTooBalanceExpect {
		t.Errorf(
			"The Transfer not effected well. \n Account 'transfer from' balance: \n\t Expected: %v \n\t Actual: %v \n Account 'transfer To' balance: \n\t Expected: %v \n\t Actual: %v",
			accFromBalanceExpect,
			accFrom.Balance,
			accTooBalanceExpect,
			accTo.Balance,
		)
	}
}
