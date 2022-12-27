package models

import "github.com/uadmin/uadmin"

type TagType struct {
	uadmin.Model
	Name string
}

func (t *TagType) String() string {
	return t.Name
}

func (t *TagType) Save() {
	uadmin.Save(t)
}
