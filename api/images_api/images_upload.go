package images_api

import (
	"gvb_server/global"
	"gvb_server/service/image_ser"

	"gvb_server/models/res"
	"gvb_server/service"

	"io/fs"
	"os"

	"github.com/gin-gonic/gin"
)

var (
	WhiteImageList = []string{
		"jpg",
		"png",
		"jpeg",
		"ico",
		"tiff",
		"gif",
		"svg",
		"webg",
	}
)

func (ImagesApi) ImageUploadView(c *gin.Context) {
	form, err := c.MultipartForm()
	if err != nil {
		res.FailWithMsg(err.Error(), c)
		return
	}
	filelist, ok := form.File["images"]
	if !ok {
		res.FailWithMsg("不存在的文件", c)
		return
	}
	//判断路径是否存在
	basePath := global.Config.Upload.Path
	_, err = os.ReadDir(basePath)
	if err != nil {
		err = os.MkdirAll(basePath, fs.ModePerm)
		if err != nil {
			global.Log.Error(err)
		}
	}
	var reslist []image_ser.FileUploadResponse
	for _, file := range filelist {
		//上传文件
		serviceRes := service.ServiceApp.ImageService.ImageUploadService(file)
		if !serviceRes.IsSuccess {
			reslist = append(reslist, serviceRes)
			continue
		}
		//成功后
		if !global.Config.QiNiu.Enable {
			//本地保存
			err = c.SaveUploadedFile(file, serviceRes.FileName)
			if err != nil {
				global.Log.Error(err)
				serviceRes.Msg = err.Error()
				serviceRes.IsSuccess = false
				reslist = append(reslist, serviceRes)
				continue
			}

		}
		reslist = append(reslist, serviceRes)
	}
	res.OKWithData(reslist, c)
}
