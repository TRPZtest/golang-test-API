package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testAPI/models"
)

type getMaxPriceTaskPOST struct{}

func NewGetMaxPriceHandlerPOST() *ApiHandler {
	return NewApiHandler(&getMaxPriceTaskPOST{})
}

func (getMaxPriceTaskPOST) Execute(w http.ResponseWriter, data any) {
	GetMaxPrice(w, data)
}

func (handler *getMaxPriceTaskPOST) GetPayload(r *http.Request) (any, error) {
	var reqData models.GetMaxPriceRequest

	reqJson, _ := ioutil.ReadAll(r.Body)

	error := json.Unmarshal(reqJson, &reqData)

	return reqData, error
}
