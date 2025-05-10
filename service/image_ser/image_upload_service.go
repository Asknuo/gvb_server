package image_ser

import (
	"fmt"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/ctype"
	"gvb_server/plugins/qiniu"
	"gvb_server/utils"
	"io"
	"mime/multipart"
	"path"
	"strings"
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

// 文件上传方法
func (ImageService) ImageUploadService(file *multipart.FileHeader) (res FileUploadResponse) {
	fileName := file.Filename
	basePath := global.Config.Upload.Path
	filePath := path.Join(basePath, file.Filename)
	res.FileName = filePath
	//文件白名单判断
	namelist := strings.Split(fileName, ".")
	suffix := strings.ToLower(namelist[len(namelist)-1])
	if !utils.IsList(suffix, WhiteImageList) {
		res.Msg = "非法文件"
		return
	}
	//判断文件大小
	size := float64(file.Size) / float64(1024*1024)
	if size >= float64(global.Config.Upload.Size) {
		res.Msg = fmt.Sprintf("图片大小超过限制的大小 %dMB", global.Config.Upload.Size)
		return
	}
	//读取文件内容hash
	fileObj, err := file.Open()
	if err != nil {
		global.Log.Error(err)
	}
	byteData, _ := io.ReadAll(fileObj)
	imageHash := utils.Md5(byteData)
	//去数据库中查这个图片是否存在
	fileType := ctype.Local
	var bannerModel models.BannerModel
	err = global.DB.Take(&bannerModel, "hash=?", imageHash).Error
	if err == nil {
		//找到了
		res.Msg = "图片已存在"
		res.FileName = bannerModel.Path
		return
	}
	fileType = ctype.Local
	res.Msg = "图片上传成功"
	res.IsSuccess = true
	if global.Config.QiNiu.Enable {
		filePath, err = qiniu.UploadImage(byteData, fileName, global.Config.QiNiu.Prefix)
		if err != nil {
			global.Log.Error(err)
			res.Msg = err.Error()
			return
		}
		res.FileName = filePath
		res.Msg = "上传到云成功"
		fileType = ctype.Qiniu
	}

	//图片入库
	global.DB.Create(&models.BannerModel{
		Path:      filePath,
		Hash:      imageHash,
		Name:      fileName,
		ImageType: fileType,
	})
	return res
}
