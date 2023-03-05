package models

type Token struct {
	Type           string  `json:"type"`
	Timestamp      int64   `json:"timestamp"`
	BlockTimestamp int64   `json:"blockTimestamp"`
	MintAddress    string  `json:"mintAddress"`
	Symbol         string  `json:"symbol"`
	Name           string  `json:"name"`
	Price          float64 `json:"price"`
	FloorPrice     float64 `json:"floorPrice"`
	TokenAddress   string  `json:"tokenAddress"`
	Seller         string  `json:"seller"`
}
