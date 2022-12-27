package models

import "github.com/uadmin/uadmin"

type Tag struct {
	uadmin.Model
	MachineID     string
	Enc           string `uadmin:"read_only"`
	User          uadmin.User
	UserID        uint
	ContactNumber int
	TagType       TagType
	TagTypeID     uint
}

func (e *Tag) String() string {
	return e.MachineID
}

func (t *Tag) Save() {
	uadmin.Save(t)
}
