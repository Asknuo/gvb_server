package ctype

import "encoding/json"

type Role int

const (
	PermissionAdmin       Role = 1 //管理员
	PremissionUser        Role = 2 //普通用户
	PremissionVisitor     Role = 3 //游客
	PremissionDisableUser Role = 4 //被禁用的用户
)

// String converts the Role type to its string representation.
func (s Role) String() string {
	switch s {
	case PermissionAdmin:
		return "管理员"
	case PremissionUser:
		return "普通用户"
	case PremissionVisitor:
		return "游客"
	case PremissionDisableUser:
		return "被禁用的用户"
	default:
		return "其他"
	}
}

func (s Role) MarshallJSON() ([]byte, error) {
	return json.Marshal(s.String())
}
