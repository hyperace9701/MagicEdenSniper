package models

type NFT struct {
	MintAddress string `json:"mintAddress"`
	Collection  string `json:"collection"`
	Name        string `json:"name"`
}
