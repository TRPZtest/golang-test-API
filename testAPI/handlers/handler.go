package handlers

import (
	"log"
	"net/http"
	response "testAPI/rest"
)

type ApiHandler struct {
	service TaskInterface
}

type TaskInterface interface {
	Execute(w http.ResponseWriter, data any)
	GetPayload(r *http.Request) (any, error)
}

func NewApiHandler(service TaskInterface) *ApiHandler {
	return &ApiHandler{
		service: service,
	}
}

func (handler *ApiHandler) Execute(w http.ResponseWriter, r *http.Request) {
	log.Println("******** NEW REQUEST ********")
	data, error := handler.service.GetPayload(r)

	if error != nil {
		response.BadRequestResponse(w, error)
		return
	}

	handler.service.Execute(w, data)
}
