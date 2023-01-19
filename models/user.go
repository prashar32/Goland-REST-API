package models

type User struct {
	UserId      int    `json:"userId" gorm:"primaryKey"`
	Name        string `json:"name"`
	UserEmail   string `json:"userEmail"`
	UserContact string `json:"userContact"`
	ShopVisited int    `json:"shopVisited"`
}
