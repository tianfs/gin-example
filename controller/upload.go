package controller

import (
    "github.com/gin-gonic/gin"
    "gin-example/util/e"
    "gin-example/util/upload"
)

type Upload struct {
}

func (this *Upload) Image(c *gin.Context) {



    file, image, err := c.Request.FormFile("image")
    if err != nil {
        e.FailResponse(c,e.ERROR_UPLOAD_CHECK_IMAGE_FAIL)
        return
    }
    if (image == nil){
        e.FailResponse(c,e.ERROR_UPLOAD_CHECK_IMAGE_FAIL)
        return
    }


    imageName := upload.GetImageName(image.Filename)
    fullPath := upload.GetImageFullPath()
    savePath := upload.GetImagePath()

    src := fullPath + imageName
    if !upload.CheckImageExt(imageName) {
        e.FailResponse(c,4)
        return
    }
    if !upload.CheckImageSize(file) {
        e.FailResponse(c,1)
        return
    }

    err1 := upload.CheckImage(fullPath)
    if err1 != nil{
        e.FailResponse(c,2)
        return
    }
    err2 := c.SaveUploadedFile(image, src);
    if err2 != nil{
        e.FailResponse(c,3)
        return
    }
    data := make(map[string]string)
    data["image_url"] = upload.GetImageFullUrl(imageName)
    data["image_save_url"] = savePath + imageName
    e.SuccessResponse(c, data)
}
