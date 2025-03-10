package models

type GEOrder struct {
	Id        int    `json:"id"`
	Seller    string `json:"seller"`
	Code      string `json:"code"`
	Quantity  int    `json:"quantity"`
	Price     int    `json:"price"`
	CreatedAt string `json:"created_at"`
}

type GEOrderHistory struct {
	OrderId  int    `json:"order_id"`
	Seller   string `json:"seller"`
	Buyer    string `json:"buyer"`
	Code     string `json:"code"`
	Quantity int    `json:"quantity"`
	Price    int    `json:"price"`
	SoldAt   string `json:"sold_at"`
}

type GEOrderCreated struct {
	Id         int    `json:"id"`
	CreatedAt  string `json:"created_at"`
	Code       string `json:"code"`
	Quantity   int    `json:"quantity"`
	Price      int    `json:"price"`
	TotalPrice int    `json:"total_price"`
	Tax        int    `json:"tax"`
}

type GETransaction struct {
	Id         string `json:"id"`
	Code       int    `json:"code"`
	Quantity   int    `json:"quantity"`
	Price      int    `json:"price"`
	TotalPrice int    `json:"total_price"`
}

type GEItem struct {
	Code     string `json:"code"`
	Quantity int    `json:"quantity"`
	Price    int    `json:"price,omitempty"`
}

type GEItemSchema struct {
	Code        string `json:"code"`
	Stock       int    `json:"stock"`
	SellPrice   int    `json:"sell_price"`
	BuyPrice    int    `json:"buy_price"`
	MaxQuantity int    `json:"max_quantity"`
}
