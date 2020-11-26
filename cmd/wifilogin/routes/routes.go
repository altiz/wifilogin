package routes

import (
	_ "wifilogin/cmd/wifilogin/docs" // docs is generated by Swag CLI, you have to import it.
	"wifilogin/cmd/wifilogin/handlers"

	"github.com/gin-gonic/gin" // swagger embed files
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitializeRoutes(router *gin.Engine) {

	url := ginSwagger.URL("http://localhost:5000/swagger/doc.json") // The url pointing to API definition
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	api := router.Group("billing/api")
	version1 := api.Group("/v1")

	version1.POST("/test", handlers.TestJSON)
	version1.GET("/", handlers.IndexPage)
	version1.POST("/home", handlers.HomeHandlers)
	version1.POST("/set-msisdn", handlers.Setmsisdn)
}
