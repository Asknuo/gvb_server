package images_api

import (
	"fmt"
	"gvb_server/global"
	"gvb_server/models/res"
	"gvb_server/utils"
	"io/fs"
	"os"
	"path"
	"strings"

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

type FileUploadResponse struct {
	FileName  string `json:"file_name"`
	IsSuccess bool   `json:"is_success"`
	Msg       string `json:"msg"`
}

// 上传单个图片，返回图片的url
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
	var reslist []FileUploadResponse
	for _, file := range filelist {
		fileName := file.Filename
		namelist := strings.Split(fileName, ".")
		suffix := strings.ToLower(namelist[len(namelist)-1])
		if !utils.IsList(suffix, WhiteImageList) {
			reslist = append(reslist, FileUploadResponse{
				FileName:  file.Filename,
				IsSuccess: false,
				Msg:       "非法文件",
			})
			continue
		}
		filePath := path.Join(basePath, file.Filename)
		//判断文件大小
		size := float64(file.Size) / float64(1024*1024)
		if size >= float64(global.Config.Upload.Size) {
			reslist = append(reslist, FileUploadResponse{
				FileName:  file.Filename,
				IsSuccess: false,
				Msg:       fmt.Sprintf("图片大小超过限制的大小 %dMB", global.Config.Upload.Size),
			})
			continue
		}
		err := c.SaveUploadedFile(file, filePath)
		if err != nil {
			global.Log.Error(err)
			reslist = append(reslist, FileUploadResponse{
				FileName:  file.Filename,
				IsSuccess: false,
				Msg:       err.Error(),
			})
			continue
		}
		reslist = append(reslist, FileUploadResponse{
			FileName:  file.Filename,
			IsSuccess: true,
			Msg:       "上传成功",
		})
	}
	res.OKWithData(reslist, c)
}
