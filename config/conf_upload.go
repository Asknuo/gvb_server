package config

type Upload struct {
	Size int    `yaml:"size" json:"size"` //图片大小
	Path string `yaml:"path" json:"path"` //图片目录
}
