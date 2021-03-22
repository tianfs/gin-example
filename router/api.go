package router

import (
    "gin-example/controller"
    _ "gin-example/docs"
    "gin-example/middleware/recovery"
    "gin-example/util/uploadTool"
    "github.com/gin-gonic/gin"
    swaggerFiles "github.com/swaggo/files"
    ginSwagger "github.com/swaggo/gin-swagger"
    "net/http"
)

func SetRouter() *gin.Engine {
    r := gin.Default()
    //增加请求中panic拦截防止进程退出
    r.Use(recovery.Recovery())
    r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

    //获取控制器
    auth := controller.NewAuth()
    saleOrder := controller.NewSaleOrder()
    upload := controller.NewUpload()

    authG := r.Group("/auth")
    {
        authG.GET("/getToken", auth.GetToken)

    }

    saleOrderG := r.Group("/saleOrder")
    //saleOrderG.Use(jwt.JWT());
    {
        saleOrderG.GET("/list", saleOrder.List)
        saleOrderG.POST("/create", saleOrder.Create)
        saleOrderG.POST("/update", saleOrder.Update)
        saleOrderG.POST("/delete", saleOrder.Delete)
        saleOrderG.POST("/kafkaAsyncProducer", saleOrder.KafkaAsyncProducer)
        saleOrderG.POST("/kafkaProducer", saleOrder.KafkaProducer)

    }

    uploadG := r.Group("/upload")
    {
        uploadG.StaticFS("/images", http.Dir(uploadTool.GetImageFullPath()))
        uploadG.POST("/image", upload.Image)
    }

    return r
}
