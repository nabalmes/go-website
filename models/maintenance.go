package models

import (
	"time"

	"github.com/uadmin/uadmin"
)

type Maintenance struct {
	uadmin.Model
	Tag       Tag
	TagID     uint
	Item      string
	Type      string
	TimeStart *time.Time
	TimeEnd   *time.Time
}

func (m *Maintenance) String() string {
	return m.Item
}

func (m *Maintenance) Save() {
	uadmin.Save(m)
}
