package mock

import (
	"context"
	"errors"
	"github.com/Daerys/avito-shop/internal/entity"
)

type ItemRepository struct {
	items     map[string]*entity.Item
	inventory map[string]map[string]*entity.InventoryItem
}

func NewItemRepository() *ItemRepository {
	return &ItemRepository{
		items: map[string]*entity.Item{
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
		inventory: make(map[string]map[string]*entity.InventoryItem),
	}
}

func (m *ItemRepository) GetItem(ctx context.Context, item *entity.Item) error {
	savedItem, exists := m.items[item.Name]
	if !exists {
		return errors.New("item not found")
	}
	*item = *savedItem
	return nil
}

func (m *ItemRepository) AddItem(ctx context.Context, user *entity.User, item *entity.Item) error {
	if m.inventory[user.Username] == nil {
		m.inventory[user.Username] = make(map[string]*entity.InventoryItem)
	}

	if invItem, exists := m.inventory[user.Username][item.Name]; exists {
		invItem.Quantity++
	} else {
		m.inventory[user.Username][item.Name] = &entity.InventoryItem{
			UserID:   int(user.ID),
			ItemType: item.Name,
			Quantity: 1,
		}
	}
	return nil
}

func (m *ItemRepository) GetInventory(ctx context.Context, user *entity.User) ([]entity.InventoryItem, error) {
	var inventory []entity.InventoryItem
	for _, item := range m.inventory[user.Username] {
		inventory = append(inventory, *item)
	}
	return inventory, nil
}
