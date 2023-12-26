package web

type ItemRequest []struct {
	Qty    float64 `json:"qty"`
	Price  float64 `json:"price"`
	Amount float64 `json:"amount"`
}
