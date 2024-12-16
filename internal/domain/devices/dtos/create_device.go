package dtos

type CreateDevice struct {
	BrandName  string `json:"brandName" doc:"Brand name" minLength:"3" maxLength:"20"`
	DeviceName string `json:"deviceName" doc:"Device name" minLength:"1" maxLength:"30"`
}
