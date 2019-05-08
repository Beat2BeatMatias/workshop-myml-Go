package external_api

type Address struct {
    ID           int    `json:"id"`
    StreetName   string `json:"street_name"`
    StreetNumber int    `json:"street_number"`
    City         string `json:"city"`
}
