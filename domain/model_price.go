package domain


//Price represents the info about price and purchase
type Price struct {
    Price    int32  `json:"price"`
    PurchaseID string `json:"purchaseId"`
}