package usecase

import (
	"sync"

	"github.com/andersonlira/item-api/domain"
	"github.com/andersonlira/item-api/gateway/txtdb"
)


type Option int

const (
	Low Option = iota
	High
	Last
)

var m = &sync.Mutex{}

func LastPrice(ID string, price domain.Price,op Option) error {
	m.Lock()
	item, err := txtdb.GetItemByID(ID)
	if err != nil {
		return err
	}

	switch op {
	case Low:
		item.LowestPrice = price.Price
		item.LowestPurchaseID = price.PurchaseID
	case High:
		item.HighestPrice = price.Price
		item.HighestPurchaseID = price.PurchaseID
	default: 
		item.LastPrice = price.Price
		item.LastPurchaseID = price.PurchaseID

	}

	txtdb.UpdateItem(ID, item)
	m.Unlock()
	return nil
}