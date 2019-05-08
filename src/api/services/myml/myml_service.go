package myml

import (
    "encoding/json"
    "fmt"
    "github.com/mercadolibre/go-meli-toolkit/goutils/apierrors"
    "github.com/mercadolibre/workshop-myml/src/api/domain/external_api"
    "github.com/mercadolibre/workshop-myml/src/api/domain/myml"
    "github.com/mercadolibre/workshop-myml/src/api/rest_client"
    "sync"
)

const (
    baseURL = "http://localhost:8081/%s/%v"
)

func GetOrderInformation(orderID int64) (*myml.FullOrder, apierrors.ApiError) {
    orderURL := fmt.Sprintf(baseURL, "orders", orderID)
    resp, apiErr := rest_client.Get(orderURL)
    if apiErr != nil {
        return nil, apiErr
    }

    var order external_api.Order
    if err := json.Unmarshal(resp.Bytes(), &order); err != nil {
        return nil, apierrors.NewInternalServerApiError(err.Error(), err)
    }

    var fullOrder myml.FullOrder
    fullOrder.ID = order.ID
    fullOrder.DateCreated = order.DateCreated

    c := make(chan *myml.FullOrder)
    defer close(c)

    var wg sync.WaitGroup
    goRoutinesCount := 3

    go func() {
        for i := 0; i < goRoutinesCount; i++ {
            currentOrder := <-c
            wg.Done()
            if currentOrder.Item != nil {
                fullOrder.Item = currentOrder.Item
                continue
            }

            if currentOrder.Payment != nil {
                fullOrder.Payment = currentOrder.Payment
                continue
            }

            if currentOrder.Address != nil {
                fullOrder.Address = currentOrder.Address
                continue
            }
        }
    }()

    wg.Add(goRoutinesCount)
    go getItem(order.OrderItems[0], c)
    go getPayment(order.Payments[0], c)
    go getAddress(order.Address, c)
    wg.Wait()

    return &fullOrder, nil
}

func getItem(itemID string, c chan *myml.FullOrder) {
    itemURL := fmt.Sprintf(baseURL, "items", itemID)
    resp, apiErr := rest_client.Get(itemURL)
    if apiErr != nil {
        c <- nil
        return
    }

    var item external_api.Item
    if err := json.Unmarshal(resp.Bytes(), &item); err != nil {
        c <- nil
        return
    }

    c <- &myml.FullOrder{
        Item: &item,
    }
}

func getPayment(paymentID string, c chan *myml.FullOrder) {
    paymentURL := fmt.Sprintf(baseURL, "payments", paymentID)
    resp, apiErr := rest_client.Get(paymentURL)
    if apiErr != nil {
        c <- nil
        return
    }
    var payment external_api.Payment
    if err := json.Unmarshal(resp.Bytes(), &payment); err != nil {
        c <- nil
        return
    }
    c <- &myml.FullOrder{
        Payment: &payment,
    }
}

func getAddress(addressID string, c chan *myml.FullOrder) {
    addressURL := fmt.Sprintf(baseURL, "addresses", addressID)
    resp, apiErr := rest_client.Get(addressURL)
    if apiErr != nil {
        c <- nil
        return
    }
    var address external_api.Address
    if err := json.Unmarshal(resp.Bytes(), &address); err != nil {
        c <- nil
        return
    }
    c <- &myml.FullOrder{
        Address: &address,
    }
}
