package advert_api

import (
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/res"

	"github.com/gin-gonic/gin"
)

type AdvertRequest struct {
	Title  string `json:"title" binding:"required" msg:"请输入标题" structs:"title"`        //显示的标题
	Herf   string `json:"herf" binding:"required,url" msg:"跳转链接非法" structs:"herf"`     //跳转到的链接
	Images string `json:"images" binding:"required,url" msg:"图片地址非法" structs:"images"` //图片
	IsShow bool   `json:"is_show" structs:"is_show"`                                   //是否展示
}

//AdvertCreateView 添加广告
// @Tags 广告管理
// @Summary 创建广告
// @Description 创建广告
// @Param data body AdvertRequest true "表示多个参数"
// @Router /api/adverts [post]
// @Produce json
// @Success 200 {object} res.Response{}

func (AdvertApi) AdvertCreateView(c *gin.Context) {
	var cr AdvertRequest
	err := c.ShouldBindBodyWithJSON(&cr)
	if err != nil {
		res.FailWithError(err, &cr, c)
		return
	}
	//判断是否重复
	var advert models.AdvertModel
	err = global.DB.Take(&advert, "title=?", cr.Title).Error
	if err == nil {
		res.FailWithMsg("该广告已存在", c)
		return
	}
	err = global.DB.Create(&models.AdvertModel{
		Title:  cr.Title,
		Herf:   cr.Herf,
		Images: cr.Images,
		IsShow: cr.IsShow,
	}).Error
	if err != nil {
		global.Log.Error(err)
		res.FailWithMsg("添加失败", c)
		return
	}
	res.OKWithMsg("添加广告成功", c)
}
