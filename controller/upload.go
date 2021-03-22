package controller

import (
    "github.com/gin-gonic/gin"
    "gin-example/util/e"
    "gin-example/util/uploadTool"
)

type Upload struct {
}
func NewUpload() Upload {
    return Upload{};
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


    imageName := uploadTool.GetImageName(image.Filename)
    fullPath := uploadTool.GetImageFullPath()
    savePath := uploadTool.GetImagePath()

    src := fullPath + imageName
    if !uploadTool.CheckImageExt(imageName) {
        e.FailResponse(c,4)
        return
    }
    if !uploadTool.CheckImageSize(file) {
        e.FailResponse(c,1)
        return
    }

    err1 := uploadTool.CheckImage(fullPath)
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
    data["image_url"] = uploadTool.GetImageFullUrl(imageName)
    data["image_save_url"] = savePath + imageName
    e.SuccessResponse(c, data)
}
