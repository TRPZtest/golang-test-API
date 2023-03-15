package handlers

import (
	"net/http"
	"testAPI/models"

	"github.com/gorilla/schema"
)

type getMaxPriceTaskGET struct{}

func NewGetMaxPriceHandlerGET() *ApiHandler {
	return NewApiHandler(&getMaxPriceTaskGET{})
}

func (getMaxPriceTaskGET) GetPayload(r *http.Request) (any, error) {

	reqData := &models.GetMaxPriceRequest{}

	d := schema.NewDecoder()

	err := d.Decode(reqData, r.URL.Query())

	return *reqData, err
}

func (getMaxPriceTaskGET) Execute(w http.ResponseWriter, data any) {
	GetMaxPrice(w, data)
}
