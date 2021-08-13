package router

import (
	"github.com/gin-gonic/gin"
	"microshop/application"
	"microshop/router/middleware"
	"net/http"
)

func Load(g *gin.Engine) *gin.Engine {
	g.Use(middleware.NoCache)
	g.Use(middleware.Options)
	g.Use(middleware.Secure)

	g.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"error":"invalid route"})
	})
	authGroup := g.Group("/api")
	authGroup.Use(middleware.StrictTokenHandler())
	{
		// address
		authGroup.GET("address/list", application.ListAddress)
		authGroup.GET("address/detail", application.DetailAddress)
		authGroup.POST("address/save", application.SaveAddress)
		authGroup.DELETE("address/delete", application.DeleteAddress)

		// cart
		authGroup.GET("cart/index", application.CartList)
		authGroup.POST("cart/add", application.CartAdd)
		authGroup.GET("cart/cartcount", application.CartCount)
		authGroup.POST("cart/update", application.CartUpdate)
		authGroup.POST("cart/delete", application.CartDelete)
		authGroup.POST("cart/checked", application.CartChecked)
		authGroup.GET("cart/checkout", application.CartCheckout)

		// order
		authGroup.POST("order/submit", application.OrderSubmit)
		authGroup.GET("order/list", application.OrderList)
		authGroup.GET("order/detail", application.OrderDetail)
		authGroup.POST("order/updatestatus", application.OrderStatusUpdate)


	}
	noAuthGroup := g.Group("/api")
	{
		noAuthGroup.GET("list", application.CartList)
		noAuthGroup.GET("index/index", application.Home)
		noAuthGroup.GET("catalog/index", application.CatalogList)
		noAuthGroup.GET("goods/list", application.GoodsList)
		noAuthGroup.GET("goods/detail", application.GoodsDetail)
		noAuthGroup.POST("sms/send", application.SendSms)
		noAuthGroup.POST("auth/loginByWeixin", application.Login)
		noAuthGroup.POST("alipay/pay", application.Pay)
		noAuthGroup.POST("alipay/notify_url", application.PayNotify)
	}
	return g
}
