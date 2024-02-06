package adminRouters

import "github.com/gin-gonic/gin"

func AdminRouter(r *gin.RouterGroup){
	usersRouter(r.Group("/users"))
	productsRouter(r.Group("/products"))
	promoRouters(r.Group("/promo"))
	psRouters(r.Group("/product-size"))
}