package controller

import (
    "fmt"
    "github.com/gin-gonic/gin"
    "gin-example/service"
    "gin-example/util/e"
)

type SaleOrder struct {
}

func (this *SaleOrder) List(c *gin.Context) {
    saleOrderS := service.SaleOrder{}
    list, err := saleOrderS.List()
    if err != nil {
        e.FailResponse(c, 999, "错误了")
        return
    }
    e.SuccessResponse(c, list)

}

type SaleOrderCreate struct {
    InvCode      string  `form:"inv_code" binding:"required"`
    InvId        int     `form:"inv_id" binding:"required"`
    InvName      string  `form:"inv_name" binding:"required"`
    InvType      int     `form:"inv_type" binding:"required" `
    Number       int     `form:"number" binding:"required"`
    UnitPrice    float32 `form:"unit_price" binding:"required"`
    CustomerId   int     `form:"customer_id" binding:"required"`
    CustomerName string  `form:"customer_name" binding:"required"`
    Remark       string  `form:"remark"`
    Type         int     `form:"type" binding:"required"`
}

func (this *SaleOrder) Create(c *gin.Context) {

    var from SaleOrderCreate
    if err := c.ShouldBind(&from); err != nil {
        e.ErrorResponse(c, 999, fmt.Sprint(err.Error()))
        return
    }
    saleOrderS := service.SaleOrder{
        InvCode:      from.InvCode,
        InvId:        from.InvId,
        InvName:      from.InvName,
        InvType:      from.InvType,
        Number:       from.Number,
        UnitPrice:    from.UnitPrice,
        CustomerId:   from.CustomerId,
        CustomerName: from.CustomerName,
        Remark:       from.Remark,
        Type:         from.Type,
        OperatorId:   1,
        OperatorName: "默认操作人",
    }
    id, err := saleOrderS.Create()
    if err != nil {
        e.FailResponse(c, 999, fmt.Sprint(err.Error()))
    }
    e.SuccessResponse(c, map[string]interface{}{"id": id})

}

type SaleOrderUpdata struct {
    Id           int     `form:"id" binding:"required"`
    InvCode      string  `form:"inv_code" binding:"required"`
    InvId        int     `form:"inv_id" binding:"required"`
    InvName      string  `form:"inv_name" binding:"required"`
    InvType      int     `form:"inv_type" binding:"required" `
    Number       int     `form:"number" binding:"required"`
    UnitPrice    float32 `form:"unit_price" binding:"required"`
    CustomerId   int     `form:"customer_id" binding:"required"`
    CustomerName string  `form:"customer_name" binding:"required"`
    Remark       string  `form:"remark"`
    Type         int     `form:"type" binding:"required"`
}

func (this *SaleOrder) Update(c *gin.Context) {
    var from SaleOrderUpdata
    if err := c.ShouldBind(&from); err != nil {
        e.FailResponse(c, 999, fmt.Sprint(err.Error()))
        return
    }
    saleOrderS := service.SaleOrder{
        Id:           from.Id,
        InvCode:      from.InvCode,
        InvId:        from.InvId,
        InvName:      from.InvName,
        InvType:      from.InvType,
        Number:       from.Number,
        UnitPrice:    from.UnitPrice,
        CustomerId:   from.CustomerId,
        CustomerName: from.CustomerName,
        Remark:       from.Remark,
        Type:         from.Type,
        OperatorId:   1,
        OperatorName: "默认操作人",
    }
    saleOrderS.Update()
    e.SuccessResponse(c, 1)
}

type SaleOrderDelete struct {
    Id int `form:"id" binding:"required"`
}

func (this *SaleOrder) Delete(c *gin.Context) {
    var from SaleOrderDelete
    if err := c.ShouldBind(&from); err != nil {
        e.FailResponse(c, 999, fmt.Sprint(err.Error()))
        return
    }
    saleOrderS := service.SaleOrder{
        Id: from.Id,
    }
    if err := saleOrderS.Delete(); err != nil {
        e.FailResponse(c, 999, "错误了")
        return
    }
    e.SuccessResponse(c, 1)
}

func (this *SaleOrder) KafkaProducer(c *gin.Context) {
    saleOrderS := service.SaleOrder{}
    saleOrderS.KafkaProducer()
    e.SuccessResponse(c, 1)
}
func (this *SaleOrder) KafkaAsyncProducer(c *gin.Context) {

    saleOrderS := service.SaleOrder{}
    saleOrderS.KafkaAsyncProducer()
    e.SuccessResponse(c, 1)

}
