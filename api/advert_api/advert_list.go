package advert_api

import (
	"gvb_server/models"
	"gvb_server/models/res"
	"gvb_server/service/common"
	"strings"

	"github.com/gin-gonic/gin"
)

// AdvertListView 广告列表
//@Tags 广告列表
//@Summary 广告列表
//@Description 广告列表
//@Param data query models.PageInfo false "查询参数"
//@Router /api/adverts [get]
//@Produce json
//@Success 200 {object} res.Response{data=res.ListResponse[models.AdvertModel]}

func (AdvertApi) AdvertListView(c *gin.Context) {
	var cr models.PageInfo
	if err := c.ShouldBindQuery(&cr); err != nil {
		res.FailWithCode(res.ArguementError, c)
		return
	}
	is_show := true
	referer := c.GetHeader("Referer")
	if strings.Contains(referer, "admin") {
		//admin来的
		is_show = false
	}
	list, count, _ := common.ComList(models.AdvertModel{IsShow: is_show}, common.Option{
		PageInfo: cr,
		Debug:    true,
	})

	res.OKWithList(list, count, c)
}
