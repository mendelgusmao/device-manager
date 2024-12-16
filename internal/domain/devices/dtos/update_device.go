package dtos

type UpdateDevice struct {
	BrandName  *string `json:"brandName,omitempty" doc:"Brand name" minLength:"3" maxLength:"20" required:"false"`
	DeviceName *string `json:"deviceName,omitempty" doc:"Device name" minLength:"1" maxLength:"30" required:"false"`
}
