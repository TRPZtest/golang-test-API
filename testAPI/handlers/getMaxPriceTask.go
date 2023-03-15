package handlers

import (
	"log"
	"net"
	"net/http"
	dal "testAPI/dal"
	"testAPI/models"
	response "testAPI/rest"
	"testAPI/services"
)

func GetMaxPrice(w http.ResponseWriter, data any) {

	req := data.(models.GetMaxPriceRequest)

	if len(req.UrlPackage) == 0 || !_IPAdressIsValid(req.IP) {
		response.NoContentResponse(w)
		return
	}

	service := services.NewAdmixerService()

	repository, err := dal.NewUrlRepository()

	if err != nil {
		response.InternalServerErrorResponse(w, err)
		return
	}

	urlPackages, err := repository.GetUrlPackages(req.UrlPackage)

	if len(urlPackages) != len(req.UrlPackage) {
		response.NoContentResponse(w)
		return
	}
	var getPriceResponse = models.GetPriceResponse{}

	if err != nil {
		response.InternalServerErrorResponse(w, err)
		return
	}

	maxPrice := 0.0
	for i := 0; i < len(urlPackages); i++ {
		err = service.GetPrice(&getPriceResponse, urlPackages[i].Url)

		if getPriceResponse.Price > maxPrice {
			maxPrice = getPriceResponse.Price
		}

		log.Printf("Price: %f", getPriceResponse.Price)
	}

	response.SuccessJsonResponse(w, models.GetPriceResponse{Price: maxPrice})
}

func _IPAdressIsValid(ip string) bool {
	if net.ParseIP(ip) == nil {
		log.Printf("IP Address: %s - Invalid\n", ip)
		return false
	} else {
		log.Printf("IP Address: %s - Valid\n", ip)
		return true
	}
}
