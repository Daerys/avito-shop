package entity

import "sync"

/*
User represents an employee of the shop.
*/
type User struct {
	ID           int
	Username     string
	PasswordHash string
	Coins        int
	Inventory    *sync.Map
}
