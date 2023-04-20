package user

type Users struct {
	Users  []User `json:"users"`
	Page   int    `json:"page"`
	Amount int    `json:"amount"`
	Offset int    `json:"offset"`
}
