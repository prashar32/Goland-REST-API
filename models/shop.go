package models

type Shop struct {
	ShopId       int    `json:"shopId" gorm:"primaryKey"`
	ShopNo       int    `json:"shopNo" gorm:"autoIncrement:true"`
	ShopName     string `json:"shopName"`
	FloorNo      int    `json:"floorNo"`
	ShopCategory string `json:"shopCategory"`
	OpeningTime  string `json:"openingTime"`
	ClosingTime  string `json:"closingTime"`
}
