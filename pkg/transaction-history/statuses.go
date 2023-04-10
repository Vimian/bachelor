package transactionhistory

type AStatus struct {
	ID   int32  `json:"id"`
	Name string `json:"name"`
}

type Statuses struct {
	Statuses []AStatus `json:"statuses"`
}
