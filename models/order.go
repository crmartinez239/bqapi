package models

type OrderModel struct {
	ID    uint   `json:"id" gorm:"primary_key"`
	Type  string `json:"type"`
	Name  string `json:"name"`
	When  string `json:"when"`
	IsNew string `json:"isnew"`
}

type CreateOrderModel struct {
	Type string `json:"type" binding:"required"`
	Name string `json:"name" binding:"required"`
}

type UpdateOrderModel struct {
	Type string `json:"type"`
	Name string `json:"name"`
}
