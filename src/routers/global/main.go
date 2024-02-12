package globalRouters

import "github.com/gin-gonic/gin"


func GlobalRouter(r *gin.RouterGroup){
	productsRouter(r.Group("/products"))
	sizeRouters(r.Group("/product-size"))
	variantRouters(r.Group("/product-variant"))
}