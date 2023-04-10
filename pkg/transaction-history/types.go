package transactionhistory

type Type struct {
	ID   int32  `json:"id"`
	Name string `json:"name"`
}

type Types struct {
	Types []Type `json:"types"`
}
