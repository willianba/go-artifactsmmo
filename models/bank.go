package models

type Bank struct {
	Slots             int `json:"slots"`
	Expansions        int `json:"expansions"`
	NextExpansionCost int `json:"next_expansion_cost"`
	Gold              int `json:"gold"`
}

type BankExtension struct {
	Price int `json:"price"`
}
