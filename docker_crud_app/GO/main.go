package main

import (
	"amazon-backend/config"
	sellerauth "amazon-backend/handlers/authentication/seller_auth"
	userauth "amazon-backend/handlers/authentication/user_auth"
	sellerhandler "amazon-backend/handlers/seller_handler"
	userhandler "amazon-backend/handlers/user_handler"
	"amazon-backend/middleware"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	config.Connect()
	router := gin.Default()
	router.Use(cors.Default())
	initRoutes(router)
	router.Run(":8080")
}

func initRoutes(r *gin.Engine) {
	r.POST("/api/user/login", userauth.Login)
	r.POST("/api/user/register", userauth.Register)
	r.POST("/api/seller/login", sellerauth.Login)
	r.POST("/api/seller/register", sellerauth.Register)
	r.POST("/api/seller/addproduct", middleware.RequireSellerAuth, sellerhandler.AddProduct)
	r.GET("/api/seller/getproduct", middleware.RequireSellerAuth, sellerhandler.Get_Product)
	r.GET("/api/getallproducts", userhandler.Search_Product)
	r.DELETE("/api/seller/delete", middleware.RequireSellerAuth, sellerhandler.RemoveProduct)
	r.PUT("/api/seller/update", middleware.RequireSellerAuth, sellerhandler.Update_Product)
	r.POST("/api/user/addcart", userhandler.Add_Item_In_Cart)
	r.DELETE("/api/user/removecart", userhandler.Remove_Item_In_Cart)
	r.GET("/api/user/viewcart", userhandler.View_Item_In_Cart)
	r.POST("/api/user/placeorder", middleware.RequireUserAuth, userhandler.Place_Order)
	r.DELETE("/api/user/cancelorder", middleware.RequireUserAuth, userhandler.Cancel_Order)
	r.GET("/api/user/vieworder", middleware.RequireUserAuth, userhandler.View_Purchased_Items)
}
