package transactionLibrary

const (
	ERR_MISSING_FROM              = "[From] parameter is required"
	ERR_MISSING_TO                = "[TO] parameter is required"
	ERR_INVALID_AMOUNT            = "[Amount] parameter should be positive number"
	ERR_INVALID_TRANSFER_SAME_ACC = "invalid operation. you are trying to transfer from and to the same account"
	ERR_BALANCE_LESS_THAN_AMOUNT  = "account's balance is less than the transaction amount"
)
