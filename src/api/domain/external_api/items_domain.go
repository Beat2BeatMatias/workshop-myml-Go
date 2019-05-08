package external_api

type Item struct {
    ID         int      `json:"id"`
    SiteID     string   `json:"site_id"`
    Title      string   `json:"title"`
    CategoryID string   `json:"category_id"`
    Price      int      `json:"price"`
    CurrencyID string   `json:"currency_id"`
    Quantity   int      `json:"quantity"`
    Condition  string   `json:"condition"`
    Pictures   []string `json:"pictures"`
}
