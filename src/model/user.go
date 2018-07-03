package model

import "golang-serialization-benchmark/src/misc/httputils"

type User struct {
	Base
	Nickname  string           `json:"nickname"`
	Password  string           `json:"-"`
	Name      string           `json:"name"`
	Status    httputils.Status `json:"-" sql:"DEFAULT:0"`
	Email     string           `json:"email"`
	Telephone string           `json:"telephone"`
	Weight    int              `json:"weight" sql:"DEFAULT:1"`
}
