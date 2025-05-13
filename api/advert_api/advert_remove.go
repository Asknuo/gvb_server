package advert_api

import (
	"fmt"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/res"

	"github.com/gin-gonic/gin"
)

// AdvertDeleteView 广告删除
//@Tags 广告删除
//@Summary 广告删除
//@Description 批量广告删除
//@Param data body models.RemoveRequest true "广告id列表"
//@Router /api/adverts [delete]
//@Produce json
//@Success 200 {object} res.Response{data=string}

func (AdvertApi) AdvertRemoveView(c *gin.Context) {
	var cr models.RemoveRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithCode(res.ArguementError, c)
		return
	}
	var advertList []models.AdvertModel
	count := global.DB.Find(&advertList, cr.IDList).RowsAffected
	if count == 0 {
		res.FailWithMsg("广告不存在", c)
		return
	}
	global.DB.Delete(&advertList)
	res.OKWithMsg(fmt.Sprintf("共删除%d个广告", count), c)
}
