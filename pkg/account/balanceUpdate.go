package account

type BalanceUpdate struct {
	BalanceChange   int64 `json:"balancechange" binding:"required"`
	TransactionType int32 `json:"transactiontype" binding:"required"`
}
