package web

type ItemRequest []struct {
	Qty    int     `json:"qty"`
	Price  float64 `json:"price"`
	Amount float64 `json:"amount"`
}