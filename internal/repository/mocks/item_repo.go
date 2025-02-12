package mocks

import (
	"avito-shop/internal/entity"
	"avito-shop/internal/repository"
	"errors"
)

/*
MockItemRepo implements the ItemRepository interface using an in-memory map.
It simulates a fixed set of items.
*/
type MockItemRepo struct {
	Items map[string]entity.Item
}

/*
NewMockItemRepo creates and returns a new instance of MockItemRepo with a predefined set of items.
*/
func NewMockItemRepo() repository.ItemRepository {
	return &MockItemRepo{
		Items: map[string]entity.Item{
			"t-shirt":    {Name: "t-shirt", Price: 80},
			"cup":        {Name: "cup", Price: 20},
			"book":       {Name: "book", Price: 50},
			"pen":        {Name: "pen", Price: 10},
			"powerbank":  {Name: "powerbank", Price: 200},
			"hoody":      {Name: "hoody", Price: 300},
			"umbrella":   {Name: "umbrella", Price: 200},
			"socks":      {Name: "socks", Price: 10},
			"wallet":     {Name: "wallet", Price: 50},
			"pink-hoody": {Name: "pink-hoody", Price: 500},
		},
	}
}

/*
GetItem retrieves an item by its name.
It returns an error if the item is not found.
*/
func (f *MockItemRepo) GetItem(name string) (*entity.Item, error) {
	item, ok := f.Items[name]
	if !ok {
		return nil, errors.New("item not found")
	}
	return &item, nil
}

/*
GetAllItems returns all items stored in the repository.
*/
func (f *MockItemRepo) GetAllItems() ([]entity.Item, error) {
	var items []entity.Item
	for _, item := range f.Items {
		items = append(items, item)
	}
	return items, nil
}
