package account

type BalanceUpdate struct {
	BalanceChange int64 `json:"balancechange" binding:"required"`
}
