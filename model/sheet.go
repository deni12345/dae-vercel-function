package model

import (
	"time"

	"cloud.google.com/go/firestore"
)

type Sheet struct {
	ID        string                 `firestore:"-" json:"id"` // Firestore ID, not stored in Firestore
	Title     string                 `firestore:"title" json:"title"`
	EndAt     time.Time              `firestore:"endAt" json:"endAt"` // End time for the sheet
	Menu      Menu                   `firestore:"menu" json:"menu"`   // Changed to match original design
	Status    string                 `firestore:"status" json:"status"`
	CreatedBy *firestore.DocumentRef `firestore:"createdBy" json:"createdBy"` // Reference to the user who created the sheet
	CreatedAt time.Time              `firestore:"createdAt" json:"createdAt"`
}

func (s *Sheet) SetID(id string) {
	s.ID = id
}

func NewSheet(title, status string, menu Menu, createdBy *firestore.DocumentRef) Sheet {
	return Sheet{
		Title:     title,
		CreatedBy: createdBy,
		CreatedAt: time.Now(),
		Status:    status,
		Menu:      menu,
	}
}

type Menu struct {
	Name      string     `firestore:"name"`
	MenuItems []MenuItem `firestore:"menu"` // Embedded items, not references
}

type MenuItem struct {
	ItemID    string  `firestore:"itemId"`
	Name      string  `firestore:"name"`
	Price     float64 `firestore:"price"` // float64 is more common for currency
	Category  string  `firestore:"category"`
	Available bool    `firestore:"available"`
}

func NewMenu(name string, items []MenuItem) *Menu {
	return &Menu{
		Name:      name,
		MenuItems: items,
	}
}

func NewMenuItem(itemID, name, category string, price float64, available bool) *MenuItem {
	return &MenuItem{
		ItemID:    itemID,
		Name:      name,
		Category:  category,
		Price:     price,
		Available: available,
	}
}

// Order represents an individual order within a sheet
type Order struct {
	UserID      string      `firestore:"userId"`   // Reference to users collection
	Username    string      `firestore:"username"` // Denormalized for display
	OrderItems  []OrderItem `firestore:"items"`
	Description string      `firestore:"description,omitempty"`
	CreatedAt   time.Time   `firestore:"createdAt"`
	Status      string      `firestore:"status"` // "pending" | "fulfilled" | "cancelled"
}

type OrderItem struct {
	ItemID      string    `firestore:"itemId"` // Reference to menu item
	Name        string    `firestore:"name"`   // Denormalized
	Quantity    int       `firestore:"quantity"`
	Toppings    []Topping `firestore:"toppings,omitempty"`
	Description string    `firestore:"description,omitempty"`
}

type Topping struct {
	Name  string  `firestore:"name"`
	Price float64 `firestore:"price,omitempty"`
}

func NewOrder(userID, username, description, status string, items []OrderItem) *Order {
	return &Order{
		UserID:      userID,
		Username:    username,
		OrderItems:  items,
		Description: description,
		Status:      status,
		CreatedAt:   time.Now(),
	}
}
