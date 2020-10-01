package domain


//Price represents the info about price and purchase
type Price struct {
    Price    int32  `json:"lastPrice"`
    PurchaseID string `json:"lastPriceId"`
}