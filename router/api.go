package router

import (
    ctl "gin-example/controller"
    _ "gin-example/docs"
    "gin-example/middleware/recovery"
    "gin-example/util/uploadTool"
    "github.com/gin-contrib/pprof"
    "github.com/gin-gonic/gin"
    swaggerFiles "github.com/swaggo/files"
    ginSwagger "github.com/swaggo/gin-swagger"
    "net/http"
)

func SetRouter() *gin.Engine {
    r := gin.Default()
    pprof.Register(r) // 性能
    //增加请求中panic拦截防止进程退出
    r.Use(recovery.Recovery())

    //文档访问链接
    r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
    //文件上传
    r.StaticFS("/upload/images", http.Dir(uploadTool.GetImageFullPath()))
    r.POST("/upload/image", ctl.Image)
    //token相关
    r.GET("/getToken", ctl.GetToken)


    apiv1 := r.Group("/v1")
    {
        //查新列表
        apiv1.GET("/saleOrder", ctl.List)
        //查新单个
        apiv1.GET("/saleOrder/:id", ctl.List)
        //新建
        apiv1.POST("/saleOrder", ctl.Create)
        //更新
        apiv1.PUT("/saleOrder/:id", ctl.Update)
        //删除
        apiv1.DELETE("/saleOrder/:id", ctl.Delete)


        apiv1.POST("/kafkaAsyncProducer", ctl.KafkaAsyncProducer)
        apiv1.POST("/kafkaProducer", ctl.KafkaProducer)
    }






    return r
}
