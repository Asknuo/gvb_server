package ctype

import "encoding/json"

type ImageType int

const (
	Local ImageType = 1 // 本地
	Qiniu ImageType = 2 // 七牛云
)

func (i ImageType) String() string {
	switch i {
	case Local:
		return "本地"
	case Qiniu:
		return "七牛云"
	default:
		return "未知"
	}
}

func (s ImageType) MarshallJSON() ([]byte, error) {
	return json.Marshal(s.String())
}
