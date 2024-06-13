package server

import (
	"book-store-management-backend/component/appctx"
	"book-store-management-backend/middleware"
	"book-store-management-backend/server/routerfacade"

	"book-store-management-backend/docs"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type ConcreteServerBuilder struct {
	router *gin.Engine
	appCtx appctx.AppContext
}

func NewServerBuilder(appCtx appctx.AppContext) *ConcreteServerBuilder {
	router := gin.Default()
	return &ConcreteServerBuilder{
		router: router,
		appCtx: appCtx,
	}
}

func (b *ConcreteServerBuilder) SetMiddlewares() ServerBuilder {
	b.router.Use(middleware.CORSMiddleware())
	b.router.Use(middleware.Recover(b.appCtx))
	return b
}

func (b *ConcreteServerBuilder) SetRoutes() ServerBuilder {
	docs.SwaggerInfo.BasePath = "/v1"
	v1 := b.router.Group("/v1")
	v1.GET("/docs/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	routerFacade := routerfacade.NewRouteFacade(v1, b.appCtx)
	{
		routerFacade.SetUpload()
		routerFacade.SetAuthor()
		routerFacade.SetCategory()
		routerFacade.SetBookTitle()
		routerFacade.SetBook()
		routerFacade.SetPublisher()
		routerFacade.SetInvoice()
		routerFacade.SetImportNote()
		routerFacade.SetInventoryCheckNote()
		routerFacade.SetSupplier()
		routerFacade.SetCustomer()
		routerFacade.SetRole()
		routerFacade.SetFeature()
		routerFacade.SetUser()
		routerFacade.SetShopGeneral()
		routerFacade.SetDashboard()
		routerFacade.SetReport()
	}
	return b
}

func (b *ConcreteServerBuilder) Build() *Server {
	return &Server{
		router: b.router,
		appCtx: b.appCtx,
	}
}
