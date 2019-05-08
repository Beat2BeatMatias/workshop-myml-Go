package external_api

//import (
//  "github.com/mercadolibre/workshop-shipping/src/api/domain/addresses"
//  "github.com/mercadolibre/workshop-shipping/src/api/domain/items"
//  "github.com/mercadolibre/workshop-shipping/src/api/domain/payments"
//)

type Order struct {
    ID          int      `json:"id"`
    DateCreated string   `json:"date_created"`
    OrderItems  []string `json:"order_items"`
    Address     string   `json:"address"`
    TotalAmount int      `json:"total_amount"`
    Status      string   `json:"status"`
    CurrencyID  string   `json:"currency_id"`
    Tags        []string `json:"tags"`
    Buyer       string   `json:"buyer"`
    Seller      string   `json:"seller"`
    Payments    []string `json:"payments"`
}
