package router

import (
    "gin-example/middleware/recovery"
    "github.com/gin-gonic/gin"
    swaggerFiles "github.com/swaggo/files"
    ginSwagger "github.com/swaggo/gin-swagger"
    "net/http"
    "gin-example/controller"
    _ "gin-example/docs"
    "gin-example/util/upload"
)

func SetRouter() *gin.Engine {
    r := gin.Default()
    //增加请求中panic拦截防止进程退出
    r.Use(recovery.Recovery())
    r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

    authG := r.Group("/auth")
    {
        var authC controller.Auth

        authG.GET("/getToken", authC.GetToken)

    }

    saleOrderG := r.Group("/saleOrder")
    //saleOrderG.Use(jwt.JWT());
    {
        var saleOrderC controller.SaleOrder

        saleOrderG.GET("/list", saleOrderC.List)
        saleOrderG.POST("/create", saleOrderC.Create)
        saleOrderG.POST("/update", saleOrderC.Update)
        saleOrderG.POST("/delete", saleOrderC.Delete)
        saleOrderG.POST("/kafkaAsyncProducer", saleOrderC.KafkaAsyncProducer)
        saleOrderG.POST("/kafkaProducer", saleOrderC.KafkaProducer)

    }

    uploadG := r.Group("/upload")
    {
        var uploadC controller.Upload
        uploadG.StaticFS("/images", http.Dir(upload.GetImageFullPath()))
        uploadG.POST("/image", uploadC.Image)
    }



    return r
}
