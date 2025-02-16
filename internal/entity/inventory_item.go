package entity

type InventoryItem struct {
	UserID   int    `db:"user_id"`
	ItemType string `db:"item_type"`
	Quantity int    `db:"quantity"`
}
