package upload

import (
    "fmt"
    "log"
    "mime/multipart"
    "os"
    "path"
    "strings"
    "gin-example/config"
    "gin-example/util"
    "gin-example/util/file"
)



func GetImageFullUrl(name string) string {
    return config.Upload.ImagePrefixUrl + "/" + GetImagePath() + name
}

func GetImageName(name string) string {
    ext := path.Ext(name)
    fileName := strings.TrimSuffix(name, ext)
    fileName = util.EncodeMD5(fileName)

    return fileName + ext
}

func GetImagePath() string {
    return config.Upload.ImageSavePath
}

func GetImageFullPath() string {
    return config.Upload.RuntimeRootPath + GetImagePath()
}

func CheckImageExt(fileName string) bool {
    ext := file.GetExt(fileName)
    for _, allowExt := range config.Upload.ImageAllowExts {
        if strings.ToUpper(allowExt) == strings.ToUpper(ext) {
            return true
        }
    }

    return false
}

func CheckImageSize(f multipart.File) bool {
    size, err := file.GetSize(f)
    if err != nil {
        log.Println(err)
        return false
    }
    fmt.Println("图片大小",size,config.Upload.ImageMaxSize,size <= config.Upload.ImageMaxSize)
    return size <= config.Upload.ImageMaxSize
}

func CheckImage(src string) error {
    dir, err := os.Getwd()
    if err != nil {
        return fmt.Errorf("os.Getwd err: %v", err)
    }

    err = file.IsNotExistMkDir(dir + "/" + src)
    if err != nil {
        return fmt.Errorf("file.IsNotExistMkDir err: %v", err)
    }

    perm := file.CheckPermission(src)
    if perm == true {
        return fmt.Errorf("file.CheckPermission Permission denied src: %s", src)
    }

    return nil
}
