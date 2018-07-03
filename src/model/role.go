package model

import "golang-serialization-benchmark/src/misc/httputils"

type Role struct {
	Base
	Name        string           `json:"name"`
	Status      httputils.Status `json:"-" sql:"DEFAULT:0"`
	Permission1 int64            `json:"permission1"`
	Permission2 int64            `json:"permission2"`
	Permission3 int64            `json:"permission3"`
	Weight      int              `json:"weight" sql:"DEFAULT:1"`
	PermissionArray string	 	 `json:"permission_array" gorm:"type:varchar(1024);not null;"`
}
