package menu_api

import (
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/ctype" // Adjust the import path to the correct location of the ctype package
	"gvb_server/models/res"

	"github.com/gin-gonic/gin"
)

type ImageSort struct {
	ImageID uint `json:"menu_id"`
	Sort    int  `json:"sort"`
}
type MenuRequest struct {
	Title         string      `json:"title" binding:"required" msg:" 请完善菜单名称"`
	Path          string      `json:"path" binding:"required" msg:"请完善菜单路径"`
	Slogan        string      `json:"slogan"`
	Abstract      ctype.Array `json:"abstract"`
	AbstractTime  int         `json:"abstract_time"`
	BannerTime    int         `json:"banner_time"`
	Sort          int         `json:"sort" binding:"required" msg:"请输入菜单序号"`
	ImageSortList []ImageSort `json:"image_sort_list"`
}

func (MenuApi) MenuCreateView(c *gin.Context) {
	var cr MenuRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(err, &cr, c)
		return
	}
	//重复判断
	var menuList []models.MenuModel
	count := global.DB.Find(&menuList, "title =? or path =?", cr.Title, cr.Path).RowsAffected
	if count > 0 {
		res.FailWithMsg("重复的菜单", c)
		return
	}
	//创建Banner数据入库
	menuModel := models.MenuModel{
		Title:        cr.Title,
		Path:         cr.Path,
		Slogan:       cr.Slogan,
		Abstract:     cr.Abstract,
		AbstractTime: cr.AbstractTime,
		BannerTime:   cr.BannerTime,
		Sort:         cr.Sort,
	}
	err = global.DB.Create(&menuModel).Error
	if err != nil {
		global.Log.Error(err)
		res.FailWithMsg("菜单添加失败", c)
		return
	}
	if len(cr.ImageSortList) == 0 {
		res.OKWithMsg("菜单添加成功", c)
		return
	}
	var menuBannerList []models.MenuBannerModel
	for _, sort := range cr.ImageSortList {
		//判断images_id是否真正有这张图片
		menuBannerList = append(menuBannerList, models.MenuBannerModel{
			MenuID:   menuModel.ID,
			BannerID: sort.ImageID,
			Sort:     sort.Sort,
		})
	}
	//入库
	err = global.DB.Create(&menuBannerList).Error
	if err != nil {
		global.Log.Error(err)
		res.FailWithMsg("图片关联失败", c)
		return
	}
	res.OKWithMsg("菜单添加成功", c)
}
