package requests

type QuantityUnitText struct {
	QuantityUnit     string  `json:"QuantityUnit"`
	Language         string  `json:"Language"`
	QuantityUnitName *string `json:"QuantityUnitName"`
}
