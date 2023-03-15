package models

type GetMaxPriceRequest struct {
	RequestID  int    `json:"request_id,omitempty" schema:"RequestID"`
	UrlPackage []uint `json:"url_package" validate:"required" schema:"UrlPackage[]"`
	IP         string `json:"ip" schema:"IP"`
}
