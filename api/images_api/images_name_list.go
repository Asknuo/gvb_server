package images_api

import (
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/res"

	"github.com/gin-gonic/gin"
)

type ImageResponse struct {
	ID   uint   `json:"id"`
	Path string `json:"path"`
	Name string `gorm:"size:38" json:"name"`
}

func (ImagesApi) ImageNameListView(c *gin.Context) {
	var imagelist []ImageResponse
	global.DB.Model(models.BannerModel{}).Select("id", "path", "name").Scan(&imagelist)
	res.OKWithData(imagelist, c)
}
