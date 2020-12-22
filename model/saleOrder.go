package model

import (
    db "gin-example/database"
)

// 销售单
type SaleOrder struct {
    Model
    InvCode      string  `gorm:"size:32;not null;`
    InvId        int     `gorm:"not null;default:0"`
    InvName      string  `gorm:"not null;size:32;default:''"`
    InvType      int     `gorm:"not null;default:0";comment:'1库存,2代发'`
    Number       int     `gorm:"not null;default:0"`
    UnitPrice    float32 `gorm:"type:decimal(10,2);default:'0.00'"`
    TotalPrice   float32 `gorm:"type:decimal(10,2);default:'0.00'"`
    CustomerId   int     `gorm:"not null;"`
    CustomerName string  `gorm:"size:128;not null;"`
    Remark       string  `gorm:"type:text;"`
    Type         int     `gorm:"not null;default:0";comment:'1卖,2退'`
    OperatorId   int     `gorm:"not null;"`
    OperatorName string  `gorm:"not null;size:32;"`
}
func init(){
    if res := db.Mysql.HasTable(&SaleOrder{}); !res {
        db.Mysql.CreateTable(&SaleOrder{})
    }
}
func (this *SaleOrder) List() ([]*SaleOrder, error) {

    var data []*SaleOrder
    if err := db.Mysql.Limit(2).Find(&data).Error; err != nil {
        return nil, err
    }
    return data, nil
}
func (this *SaleOrder) Create() (int,error) {
    res := db.Mysql.Create(this);

    if  res.Error != nil {
        // 最后返回json格式数据
        return 0,res.Error
    }
    return this.Id,nil
}

func (this *SaleOrder) Update(id int) error {

    if err := db.Mysql.Model(&SaleOrder{}).Where("id = ? ", id).Update(this).Error; err != nil {
        return err
    }

    return nil
}

func (this *SaleOrder) Delete(id int) error {

    if err := db.Mysql.Where("id = ? ", id).Delete(&SaleOrder{}).Error; err != nil {
        return err
    }

    return nil
}
