package model

import "golang-serialization-benchmark/src/misc/httputils"

type UserGroup struct {
	Base
	GroupName string           `json:"groupname"`
	Commit    string           `json:"commit"`
	Status    httputils.Status `json:"-" sql:"DEFAULT:0"`
}
