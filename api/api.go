package api

import (
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"

	_ "app/api/docs"
	"app/api/handler"
	"app/storage"
)

func NewApi(r *gin.Engine, storage storage.StorageI) {

	handlerV1 := handler.NewHandler(storage)

	r.POST("/investor", handlerV1.CreateInvestor)
	r.GET("/investor/:id", handlerV1.GetByIdInvestor)
	r.GET("/investor", handlerV1.GetListInvestor)
	r.GET("/investor-share", handlerV1.GetListInvestorFoyda)
	r.DELETE("/investor/:id", handlerV1.DeleteInvestor)
	r.PUT("/investor/:id", handlerV1.UpdateInvestor)

	r.POST("/car", handlerV1.CreateCar)
	r.GET("/car/:id", handlerV1.GetByIdCar)
	r.GET("/car", handlerV1.GetListCar)
	r.DELETE("/car/:id", handlerV1.DeleteCar)
	r.PUT("/car/:id", handlerV1.UpdateCar)

	r.POST("/client", handlerV1.CreateClient)
	r.GET("/client/:id", handlerV1.GetByIdClient)
	r.GET("/client", handlerV1.GetListClient)
	r.DELETE("/client/:id", handlerV1.DeleteClient)
	r.PUT("/client/:id", handlerV1.UpdateClient)

	r.POST("/order", handlerV1.CreateOrder)
	r.GET("/order/:id", handlerV1.GetByIdOrder)
	r.GET("/order", handlerV1.GetListOrder)
	r.DELETE("/order/:id", handlerV1.DeleteOrder)
	r.PUT("/order/:id", handlerV1.UpdateOrder)
	r.PUT("/orderstatus/:id", handlerV1.UpdateOrderStatus)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

}
