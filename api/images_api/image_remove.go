package images_api

import (
	"fmt"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/res"

	"github.com/gin-gonic/gin"
)

func (ImagesApi) ImageRemoveView(c *gin.Context) {
	var cr models.RemoveRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithCode(res.ArguementError, c)
		return
	}
	var imageList []models.BannerModel
	count := global.DB.Find(&imageList, cr.IDList).RowsAffected
	if count == 0 {
		res.FailWithMsg("文件不存在", c)
		return
	}
	global.DB.Delete(&imageList)
	res.OKWithMsg(fmt.Sprintf("共删除%d张图片", count), c)
}
