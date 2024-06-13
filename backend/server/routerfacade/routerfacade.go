package routerfacade

import (
	"book-store-management-backend/component/appctx"
	"book-store-management-backend/module/author/authortransport"
	"book-store-management-backend/module/book/booktransport"
	"book-store-management-backend/module/booktitle/booktitletransport"
	"book-store-management-backend/module/category/categorytransport"
	"book-store-management-backend/module/customer/customertransport/gincustomer"
	"book-store-management-backend/module/dashboard/dashboardtransport/gindashboard"
	"book-store-management-backend/module/feature/featuretransport/ginfeature"
	"book-store-management-backend/module/importnote/importnotetransport/ginimportnote"
	"book-store-management-backend/module/inventorychecknote/inventorychecknotetransport/gininventorychecknote"
	"book-store-management-backend/module/invoice/invoicetransport/gininvoice"
	"book-store-management-backend/module/publisher/publishertransport"
	"book-store-management-backend/module/role/roletransport/ginrole"
	"book-store-management-backend/module/salereport/salereporttransport/ginsalereport"
	"book-store-management-backend/module/shopgeneral/shopgeneraltransport/ginshopgeneral"
	ginstockreports "book-store-management-backend/module/stockreport/stockreporttransport/ginstockreport"
	"book-store-management-backend/module/supplier/suppliertransport/ginsupplier"
	"book-store-management-backend/module/supplierdebtreport/supplierdebtreporttransport/ginsupplierdebtreport"
	"book-store-management-backend/module/upload/uploadtransport"
	"book-store-management-backend/module/user/usertransport/ginuser"

	"github.com/gin-gonic/gin"
)

type RouteFacade struct {
	routerGroup *gin.RouterGroup
	appCtx      appctx.AppContext
}

func NewRouteFacade(routerGroup *gin.RouterGroup, appCtx appctx.AppContext) *RouteFacade {
	return &RouteFacade{
		routerGroup: routerGroup,
		appCtx:      appCtx,
	}
}

func (rf *RouteFacade) SetUpload() {
	uploadtransport.SetupRoutes(rf.routerGroup, rf.appCtx)
}
func (rf *RouteFacade) SetAuthor() {
	authortransport.SetupRoutes(rf.routerGroup, rf.appCtx)
}
func (rf *RouteFacade) SetCategory() {
	categorytransport.SetupRoutes(rf.routerGroup, rf.appCtx)
}
func (rf *RouteFacade) SetBookTitle() {
	booktitletransport.SetupRoutes(rf.routerGroup, rf.appCtx)
}
func (rf *RouteFacade) SetBook() {
	booktransport.SetupRoutes(rf.routerGroup, rf.appCtx)
}
func (rf *RouteFacade) SetPublisher() {
	publishertransport.SetupRoutes(rf.routerGroup, rf.appCtx)
}
func (rf *RouteFacade) SetInvoice() {
	gininvoice.SetupRoutes(rf.routerGroup, rf.appCtx)
}
func (rf *RouteFacade) SetImportNote() {
	ginimportnote.SetupRoutes(rf.routerGroup, rf.appCtx)
}
func (rf *RouteFacade) SetInventoryCheckNote() {
	gininventorychecknote.SetupRoutes(rf.routerGroup, rf.appCtx)
}
func (rf *RouteFacade) SetSupplier() {
	ginsupplier.SetupRoutes(rf.routerGroup, rf.appCtx)
}
func (rf *RouteFacade) SetCustomer() {
	gincustomer.SetupRoutes(rf.routerGroup, rf.appCtx)
}
func (rf *RouteFacade) SetRole() {
	ginrole.SetupRoutes(rf.routerGroup, rf.appCtx)
}
func (rf *RouteFacade) SetFeature() {
	ginfeature.SetupRoutes(rf.routerGroup, rf.appCtx)
}
func (rf *RouteFacade) SetUser() {
	ginuser.SetupRoutes(rf.routerGroup, rf.appCtx)
}
func (rf *RouteFacade) SetShopGeneral() {
	ginshopgeneral.SetupRoutes(rf.routerGroup, rf.appCtx)
}
func (rf *RouteFacade) SetDashboard() {
	gindashboard.SetupRoutes(rf.routerGroup, rf.appCtx)
}
func (rf *RouteFacade) SetReport() {
	report := rf.routerGroup.Group("/reports")
	{
		ginstockreports.SetupRoutes(report, rf.appCtx)
		ginsupplierdebtreport.SetupRoutes(report, rf.appCtx)
		ginsalereport.SetupRoutes(report, rf.appCtx)
	}
}
