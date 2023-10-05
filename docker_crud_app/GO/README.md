## **🔥 Introduction**
This documentation provides an overview of the APIs available for a shopping app. These APIs allow users and sellers to interact with the shopping platform by performing actions such as user registration, login, product management, cart management, and order management etc.

## **💥 APIs**

```go
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
```
## **🛠️ Local Development** :

1. install all go dependencies
   ```shell
   $ go get ./...
   ```
2. Start Database on :5432
   ```shell
   docker compose up
   ```
3. Start the server on :8080
    ```shell
    go run main.go
    ```


**Note : You need to restart backend server after every change in any .go file.**
