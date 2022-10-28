package domain

type Item struct {
	ID          string `json:"id" gorm:"primaryKey"`
	ItemCode    string `json:"item_code"`
	Description string `json:"description"`
	Quantity    uint   `json:"quantity"`
	OrderID     uint   `json:"order_id"`
	Order       Order
}
