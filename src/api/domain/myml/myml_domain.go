package myml

import (
    "github.com/mercadolibre/workshop-myml/src/api/domain/external_api"
)

type FullOrder struct {
    ID          int                   `json:"id"`
    DateCreated string                `json:"date_created"`
    Item        *external_api.Item    `json:"item"`
    Payment     *external_api.Payment `json:"payment"`
    Address     *external_api.Address `json:"address"`
}
