package models

import "time"

type Item struct {
	ID          uint   		 `gorm:"primaryKey" json:"id"`
	UserID			uint			 `json:"user_id"`
	User				User			 
	Name        string 		 `json:"name"`
	Quantity    int    		 `json:"quantity"`
	IsPurchased bool   		 `json:"is_purchased"`
	Note        string 		 `json:"note"`
	CreatedAt   time.Time  `json:"created_at"`
	PurchasedAt *time.Time `json:"purchased_at"`
}