package models

type Order struct {
	ID      int32   `json:"id"`
	Phone   string  `json:"phone"`
	ItemIDs []int32 `json:"item_ids"`
}
