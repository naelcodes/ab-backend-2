package invoice

import (
	. "neema.co.za/rest/modules/payment/internal/api"
	"neema.co.za/rest/utils/middlewares"
	. "neema.co.za/rest/utils/payloads"
)

func GetModule() *Module {
	api := BuildApi()
	handleRoutes(api)
	module := Module(*api) //Module is an alias of Api
	return &module
}

func handleRoutes(api *Api) {
	api.Get("", api.GetAllPaymentsHandler)
	api.Get("/:id", api.GetPaymentHandler)
	api.Post("", middlewares.PayloadValidator(new(CreatePaymentPayload)), api.CreatePaymentHandler)
}
