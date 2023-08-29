package routes

import "github.com/gin-gonic/gin"

func InvoiceRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/invoices", controller.GetInvoice())
	incomingRoutes.GET("/invoices/:invoice_id", controller.GETInvoice())
	incomingRoutes.POST("/invoice", controller.CreateInvoice())
	incomingRoutes.PATCH("/invoices/:invoices_id", controller.UpdateInvoice())
}
