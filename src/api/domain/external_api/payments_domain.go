package external_api

type Payment struct {
    ID                int    `json:"id"`
    OrderID           int    `json:"order_id"`
    PaymentMethodID   string `json:"payment_method_id"`
    CurrencyID        string `json:"currency_id"`
    Status            string `json:"status"`
    TransactionAmount int    `json:"transaction_amount"`
}
