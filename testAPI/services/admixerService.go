package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testAPI/models"
)

type AdmixerServiceInterFace interface {
	GetPrice(target *models.GetPriceResponse, url string) error
}

type AdmixerService struct {
	servicePort string
}

func NewAdmixerService() *AdmixerService {

	return &AdmixerService{servicePort: "3333"}
}

func (service *AdmixerService) GetPrice(target *models.GetPriceResponse, url string) error {
	res, err := http.Get(url)

	if err != nil {
		fmt.Printf("error making http request: %s\n", err)
		return err
	}

	fmt.Printf("client: got response from %s!\n", url)
	fmt.Printf("client: status code: %d\n", res.StatusCode)

	return json.NewDecoder(res.Body).Decode(target)
}
