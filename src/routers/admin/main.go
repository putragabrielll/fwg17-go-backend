package adminRouters

import (
	"github.com/gin-gonic/gin"
	"github.com/putragabrielll/go-backend/src/middlewares"
)

func AdminRouter(r *gin.RouterGroup) {
	authMiddleware, _ := middlewares.Auth()
	r.Use(authMiddleware.MiddlewareFunc())
	
	usersRouter(r.Group("/users"))
	productsRouter(r.Group("/products"))
	promoRouters(r.Group("/promo"))
	psRouters(r.Group("/product-size"))
	pvRouters(r.Group("/product-variant"))
	categoriesRouters(r.Group("/categories"))
	tagRouters(r.Group("/tags"))
	PTRouters(r.Group("/product-tags"))
	PCRouters(r.Group("/product-categories"))
	PRRouters(r.Group("/product-ratings"))
	ordersRouters(r.Group("/orders"))
	ordersDetailsRouters(r.Group("/order-details"))
}
