package qiniu

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"gvb_server/config"
	"gvb_server/global"
	"time"

	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
)

func getToken(q config.QiNiu) string {
	accessKey := q.AccessKey
	secretKey := q.SecretKey
	bucket := q.Bucket
	putPolicy := storage.PutPolicy{
		Scope: bucket,
	}
	mac := qbox.NewMac(accessKey, secretKey)
	upToken := putPolicy.UploadToken(mac)
	return upToken
}

func getCfg(q config.QiNiu) storage.Config {
	cfg := storage.Config{}
	//空间对应的机房
	zone, _ := storage.GetRegionByID(storage.RegionID(q.Zone))
	cfg.Zone = &zone
	cfg.UseHTTPS = false
	cfg.UseCdnDomains = false
	return cfg
}

func UploadImage(data []byte, imageName string, prefix string) (filePath string, err error) {
	if !global.Config.QiNiu.Enable {
		return "", errors.New("没有启用服务器")
	}
	//文件名不能重复
	q := global.Config.QiNiu
	if q.AccessKey == "" || q.SecretKey == "" {
		return "", errors.New("请配置AccessKey以及secretKey")
	}
	if float64(len(data))/1024/1024 > q.Size {
		return "", errors.New("文件过大")
	}
	upToken := getToken(q)
	cfg := getCfg(q)
	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}
	putExtra := storage.PutExtra{
		Params: map[string]string{},
	}
	dataLen := int64(len(data))
	//获取当前时间
	now := time.Now().Format("202505251102")
	key := fmt.Sprintf("%s/%s__%s", prefix, now, imageName)
	err = formUploader.Put(context.Background(), &ret, upToken, key, bytes.NewReader(data), dataLen, &putExtra)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s%s", q.CDN, ret.Key), nil
}
