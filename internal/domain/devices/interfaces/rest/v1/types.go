package v1

import "github.com/mendelgusmao/device-manager/internal/domain/devices/dtos"

type deviceCreateRequest struct {
	Body *dtos.CreateDevice
}

type deviceUpdateRequest struct {
	Identifier string `path:"identifier" maxLength:"36" example:"f3c0cc81-4cac-43fd-bb5a-77f79a111413" doc:"Device identifier"`
	Body       dtos.UpdateDevice
}

type deviceIdentifier struct {
	Identifier string `path:"identifier" maxLength:"36" example:"f3c0cc81-4cac-43fd-bb5a-77f79a111413" doc:"Device identifier"`
}

type deviceSearch struct {
	BrandName  string `query:"brandName" maxLength:"20" example:"Awesome Brand" doc:"Brand name"`
	DeviceName string `query:"deviceName" maxLength:"20" example:"Awesome Device" doc:"Device name"`
}

type deviceResponse struct {
	Body *dtos.Device
}

type devicesResponse struct {
	Body []dtos.Device
}
