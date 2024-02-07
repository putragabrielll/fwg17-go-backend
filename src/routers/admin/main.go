package adminRouters

import "github.com/gin-gonic/gin"

func AdminRouter(r *gin.RouterGroup){
	usersRouter(r.Group("/users"))
	productsRouter(r.Group("/products"))
	promoRouters(r.Group("/promo"))
	psRouters(r.Group("/product-size"))
	pvRouters(r.Group("/product-variant"))
	categoriesRouters(r.Group("/categories"))
	tagRouters(r.Group("/tags"))
	PTRouters(r.Group("/product-tags"))
	PCRouters(r.Group("/product-categories"))
}