package domain

import (
    "time"
)

//Item represents a model
type Item struct {
    ID string `json:"id"`
    CreatedAt time.Time `json:"dateCreated,omitempty"`
    UpdatedAt time.Time `json:"dateUpdated,omitempty"`
	Initial      string `json:"initial"`
	Name         string `json:"name"`
    LowestPrice  int32  `json:"lowestPrice"`
    LowestPurchaseID string `json:"lowestPurchaseId"`
    HighestPrice int32  `json:"highestPrice"`
    HighestPurchaseID string `json:"highestPurchaseId"`
    LastPrice    int32  `json:"lastPrice"`
    LastPurchaseID string `json:"lastPriceId"`
	GroupID        string `json:"groupId"`
}