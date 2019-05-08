package rest_client

import (
    "errors"
    "github.com/mercadolibre/go-meli-toolkit/goutils/apierrors"
    "github.com/mercadolibre/go-meli-toolkit/restful/rest"
)

const (
    errRestClient = "error executing rest api call"
)

func Get(url string) (*rest.Response, apierrors.ApiError) {
    resp := rest.Get(url)
    if resp == nil {
        return nil, apierrors.NewInternalServerApiError(
            errRestClient,
            errors.New("nil response"))
    }
    return resp, nil
}
