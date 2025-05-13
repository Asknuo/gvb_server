package advert_api

import (
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/res"

	"github.com/fatih/structs"
	"github.com/gin-gonic/gin"
)

//AdvertUpdateView 更新广告
//@Tags 广告更新
//@Summary 广告更新
//@Description 广告更新
//@Param data body AdvertRequest true "广告的一些参数"
//@Router /api/adverts/:id [put]
//@Produce json
//@Success 200 {object} res.Response{data=string}

func (AdvertApi) AdvertUpdateView(c *gin.Context) {

	id := c.Param("id")
	var cr AdvertRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(err, &cr, c)
		return
	}
	var advert models.AdvertModel
	err = global.DB.Take(&advert, id).Error
	if err != nil {
		res.FailWithMsg("广告不存在", c)
		return
	}
	maps := structs.Map(&cr)
	err = global.DB.Model(&advert).Updates(maps).Error
	//结构体转map的第三方库

	if err != nil {
		global.Log.Error(err)
		res.FailWithMsg("修改失败", c)
		return
	}
	res.OKWithMsg("修改广告成功", c)
}
